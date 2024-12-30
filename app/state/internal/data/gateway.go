package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/hoysics/basic-im/api/gateway/v1"
	pb "github.com/hoysics/basic-im/api/state/v1"
	"github.com/hoysics/basic-im/app/state/internal/biz"
	"github.com/hoysics/basic-im/app/state/internal/conf"
	"github.com/pkg/errors"
	sysgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"

	"google.golang.org/protobuf/proto"
)

var _gatewayRepo *gatewayRepo

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

var retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "api.gateway.v1.InnerGateway"}],
		  "retryPolicy": {
			  "MaxAttempts": 4,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`

type gwClient struct {
	rpc  v1.InnerGatewayClient
	pipe chan *pb.ChanMsg
}

/*
暂时屏蔽用Map记录多个网关接入地址的设计
State的性能理论上弱于Gateway，所以不太可能多Gateway对一个State
所以双向绑定，均为一对一
*/

type gatewayRepo struct {
	log    *log.Helper
	client *gwClient
}

func NewGatewayRepo(logger log.Logger, reg *conf.State) biz.GatewayRepo {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(reg.GetGatewayEndpoint()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			circuitbreaker.Client(),
		),
		grpc.WithOptions(sysgrpc.WithKeepaliveParams(kacp), sysgrpc.WithDefaultServiceConfig(retryPolicy)),
	)
	if err != nil {
		log.Fatalf(`grpc NewClient failed: %v`, err)
	}
	client := v1.NewInnerGatewayClient(conn)
	gwC := &gwClient{
		rpc:  client,
		pipe: make(chan *pb.ChanMsg, reg.GetPerGwPushChanSize()),
	}
	_gatewayRepo = &gatewayRepo{
		log:    log.NewHelper(logger),
		client: gwC,
	}
	return _gatewayRepo
}

func (g *gatewayRepo) PushPipe() chan *pb.ChanMsg {
	return g.client.pipe
}

func (g *gatewayRepo) Push(ctx context.Context, cmd *biz.CmdCtx) {
	if g.client == nil {
		g.log.Info("gateway repo client is nil")
		return
	}
	data, err := proto.Marshal(cmd.Msg)
	if err != nil {
		g.log.Errorf("Marshal: %v", err)
	}
	select {
	case <-ctx.Done():
		g.log.Infof(`Push.err: ctx is canceled: %v`, ctx.Err())
	case g.client.pipe <- &pb.ChanMsg{ConnId: cmd.ConnId, Data: data}:
	}
	return
}

func (g *gatewayRepo) DelConn(ctx context.Context, connId uint64) error {
	rp, err := g.client.rpc.DelConn(ctx, &v1.InnerGatewayRequest{ConnId: connId, Data: nil})
	if err != nil {
		return errors.Wrap(biz.ErrWhenCallGwDelConn, err.Error())
	}
	if rp == nil {
		return biz.ErrGwDelConn
	}
	if rp.Code == 0 {
		return nil
	}
	g.log.Infof(`DelConn.删除失败: %v`, rp)
	return errors.Wrap(biz.ErrGwDelConn, rp.Msg)
}

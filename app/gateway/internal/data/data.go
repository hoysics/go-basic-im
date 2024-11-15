package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	pb "github.com/hoysics/basic-im/api/state/v1"
	"github.com/hoysics/basic-im/app/gateway/internal/conf"
	clientV3 "go.etcd.io/etcd/client/v3"
	sysgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"strings"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewStateRepo)

// Data .
type Data struct {
	ctx      context.Context
	endpoint string
	state    pb.StateClient
}

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

var retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "api.state.v1.State"}],
		  "retryPolicy": {
			  "MaxAttempts": 4,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`

// To ensure resources are not leaked due to the stream returned, one of the following
// actions must be performed:
//
//  1. Call Close on the ClientConn.
//  2. Cancel the context provided.
//  3. Call RecvMsg until a non-nil error is returned. A protobuf-generated
//     client-streaming RPC, for instance, might use the helper function
//     CloseAndRecv (note that CloseSend does not Recv, therefore is not
//     guaranteed to release all resources).
//  4. Receive a non-nil, non-io.EOF error from Header or SendMsg.
//
// If none of the above happen, a goroutine and a context will be leaked, and grpc
// will not call the optionally-configured stats handler with a stats.End message.

// NewData .
func NewData(
	logger log.Logger,
	gw *conf.Gateway,
) (*Data, func(), error) {
	// Init State Server Client
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(gw.GetStateServerEndpoint()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			circuitbreaker.Client(),
		),
		grpc.WithOptions(sysgrpc.WithKeepaliveParams(kacp), sysgrpc.WithDefaultServiceConfig(retryPolicy)),
	)
	if err != nil {
		log.Fatalf("grpc fail to dial: %v", err)
	}
	// Init Main Ctx
	ctx, cancel := context.WithCancel(context.Background())
	cleanup := func() {
		h := log.NewHelper(logger)
		h.Info("closing the data resources")
		if err := conn.Close(); err != nil {
			h.Infof(`close the state conn fail: %v`, err)
		}
		cancel()
		h.Info("closed the data resources")
	}
	//TODO endpoint改成动态取
	return &Data{
		endpoint: "127.0.0.1:9001",
		state:    pb.NewStateClient(conn),
		ctx:      ctx,
	}, cleanup, nil
}

func NewDiscovery(reg *conf.Registry) registry.Discovery {
	addr := strings.Split(reg.GetEtcd().GetAddr(), `,`)
	cli, err := clientV3.New(clientV3.Config{
		DialTimeout: reg.GetEtcd().GetTimeout().AsDuration(),
		Endpoints:   addr,
		Username:    reg.GetEtcd().GetUsername(),
		Password:    reg.GetEtcd().GetPassword(),
	})
	if err != nil {
		log.Fatalf("discovery etcd connect error: %v", err)
	}
	r := etcd.New(cli)
	return r
}

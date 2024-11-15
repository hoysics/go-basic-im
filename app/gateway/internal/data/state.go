package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/hoysics/basic-im/api/state/v1"
	"github.com/hoysics/basic-im/app/gateway/internal/biz"
	"github.com/hoysics/basic-im/app/gateway/internal/conf"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
)

var (
	_endpointHead = `GwEndpoint`
)

type stateRepo struct {
	data *Data

	log *log.Helper

	ctx context.Context

	callback biz.OnStateCall
	sendCmd  chan *biz.CmdContext
}

func (r *stateRepo) CancelConn(u uint64) {
	_, err := r.data.state.CancelConn(context.TODO(), &v1.StateRequest{
		Endpoint: r.data.endpoint,
		ConnId:   u,
	})
	if err != nil {
		r.log.Errorf(`call CancelConn(%d): %v`, u, err)
	}
}

func (r *stateRepo) RunCmdProc(callback biz.OnStateCall) error {
	r.callback = callback
	ctx, cancel := context.WithCancel(r.data.ctx)
	r.ctx = ctx
	// 传递endpoint
	md := metadata.Pairs(_endpointHead, r.data.endpoint)
	mctx := metadata.NewOutgoingContext(ctx, md)
	stream, err := r.data.state.CreateMsgChan(mctx)
	if err != nil {
		cancel()
		return err
	}
	go func(ctx context.Context) {
		r.sendCmdProc(ctx, r.sendCmd, stream)
	}(ctx)
	go func() {
		defer cancel()
		r.pushCmdProc(stream)
	}()
	return nil
}
func (r *stateRepo) SendMsg(ctx context.Context, cmd *biz.CmdContext) error {
	//if !cb.AllowRequest() {
	//	fmt.Println("Circuit breaker is open, request blocked")
	//	time.Sleep(100 * time.Millisecond) // 等待一段时间后重试
	//	return nil
	select {
	case <-r.ctx.Done():
		return errors.Wrap(r.ctx.Err(), "data context canceled")
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "req context canceled")
	case r.sendCmd <- cmd:
	}
	return nil
	//}
}

// 上行消息转发处理
func (r *stateRepo) sendCmdProc(ctx context.Context, sendCmd chan *biz.CmdContext, stream v1.State_CreateMsgChanClient) {
	r.log.Info(`start send cmd`)
	for cmd := range sendCmd {
		if ctx.Err() != nil {
			break
		}
		err := stream.Send(&v1.ChanMsg{
			ConnId: cmd.ConnId,
			Data:   cmd.Payload,
		})
		if err == io.EOF {
			fmt.Println("Stream RPC completed successfully")
			//cb.RecordSuccess()
			break
		}
		if err != nil {
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.Unavailable:
					// 可能需要重试或使用备用服务
				case codes.DeadlineExceeded:
					// 可能需要调整超时设置或重试
				default:
					// 其他错误处理
				}
			}
			r.log.Infof("Receive error: %v", err)
			//cb.RecordFailure()
			//if shouldRetryRPC(err) {
			//	// 实施重试逻辑
			//	time.Sleep(retryDelay(1)) // 示例：简单重试
			//	continue
			//}
			break
		}
	}
	r.log.Info("send cmd completed successfully")
}

// 下行消息转发处理
func (r *stateRepo) pushCmdProc(stream v1.State_CreateMsgChanClient) {
	// 正常处理流式 RPC
	r.log.Info("start push cmd")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			r.log.Infof("Stream RPC completed successfully")
			//cb.RecordSuccess()
			break
		}
		if err != nil {
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.Unavailable:
					// 可能需要重试或使用备用服务
				case codes.DeadlineExceeded:
					// 可能需要调整超时设置或重试
				default:
					// 其他错误处理
				}
			}
			r.log.Infof("Receive error: %v", err)
			//cb.RecordFailure()
			//if shouldRetryRPC(err) {
			//	// 实施重试逻辑
			//	time.Sleep(retryDelay(1)) // 示例：简单重试
			//	continue
			//}
			break
		}
		// 处理接收到的消息
		//r.log.Infof("Received message: %v", msg)
		r.callback.PushMsgByCmd(&biz.CmdContext{
			ConnId:  msg.ConnId,
			Payload: msg.Data,
		})
	}
	r.log.Info("push cmd completed successfully")
}

// NewStateRepo .
func NewStateRepo(data *Data, logger log.Logger, gw *conf.Gateway) biz.StateRepo {
	//cb := NewCircuitBreaker()
	return &stateRepo{
		data:    data,
		log:     log.NewHelper(logger),
		sendCmd: make(chan *biz.CmdContext, gw.CmdChannelNum),
	}
}

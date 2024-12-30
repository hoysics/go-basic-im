package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hoysics/basic-im/api/gateway/v1"
	"github.com/hoysics/basic-im/app/gateway/internal/conf"
	"github.com/panjf2000/ants/v2"
	"sync"
)

// StateRepo is a Greater repo.
type StateRepo interface {
	RunCmdProc(callback OnStateCall) error
	SendMsg(context.Context, *CmdContext) error
	CancelConn(uint64)
}

type OnStateCall interface {
	PushMsgByCmd(cmd *CmdContext)
}

// ConnectUseCase 由于gRPC限制，不能在应用层使用Conn管理Socket，故此处仅仅能完成原定方案中一半的设计，即：map+(wait->call state)
type ConnectUseCase struct {
	repo StateRepo
	log  *log.Helper

	epoll *HalfEpoll
}

func NewWorkPool(size int) (wPool *ants.Pool) {
	var err error
	if wPool, err = ants.NewPool(size); err != nil {
		log.Fatalf("InitWorkPoll.err :%v num:%d\n", err, size)
	}
	return
}

// NewConnUseCase new a half epoll
func NewConnUseCase(logger log.Logger, repo StateRepo, gw *conf.Gateway) *ConnectUseCase {
	wPool := NewWorkPool(int(gw.WorkerPoolNum))
	l := log.NewHelper(logger)
	ep := &HalfEpoll{
		log:                l,
		table:              sync.Map{},
		wPool:              wPool,
		perConnMsgChanSize: int(gw.PerConnMsgChanSize),
	}
	go func() {
		if err := repo.RunCmdProc(ep); err != nil {
			log.Fatalf("repo.RunCmdProc.err :%v", err)
		}
	}()
	return &ConnectUseCase{
		log:   l,
		epoll: ep,
		repo:  repo,
	}
}

// Connect 接入新连接
func (c *ConnectUseCase) Connect(ctx context.Context, cancel context.CancelFunc, conn pb.Gateway_ConnectServer) (uint64, error) {
	return c.epoll.Add(ctx, cancel, conn)
}

// Disconnect 由客户端断开连接触发的普通清理
func (c *ConnectUseCase) Disconnect(connId uint64) {
	c.epoll.CloseConn(connId)
	c.repo.CancelConn(connId)
}

// DeleteConn 由State因心跳过期等原因主动发起的断开连接
func (c *ConnectUseCase) DeleteConn(connId uint64) {
	//不能再调用state的接口 避免回环
	c.epoll.CloseConn(connId)
}

func (c *ConnectUseCase) SendMsg(ctx context.Context, connId uint64, body []byte) error {
	return c.repo.SendMsg(ctx, &CmdContext{
		ConnId:  connId,
		Payload: body,
	})
}

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hoysics/basic-im/api/gateway/v1"
	ants "github.com/panjf2000/ants/v2"
	"sync"
)

type CmdContext struct {
	ConnId  uint64
	Payload []byte
}

type HalfEpoll struct {
	log   *log.Helper
	table sync.Map

	wPool *ants.Pool

	perConnMsgChanSize int
}

func (ep *HalfEpoll) PushMsgByCmd(cmd *CmdContext) {
	if connPtr, ok := ep.table.Load(cmd.ConnId); ok {
		conn, _ := connPtr.(*Connection)
		if err := conn.Push(cmd.Payload); err != nil {
			ep.log.Errorf(`run proc failed: %v`, err)
		}
	}
}

// Add 接入新连接
func (ep *HalfEpoll) Add(ctx context.Context, cancel context.CancelFunc, conn pb.Gateway_ConnectServer) (uint64, error) {
	connection, err := NewConnection(ctx, cancel, conn, ep.perConnMsgChanSize)
	if err != nil {
		return 0, err
	}
	if err = ep.wPool.Submit(func() {
		ep.log.Info("begin handle connection")
		connection.RunPushProc()
		ep.log.Info("end handle connection")
	}); err != nil {
		return 0, err
	}
	ep.table.Store(connection.id, connection)
	return connection.id, nil
}

func (ep *HalfEpoll) CloseConn(connId uint64) {
	if connPtr, ok := ep.table.Load(connId); ok {
		conn, _ := connPtr.(*Connection)
		conn.Close()
		ep.table.Delete(connId)
	}
}

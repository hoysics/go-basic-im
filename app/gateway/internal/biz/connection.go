package biz

import (
	"context"
	pb "github.com/hoysics/basic-im/api/gateway/v1"
)

type Connection struct {
	id   uint64
	conn pb.Gateway_ConnectServer

	ctx    context.Context
	cancel context.CancelFunc

	msg chan []byte
}

func NewConnection(ctx context.Context, cancel context.CancelFunc, conn pb.Gateway_ConnectServer, chanSize int) (*Connection, error) {
	id, err := node.NextId()
	if err != nil {
		return nil, err
	}
	return &Connection{
		id:     id,
		conn:   conn,
		ctx:    ctx,
		cancel: cancel,
		msg:    make(chan []byte, chanSize),
	}, nil
}
func (c *Connection) Push(msg []byte) error {
	select {
	case <-c.ctx.Done():
		return c.ctx.Err()
	case c.msg <- msg:
	}
	return nil
}

func (c *Connection) RunPushProc() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case msg := <-c.msg:
			_ = c.conn.Send(&pb.Message{Data: msg})
		}
	}
}

func (c *Connection) Close() {
	c.cancel()
	close(c.msg)
}

func (c *Connection) RemoteAddr() string {
	return ""
}

package sdk

import (
	"context"
	"encoding/json"
	pb "github.com/hoysics/basic-im/api/common/v1"
	v1 "github.com/hoysics/basic-im/api/gateway/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"sync/atomic"
)

type connect struct {
	rawCtx   context.Context
	client   v1.GatewayClient
	connId   uint64
	ctx      context.Context
	cancel   context.CancelFunc
	sendChan chan []byte
	recvChan chan *Message
}

func newConnect(ct context.Context, serverAddr string) *connect {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		_log.Fatal(err)
	}
	client := v1.NewGatewayClient(conn)
	ctx, cancel := context.WithCancel(ct)
	stream, err := client.Connect(ctx)
	if err != nil {
		_log.Fatal(err)
	}
	con := &connect{
		rawCtx:   ct,
		client:   client,
		ctx:      ctx,
		cancel:   cancel,
		sendChan: make(chan []byte),   //由于cui的读写限制，不能重建，必须复用
		recvChan: make(chan *Message), //由于cui的读写限制，不能重建，必须复用
	}
	go sendStreamProc(ctx, stream, con.sendChan)
	go recvStreamProc(ctx, &con.connId, stream, con.sendChan, con.recvChan)
	return con
}

func (c *connect) reConn() {
	c.close()
	c.ctx, c.cancel = context.WithCancel(c.rawCtx)
	stream, err := c.client.Connect(c.ctx)
	if err != nil {
		_log.Error(err)
		return
	}
	go sendStreamProc(c.ctx, stream, c.sendChan)
	go recvStreamProc(c.ctx, &c.connId, stream, c.sendChan, c.recvChan)
}

func (c *connect) send(ty pb.CmdType, payload []byte) error {
	mm := &pb.MajorMsg{
		Cmd:     ty,
		Payload: payload,
	}
	data, err := proto.Marshal(mm)
	if err != nil {
		return err
	}
	_log.Infof(`send major msg: %v`, mm)
	select {
	case <-c.ctx.Done():
		return c.ctx.Err()
	case c.sendChan <- data:
		return nil
	}
}

func (c *connect) recv() <-chan *Message {
	return c.recvChan
}

func (c *connect) close() {
	c.cancel()
	// 目前没啥值得回收的
}

func sendStreamProc(ctx context.Context, stream v1.Gateway_ConnectClient, send chan []byte) {
	_log.Info(`begin send message`)
	for {
		select {
		case <-ctx.Done():
			_log.Infof("stream.Context().Done() %v", stream.Context().Err())
			return
		case msg := <-send:
			if err := stream.Send(&v1.Message{Data: msg}); err != nil {
				_log.Error(err)
			}
		}
	}
}

func recvStreamProc(ctx context.Context, connId *uint64, stream v1.Gateway_ConnectClient, sendChan chan<- []byte, recvChan chan<- *Message) {
	_log.Info(`begin receive message`)
	for {
		by, err := stream.Recv()
		if err != nil {
			_log.Error(err)
			return
		}
		_log.Infof(`receive message: %v`, by)
		mm := &pb.MajorMsg{}
		if err := proto.Unmarshal(by.Data, mm); err != nil {
			_log.Error(err)
			continue
		}
		switch mm.Cmd {
		case pb.CmdType_ACK:
			ackMsg := &pb.AckMsg{}
			if err := proto.Unmarshal(mm.Payload, ackMsg); err != nil {
				_log.Error(err)
				continue
			}
			switch ackMsg.Cmd {
			case pb.CmdType_LOGIN:
				atomic.StoreUint64(connId, ackMsg.ConnId)
			}
			msg := &Message{
				Type:       MsgTypeAck,
				Name:       "SDK-Info",
				FormUserID: "1212121",
				ToUserID:   "222212122",
				Content:    ackMsg.Msg,
				Session:    "",
			}
			select {
			case <-ctx.Done():
			case recvChan <- msg:
			}
		case pb.CmdType_PUSH:
			pushMsg := &pb.PushMsg{}
			if err := proto.Unmarshal(mm.Payload, pushMsg); err != nil {
				_log.Error(err)
				continue
			}
			//TODO MaxClientId的校验
			// if pushMsg.MsgID == c.maxMsgID+1 {
			// 	c.maxMsgID++
			//发Ack
			ackMsg := &pb.AckMsg{
				Cmd:       pb.CmdType_PUSH,
				ConnId:    atomic.LoadUint64(connId),
				SessionId: pushMsg.SessionId,
				MsgId:     pushMsg.MsgId,
			}
			pd, err := proto.Marshal(ackMsg)
			if err != nil {
				_log.Error(err)
				continue
			}
			mm := &pb.MajorMsg{
				Cmd:     pb.CmdType_ACK,
				Payload: pd,
			}
			data, err := proto.Marshal(mm)
			if err != nil {
				return
			}
			_log.Infof(`send push-ack msg: %v`, mm)
			select {
			case <-ctx.Done():
			case sendChan <- data:
			}
			//推送消息显示
			msg := &Message{}
			if err = json.Unmarshal(pushMsg.Content, msg); err != nil {
				_log.Error(err)
				continue
			}
			select {
			case <-ctx.Done():
			case recvChan <- msg:
			}
			// }
		}
	}
}

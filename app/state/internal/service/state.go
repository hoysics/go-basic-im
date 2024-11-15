package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/hoysics/basic-im/api/common/v1"
	pb "github.com/hoysics/basic-im/api/state/v1"
	"github.com/hoysics/basic-im/app/state/internal/biz"
	"github.com/hoysics/basic-im/app/state/internal/conf"
	"google.golang.org/protobuf/proto"
	"io"
)

type StateService struct {
	pb.UnimplementedStateServer

	log *log.Helper

	ch *biz.CmdHandler
	gw *biz.GatewayClient

	cmdSize uint32
}

func NewStateService(logger log.Logger, reg *conf.State, ch *biz.CmdHandler, gw *biz.GatewayClient) *StateService {
	return &StateService{
		log:     log.NewHelper(logger),
		ch:      ch,
		gw:      gw,
		cmdSize: reg.PerGwCmdChanSize,
	}
}

func (s *StateService) CancelConn(ctx context.Context, req *pb.StateRequest) (*pb.StateResponse, error) {
	s.log.Infof("CancelConn endpoint:%s, fd:%d, data:%+v", req.Endpoint, req.ConnId, req.Data)
	msg, err := s.ch.Logout(ctx, req.ConnId)
	rp := &pb.StateResponse{Code: 0, Msg: msg}
	if err != nil {
		s.log.Errorf("CancelConn failed: %v", err)
		rp.Code = 1
	}
	return rp, nil
}

func (s *StateService) CreateMsgChan(conn pb.State_CreateMsgChanServer) error {
	rawCtx := conn.Context()
	push := s.gw.PushPipe()
	// 开始处理请求
	ctx, cancel := context.WithCancel(rawCtx)
	defer cancel()
	cmds := make(chan *biz.CmdCtx, s.cmdSize)
	msgs := make(chan *biz.CmdCtx, s.cmdSize)
	acks := make(chan *biz.CmdCtx, s.cmdSize)
	go s.runPushProc(ctx, conn, push)
	go s.runCmdProc(ctx, cmds, push)
	go s.runMsgProc(ctx, msgs, push)
	go s.runAckProc(ctx, acks)
	for {
		req, err := conn.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		major := &biz.CmdCtx{ConnId: req.ConnId, Msg: &v1.MajorMsg{}}
		if err = proto.Unmarshal(req.GetData(), major.Msg); err != nil {
			s.log.Errorf("proto.Unmarshal: %v", err)
			continue
		}
		switch major.Msg.Cmd {
		case v1.CmdType_UP:
			select {
			case <-ctx.Done():
			case msgs <- major:
			}
		case v1.CmdType_ACK:
			select {
			case <-ctx.Done():
			case acks <- major:
			}
		case v1.CmdType_PUSH:
			continue
		default:
			select {
			case <-ctx.Done():
			case cmds <- major:
			}
		}
	}
}
func (s *StateService) runAckProc(ctx context.Context, acks <-chan *biz.CmdCtx) {
	for {
		select {
		case <-ctx.Done():
		case ack := <-acks:
			//TODO 先尝试不用协程 看下性能如何 pushAck的数量在目前业务场景里可能没有上行消息那么多 而且相对来说处理状态简单 只需要重置计时器
			s.ch.AckMsg(ctx, ack)
		}
	}
}

func (s *StateService) runMsgProc(ctx context.Context, msgs <-chan *biz.CmdCtx, sendAck chan<- *pb.ChanMsg) {
	for {
		select {
		case <-ctx.Done():
		case major := <-msgs:
			//这个并发会不会影响消息的有序性 如果影响 补偿方案成本高不高 和在这里去掉并发（影响性能）比 哪个划算
			// -> 会影响：客户端快速连续发两条消息的时候 单个Session内的消息收发应该是线性、有序的
			// -> 至于性能的优化应该在后续整个功能跑通后再处理
			//err = biz.WorkerPool.Submit(func() {
			var clientId uint64
			var ok bool
			var ack *v1.AckMsg
			if clientId, ok = s.ch.UpMsg(ctx, major); !ok {
				s.log.Infof(`UpMsg.failed: %v`, major)
				return
				//ack = &v1.AckMsg{Code: 1, Msg: fmt.Sprintf(`UpMsg.err: %v`, err), Cmd: v1.CmdType_UP, ConnId: major.ConnId, ClientId: clientId}
			} else {
				ack = &v1.AckMsg{Code: 0, Msg: "up-ok", Cmd: v1.CmdType_UP, ConnId: major.ConnId, ClientId: clientId}
			}
			dl, err := proto.Marshal(ack)
			if err != nil {
				s.log.Errorf("Marshal: %v", err)
				return
			}
			data, err := proto.Marshal(&v1.MajorMsg{Cmd: v1.CmdType_ACK, Payload: dl})
			if err != nil {
				s.log.Errorf("Marshal: %v", err)
				return
			}
			select {
			case <-ctx.Done():
			case sendAck <- &pb.ChanMsg{ConnId: major.ConnId, Data: data}:
			}
			//TODO Mock一个Echo
			if err = s.ch.PushMsg(context.TODO(), major.ConnId, clientId, []byte(`{"Type":"text","Name":"EchoBot","FormUserID":"222222","ToUserID":"123213","Content":"你好 这一条Mock的回复消息","Session":""}`)); err != nil {
				s.log.Errorf(`Mock.PushMsg.err: %v`, err)
			}
			//})
			//if err != nil {
			//	s.log.Errorf("biz.WorkerPool.Submit: %v", err)
			//}
		}
	}
}

func (s *StateService) runCmdProc(ctx context.Context, cmds <-chan *biz.CmdCtx, sendAck chan<- *pb.ChanMsg) {
	for {
		select {
		case <-ctx.Done():
		case cmd := <-cmds:
			var err error
			var rp string
			switch cmd.Msg.Cmd {
			case v1.CmdType_LOGIN:
				if rp, err = s.ch.Login(ctx, cmd); err != nil {
					s.log.Errorf("Login: %v", err)
				}
			case v1.CmdType_HEARTBEAT:
				//TODO 为了减少通信量 可以暂时不回复心跳的ACK
				s.ch.Heartbeat(ctx, cmd)
				continue
			case v1.CmdType_RE_CONN:
				if rp, err = s.ch.ReConn(ctx, cmd); err != nil {
					s.log.Errorf("Login: %v", err)
				}
			}
			s.log.Infof(`cmd ack: %v`, cmd)
			//这里的CmdType有什么用 => 客户端SDK中根据这个type判断是什么类型的消息、分别该怎么处理
			dl, err := proto.Marshal(&v1.AckMsg{Code: 0, Msg: rp, Cmd: cmd.Msg.Cmd, ConnId: cmd.ConnId, ClientId: 0})
			if err != nil {
				s.log.Errorf("Marshal: %v", err)
				continue
			}
			data, err := proto.Marshal(&v1.MajorMsg{Cmd: v1.CmdType_ACK, Payload: dl})
			if err != nil {
				s.log.Errorf("Marshal: %v", err)
				continue
			}
			select {
			case <-ctx.Done():
			case sendAck <- &pb.ChanMsg{ConnId: cmd.ConnId, Data: data}:
			}
		}
	}
}

func (s *StateService) runPushProc(ctx context.Context, conn pb.State_CreateMsgChanServer, push <-chan *pb.ChanMsg) {
	//TODO 通过引用计数 处理chan的close
	for {
		select {
		case <-ctx.Done():
		case msg := <-push:
			s.log.Debugf("push msg: %v", msg)
			if err := conn.Send(msg); err != nil {
				s.log.Errorf("Send: %v", err)
			}
		}
	}
}

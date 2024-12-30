package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/hoysics/basic-im/api/common/v1"
	"github.com/hoysics/basic-im/app/state/internal/conf"
	"google.golang.org/protobuf/proto"
)

var (
	ErrNilHead = errors.New(1, `no head`, ``)
)

//var WorkerPool *ants.Pool

// CmdCtx is a Greeter model.
type CmdCtx struct {
	ConnId uint64
	Msg    *v1.MajorMsg
}

// ConnStateRepo 长连接状态管理，负责定时器及和Cache交互等
type ConnStateRepo interface {
	ConnLogin(ctx context.Context, did, connId uint64) error //Login(ctx context.Context, endpoint string, connId uint64) error
	ConnLogout(ctx context.Context, connId uint64, callByGw bool) (uint64, error)
	ResetHeartTimer(connId uint64)
	ReConn(ctx context.Context, oldConnId, newConnId uint64) error
	CompareAndIncrClientId(ctx context.Context, connId, oldMaxClientId uint64, sessionId string) bool
	Push(ctx context.Context, connId uint64, msg *v1.PushMsg) error
	AckLastMsg(ctx context.Context, connId, sessionId, msgId uint64)
}

// CmdHandler is a control signal use case.
type CmdHandler struct {
	connToState ConnStateRepo
	log         *log.Helper
}

// NewCmdHandler new a control signal use case.
func NewCmdHandler(connToState ConnStateRepo, logger log.Logger, reg *conf.State) *CmdHandler {
	//var err error
	//if WorkerPool, err = ants.NewPool(int(reg.WorkerPoolNum)); err != nil {
	//	log.Fatalf("InitWorkPoll.err :%s num:%d\n", err.Error(), reg.WorkerPoolNum)
	//}
	return &CmdHandler{
		connToState: connToState,
		log:         log.NewHelper(logger),
	}
}

// Login 登录信令处理
func (ch *CmdHandler) Login(ctx context.Context, g *CmdCtx) (string, error) {
	login := &v1.LoginMsg{}
	err := proto.Unmarshal(g.Msg.Payload, login)
	if err != nil {
		return "login failed", err
	}
	if login.Head != nil {
		//TODO 把login Msg传给业务层做处理
		ch.log.Infof("login head: %v", login.Head)
	}
	ch.log.WithContext(ctx).Infof("Login: %v", g.ConnId)
	if err = ch.connToState.ConnLogin(ctx, login.Head.DeviceId, g.ConnId); err != nil {
		return "login failed", err
	}
	return "login ok", nil
}

func (ch *CmdHandler) Logout(ctx context.Context, connId uint64) (string, error) {
	if _, err := ch.connToState.ConnLogout(ctx, connId, true); err != nil {
		return "logout failed", err
	}
	return "logout ok", nil
}

// Heartbeat 心跳处理
func (ch *CmdHandler) Heartbeat(ctx context.Context, g *CmdCtx) {
	hb := &v1.HeartbeatMsg{}
	if err := proto.Unmarshal(g.Msg.Payload, hb); err != nil {
		ch.log.Errorf("Heartbeat(%v) Unmarshal err: %v", g.ConnId, err)
		return
	}
	ch.log.WithContext(ctx).Infof("Heartbeat connId: %v", g.ConnId)
	ch.connToState.ResetHeartTimer(g.ConnId)
}

// ReConn 快速重连处理
func (ch *CmdHandler) ReConn(ctx context.Context, g *CmdCtx) (string, error) {
	//状态复用，仅更换connId
	reConn := &v1.ReConnMsg{}
	err := proto.Unmarshal(g.Msg.Payload, reConn)
	if err != nil {
		return "re-conn failed", err
	}
	if reConn.Head == nil {
		return "re-conn failed", ErrNilHead
	}
	ch.log.WithContext(ctx).Infof("ReConn: %v", g.ConnId)
	if err = ch.connToState.ReConn(ctx, reConn.Head.ConnId, g.ConnId); err != nil {
		return "re-conn failed", err
	}
	return "re-conn ok", nil
}

func (ch *CmdHandler) UpMsg(ctx context.Context, g *CmdCtx) (uint64, bool) {
	upMsg := &v1.UpMsg{}
	err := proto.Unmarshal(g.Msg.Payload, upMsg)
	if err != nil {
		ch.log.Errorf("UpMsg(%v) Unmarshal err: %v", g.ConnId, err)
		return 0, false
	}
	if upMsg.Head == nil {
		ch.log.Errorf("UpMsg(%v) failed: upMsg.Head == nil ", g.ConnId)
		return 0, false
	}
	return upMsg.Head.ClientId, ch.connToState.CompareAndIncrClientId(ctx, g.ConnId, upMsg.Head.ClientId, upMsg.Head.SessionId)
}

func (ch *CmdHandler) PushMsg(ctx context.Context, connId, msgId uint64, content []byte) error {
	pushMsg := &v1.PushMsg{MsgId: msgId, Content: content}
	//下行消息的下发：不管成功与否，都要更新LastMsg => 放在repo层处理
	return ch.connToState.Push(ctx, connId, pushMsg)
}

// AckMsg 处理下行消息的Ack回复
func (ch *CmdHandler) AckMsg(ctx context.Context, g *CmdCtx) {
	ackMsg := &v1.AckMsg{}
	if err := proto.Unmarshal(g.Msg.Payload, ackMsg); err != nil {
		ch.log.Errorf(`Ack Msg Unmarshal err: %v`, err)
		return
	}
	ch.log.WithContext(ctx).Infof("Ack connId: %v,%v", g.ConnId, ackMsg)
	//从时序来说，由于gRPC的封装和客户端与服务端约定的只处理最新一条Msg的Ack，不用担心因为网络问题导致的类似于tcp三次握手所处理的ack乱序问题
	ch.connToState.AckLastMsg(ctx, g.ConnId, ackMsg.SessionId, ackMsg.MsgId)
}

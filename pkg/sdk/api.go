package sdk

import (
	"context"
	"encoding/json"
	pb "github.com/hoysics/basic-im/api/common/v1"
	"google.golang.org/protobuf/proto"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MsgTypeText = "text"
	MsgTypeAck  = "ack"
	//MsgTypeReConn    = "reConn"
	//MsgTypeHeartbeat = "heartbeat"
	//MsgLogin         = "loginMsg"
)

type Chat struct {
	sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	Nick      string
	UserID    string
	SessionID string
	conn      *connect

	MsgClientIDTable map[string]uint64
}

type Message struct {
	Type       string
	Name       string
	FormUserID string
	ToUserID   string
	Content    string
	Session    string
}

func NewChat(serverAddr, nick, userID, sessionID string) *Chat {
	ctx, cancel := context.WithCancel(context.Background())
	chat := &Chat{
		ctx:              ctx,
		cancel:           cancel,
		Nick:             nick,
		UserID:           userID,
		SessionID:        sessionID,
		conn:             newConnect(ctx, serverAddr),
		MsgClientIDTable: make(map[string]uint64),
	}
	chat.login()
	go chat.heartbeat()
	return chat
}

func (chat *Chat) ReConn() {
	chat.Lock()
	defer chat.Unlock()
	chat.MsgClientIDTable = make(map[string]uint64)
	chat.conn.reConn()
	chat.reConn()
}

func (chat *Chat) getClientID(sessionID string) uint64 {
	chat.Lock()
	defer chat.Unlock()
	var res uint64
	if id, ok := chat.MsgClientIDTable[sessionID]; ok {
		res = id
	}
	chat.MsgClientIDTable[sessionID] = res + 1
	return res
}

func (chat *Chat) Send(msg *Message) {
	by, _ := json.Marshal(msg)
	_log.Infof(`send message: %v`, string(by))
	pl := &pb.UpMsg{
		Head: &pb.UpMsgHead{
			ConnId:    atomic.LoadUint64(&chat.conn.connId),
			ClientId:  chat.getClientID(chat.SessionID),
			SessionId: chat.SessionID,
		},
		Body: by,
	}
	pd, err := proto.Marshal(pl)
	if err != nil {
		return
	}
	err = chat.conn.send(pb.CmdType_UP, pd)
	if err != nil {
		return
	}
}

func (chat *Chat) GetCurClientId() uint64 {
	if id, ok := chat.MsgClientIDTable[chat.SessionID]; ok {
		return id
	}
	return 0
}

// Close chat
func (chat *Chat) Close() {
	chat.cancel()
	chat.conn.close()
}

// Recv receive message
func (chat *Chat) Recv() <-chan *Message {
	return chat.conn.recv()
}

func (chat *Chat) GetConnID() uint64 {
	return atomic.LoadUint64(&chat.conn.connId)
}

func (chat *Chat) login() {
	pl := &pb.LoginMsg{
		Head: &pb.LoginMsgHead{
			DeviceId: 123,
		},
	}
	pd, err := proto.Marshal(pl)
	if err != nil {
		return
	}
	err = chat.conn.send(pb.CmdType_LOGIN, pd)
	if err != nil {
		return
	}
}

func (chat *Chat) reConn() {
	pl := &pb.ReConnMsg{
		Head: &pb.ReConnMsgHead{ConnId: atomic.LoadUint64(&chat.conn.connId)},
	}
	_log.Infof(`reConnMsg: %v`, pl)
	pd, err := proto.Marshal(pl)
	if err != nil {
		return
	}
	err = chat.conn.send(pb.CmdType_RE_CONN, pd)
	if err != nil {
		return
	}
}

func (chat *Chat) heartbeat() {
	tc := time.NewTicker(1 * time.Second)
	defer func() {
		chat.heartbeat()
	}()
Loop:
	for {
		select {
		case <-chat.ctx.Done():
			tc.Stop()
		case <-tc.C:
			pl := &pb.HeartbeatMsg{
				Head: &pb.HeartbeatMsgHead{},
			}
			pd, err := proto.Marshal(pl)
			if err != nil {
				return
			}
			err = chat.conn.send(pb.CmdType_HEARTBEAT, pd)
			if err != nil {
				goto Loop
			}
		}
	}
}

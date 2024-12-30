package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/hoysics/basic-im/api/common/v1"
	"github.com/hoysics/basic-im/app/state/internal/biz"
	"github.com/hoysics/basic-im/pkg/cache"
	"github.com/hoysics/basic-im/pkg/router"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
	"sync"
)

/**
处理连接状态与本地缓存和Redis之间的交互
*/

var _rdb *cache.RedisClient
var _stateCache *stateCache
var _ biz.ConnStateRepo = &stateCache{}

type stateCache struct {
	log              *log.Helper
	connToStateTable sync.Map
}

// NewConnStateRepo .
func NewConnStateRepo(logger log.Logger, _ *Data) biz.ConnStateRepo {
	_stateCache = &stateCache{
		log:              log.NewHelper(logger),
		connToStateTable: sync.Map{},
	}
	_ = _stateCache.initLoginSlot(context.Background())
	return _stateCache
}

// 初始化并连接登录槽
func (c *stateCache) initLoginSlot(ctx context.Context) error {
	for _, slot := range _connSlotList {
		loginSlotKey := fmt.Sprintf(LoginSlotSetKey, slot)
		go func() {
			loginSlot, err := _rdb.SMembersStrSlice(ctx, loginSlotKey)
			if err != nil {
				c.log.Fatalf("initLoginSlot rdb.SMembersStrSlice err: %v", err)
			}
			for _, meta := range loginSlot {
				did, connId, err := c.loginSlotUnmarshal(meta)
				if err != nil {
					c.log.Errorf("initLoginSlot Unmarshal err: %v", err)
					continue
				}
				if err := c.connReLogin(ctx, did, connId); err != nil {
					c.log.Errorf("initLoginSlot connReLogin error: %v", err)
				}
			}
		}()
	}
	return nil
}

func (c *stateCache) newConnState(did, connId uint64) *connState {
	state := &connState{
		connId: connId,
		did:    did,
	}
	state.resetHeartTimer()
	return state
}

/*
处理连接登录/登出/重连状态
*/

// ConnLogin 连接登录
func (c *stateCache) ConnLogin(ctx context.Context, did, connId uint64) error {
	state := c.newConnState(did, connId)
	// 登录槽存储
	slotKey := c.getLoginSlotKey(connId)
	meta := c.loginSlotMarshal(did, connId)
	if err := _rdb.SAdd(ctx, slotKey, meta); err != nil {
		return err
	}
	// 添加路由记录
	err := router.AddRecord(ctx, did, _stateServerAddr, connId)
	if err != nil {
		return err
	}
	//TODO 上行消息 max_client_id 初始化, 现在相当于生命周期在conn维度，后面重构sdk时会调整到会话维度
	// 本地状态存储
	c.storeConnIdState(connId, state)
	return nil
}

// 处理服务重启后 通过缓存中的信息重新加载连接数
func (c *stateCache) connReLogin(ctx context.Context, did, connId uint64) error {
	//TODO 有问题 没有继承之前的消息状态
	state := c.newConnState(did, connId)
	c.storeConnIdState(connId, state)
	return state.loadMsgTimer(ctx)
}

func (c *stateCache) ConnLogout(ctx context.Context, connId uint64, callByGw bool) (uint64, error) {
	if state, ok := c.loadConnIdState(connId); ok {
		return state.did, state.close(ctx, callByGw)
	}
	return 0, nil
}

// ReConn 客户端发起重连时调用
func (c *stateCache) ReConn(ctx context.Context, oldConnId, newConnId uint64) error {
	var (
		did uint64
		err error
	)
	if did, err = c.ConnLogout(ctx, oldConnId, false); err != nil {
		return err
	}
	return c.ConnLogin(ctx, did, newConnId)
}

func (c *stateCache) ResetHeartTimer(connId uint64) {
	if state, ok := c.loadConnIdState(connId); ok {
		state.resetHeartTimer()
	}
}

/*
处理本地连接状态缓存
*/
func (c *stateCache) checkConnIdState(connId uint64) bool {
	_, ok := c.connToStateTable.Load(connId)
	return ok
}
func (c *stateCache) loadConnIdState(connId uint64) (*connState, bool) {
	if v, ok := c.connToStateTable.Load(connId); ok {
		state, ok := v.(*connState)
		return state, ok
	}
	return nil, false
}

func (c *stateCache) storeConnIdState(connId uint64, state *connState) {
	c.connToStateTable.Store(connId, state)
}

func (c *stateCache) deleteConnIdState(connId uint64) {
	c.connToStateTable.Delete(connId)
}

/*
*
处理connId与Redis槽位映射
*/
func (c *stateCache) getLoginSlotKey(connId uint64) string {
	slot := connId % uint64(len(_connSlotList))
	slotKey := fmt.Sprintf(LoginSlotSetKey, _connSlotList[slot])
	return slotKey
}

func (c *stateCache) getConnStateSlot(connId uint64) uint64 {
	return connId % uint64(len(_connSlotList))
}

func (c *stateCache) loginSlotUnmarshal(mate string) (uint64, uint64, error) {
	sts := strings.Split(mate, "|")
	if len(sts) < 2 {
		return 0, 0, fmt.Errorf("invalid login slot format: %s", mate)
	}
	did, err := strconv.ParseUint(sts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	connId, err := strconv.ParseUint(sts[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return did, connId, nil
}

func (c *stateCache) loginSlotMarshal(did, connId uint64) string {
	return fmt.Sprintf("%d|%d", did, connId)
}

/*
*
处理最新一条推送消息 由RPC接口层调用
*/

// CompareAndIncrClientId 比对ClientId是否符合+1的规则
func (c *stateCache) CompareAndIncrClientId(ctx context.Context, connId, oldMaxClientId uint64, sessionId string) bool {
	//TODO Call业务层的RPC // 调用下游业务层rpc，只有当rpc回复成功后才能更新max_clientID
	//TODO 此处先假设RPC调用成功
	slotKey := c.getConnStateSlot(connId)
	key := getMaxClientIdKey(slotKey, connId, sessionId)
	c.log.Infof("RunLuaInt %s, %d, %d", key, oldMaxClientId, TTL7D)
	var (
		res int
		err error
	)
	if res, err = RunLuaInt(ctx, LuaCompareAndIncrClientId, []string{key}, oldMaxClientId, TTL7D); err != nil {
		c.log.Errorf("RunLuaInt %s, %d, %v", key, oldMaxClientId, err)
	}
	return res > 0
}

func (c *stateCache) appendLastMsg(ctx context.Context, connId uint64, msg *v1.PushMsg) error {
	if msg == nil {
		return errors.New("nil message")
	}
	var (
		state *connState
		ok    bool
	)
	if state, ok = c.loadConnIdState(connId); !ok {
		return errors.New("connID state is nil")
	}
	slot := _stateCache.getConnStateSlot(connId)
	key := getLastMsgKey(slot, connId)
	msgTimerLock := getMsgTimerLockKey(msg.SessionId, msg.MsgId)
	msgData, _ := proto.Marshal(msg)
	return state.appendMsg(ctx, key, msgTimerLock, msgData)
}

func (c *stateCache) AckLastMsg(ctx context.Context, connId, sessionId, msgId uint64) {
	if state, ok := c.loadConnIdState(connId); ok {
		state.ackLastMsg(ctx, sessionId, msgId)
	}
}

func (c *stateCache) Push(ctx context.Context, connId uint64, pushMsg *v1.PushMsg) error {
	by, err := proto.Marshal(pushMsg)
	if err != nil {
		return err
	}
	//Step 1: Push by gw client
	_gatewayRepo.Push(ctx, &biz.CmdCtx{ConnId: connId, Msg: &v1.MajorMsg{Cmd: v1.CmdType_PUSH, Payload: by}})
	//1. State中针对ConnState仅保存一个msgTimer，针对当前连接上最新的Push消息，仅对这个消息进行超时重试
	//2. 当有新消息Push过来时，取消之前的定时器，staet内仅保留最新push的msg
	//3. 客户端仅维护对最新消息的Ack，其他消息不用回复
	//4. 客户端若出现消息漏洞，则直接走兜底的pull模式拉取数据
	return c.appendLastMsg(ctx, connId, pushMsg)
}

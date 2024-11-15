package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/hoysics/basic-im/api/common/v1"
	"github.com/hoysics/basic-im/app/state/internal/biz"
	"github.com/hoysics/basic-im/pkg/router"
	"github.com/hoysics/basic-im/pkg/timingwheel"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

type connState struct {
	sync.RWMutex
	heartTimer   *timingwheel.Timer //心跳计时器
	reConnTimer  *timingwheel.Timer //重连计时器
	msgTimer     *timingwheel.Timer //消息重发计时器
	msgTimerLock string
	connId       uint64
	did          uint64
}

func (c *connState) close(ctx context.Context, callByGw bool) error {
	c.Lock()
	defer c.Unlock()
	if c.heartTimer != nil {
		c.heartTimer.Stop()
	}
	if c.reConnTimer != nil {
		c.reConnTimer.Stop()
	}
	if c.msgTimer != nil {
		c.msgTimer.Stop()
	}
	// TODO 这里如何保证事务性，值得思考一下，或者说有没有必要保证
	// TODO 这里也可以使用lua或者pipeline 来尽可能合并两次redis的操作 通常在大规模的应用中这是有效的
	// TODO 这里是要好好思考一下，网络调用次数的时间&空间复杂度的
	//1. 清理登录状态
	slotKey := _stateCache.getLoginSlotKey(c.connId)
	meta := _stateCache.loginSlotMarshal(c.did, c.connId)
	if err := _rdb.SRem(ctx, slotKey, meta); err != nil {
		return err
	}
	//2. 清理端侧单次连接最大的ClientId
	slot := _stateCache.getConnStateSlot(c.connId)
	key := getMaxClientIdKey(slot, c.connId, "*")
	// 使用通配符模式查找匹配的键
	keys, err := _rdb.GetKeys(ctx, key)
	if err != nil {
		return err
	}
	// 删除匹配的键
	if len(keys) > 0 {
		if err = _rdb.Del(ctx, keys...); err != nil {
			return err
		}
	}
	//3. 删除路有记录
	if err := router.DelRecord(ctx, c.did); err != nil {
		return err
	}
	//4. 清理push的最后一条消息
	if err := _rdb.Del(ctx, getLastMsgKey(slot, c.connId)); err != nil {
		return err
	}
	//5. 通知网关断开连接
	if !callByGw { //如果时网关发起的删除连接，则不回环调用网关接口
		if err := _gatewayRepo.DelConn(ctx, c.connId); err != nil {
			return err
		}
	}
	//6. 删除本地缓存
	_stateCache.deleteConnIdState(c.connId)
	return nil
}

func (c *connState) appendMsg(ctx context.Context, key, msgTimerLock string, msgData []byte) error {
	c.Lock()
	defer c.Unlock()
	c.msgTimerLock = msgTimerLock
	if c.msgTimer != nil {
		c.msgTimer.Stop()
		c.msgTimer = nil
	}
	c.msgTimer = AfterFunc(100*time.Millisecond, func() {
		if err := c.rePush(ctx); err != nil {
			log.Errorf(`appendMsg repush err: %v`, err)
		}
	})
	if err := _rdb.SetBytes(ctx, key, msgData, TTL7D); err != nil {
		log.Fatalf(`appendMsg err: %v`, err)
	}
	return nil
}

func (c *connState) resetMsgTimer(sessionId, msgId uint64) {
	//TODO 暂时屏蔽connId 不确定在高并发场景下是否会出现数据错乱的情况 理论上应该不会
	//1. 本地计时器重置
	c.Lock()
	defer c.Unlock()
	if c.msgTimer != nil {
		c.msgTimer.Stop()
	}
	c.msgTimerLock = getMsgTimerLockKey(sessionId, msgId)
	//2. redis缓存重置
	c.msgTimer = AfterFunc(100*time.Millisecond, func() {
		if err := c.rePush(context.Background()); err != nil {
			log.Errorf(`resetMsgTimer repush err: %v`, err)
		}
	})
}

func (c *connState) loadMsgTimer(ctx context.Context) error {
	lastMsg, err := c.getLastMsg(ctx)
	if err != nil {
		// 这里的处理是粗暴的，如果线上是需要更solid的方案
		return err
	}
	c.resetMsgTimer(lastMsg.SessionId, lastMsg.MsgId)
	return nil
}

func (c *connState) resetHeartTimer() {
	c.Lock()
	defer c.Unlock()
	if c.heartTimer != nil {
		c.heartTimer.Stop()
	}
	c.heartTimer = AfterFunc(5*time.Second, func() {
		c.resetReConnTimer()
	})
}

// 为了快速重连 不立即释放连接状态 而是分配一个新的计时器 给予10s的延迟
func (c *connState) resetReConnTimer() {
	c.Lock()
	defer c.Unlock()
	if c.reConnTimer != nil {
		c.reConnTimer.Stop()
	}
	//此处不Stop之前的是否会有内存泄露问题=>不会 到期自己执行 执行完了就没了
	c.reConnTimer = AfterFunc(10*time.Second, func() {
		ctx := context.TODO()
		//因为是延迟清除的 则状态可能被新的connId和endpoint继承 要确保清除的是就旧的失效连接 所以需要按照创建AfterFunc时的参数进行调用
		//整体connId状态牵出
		logout, err := _stateCache.ConnLogout(ctx, c.connId, false)
		if err != nil {
			log.Errorf(`conn(%v) logout fail: %v`, logout, err)
		}
	})
}

func (c *connState) getLastMsg(ctx context.Context) (*v1.PushMsg, error) {
	slot := _stateCache.getConnStateSlot(c.connId)
	data, err := _rdb.GetBytes(ctx, getLastMsgKey(slot, c.connId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New(`last msg data is empty`)
	}
	pushMsg := &v1.PushMsg{}
	if err = proto.Unmarshal(data, pushMsg); err != nil {
		return nil, err
	}
	return pushMsg, nil
}

func (c *connState) ackLastMsg(ctx context.Context, sessionId, msgId uint64) bool {
	c.Lock()
	defer c.Unlock()
	msgTimerLock := getMsgTimerLockKey(sessionId, msgId)
	if msgTimerLock != c.msgTimerLock {
		return false
	}
	slot := _stateCache.getConnStateSlot(c.connId)
	key := getLastMsgKey(slot, c.connId)
	if err := _rdb.Del(ctx, key); err != nil {
		return false
	}
	if c.msgTimer != nil {
		c.msgTimer.Stop()
	}
	return true
}

// 仅重推最后一条消息
func (c *connState) rePush(ctx context.Context) error {
	lastMsg, err := c.getLastMsg(ctx)
	if err != nil {
		return err
	}
	if lastMsg == nil {
		return nil
	}
	msgData, err := proto.Marshal(lastMsg)
	if err != nil {
		return err
	}
	_gatewayRepo.Push(ctx, &biz.CmdCtx{ConnId: c.connId, Msg: &v1.MajorMsg{
		Cmd:     v1.CmdType_PUSH,
		Payload: msgData,
	}})
	if _stateCache.checkConnIdState(c.connId) {
		c.resetMsgTimer(lastMsg.SessionId, lastMsg.MsgId)
	} else {
		log.Infof(`此ConnId已经断开连接，不再重发: %v`, c.connId)
	}
	return nil
}

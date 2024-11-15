package data

import (
	"fmt"
	"time"
)

const (
	MaxClientIdKey  = "max_client_id_{%d}_%d_%s"
	LastMsgKey      = "last_msg_{%d}_%d"
	LoginSlotSetKey = "login_slot_set_{%d}" // 通过 hash tag保证在cluster模式上 key都在一个shard上
	TTL7D           = 7 * 24 * time.Hour
)

func getLastMsgKey(slot, connId uint64) string {
	return fmt.Sprintf(LastMsgKey, slot, connId)
}

func getMaxClientIdKey(slot, connId uint64, sessionId string) string {
	return fmt.Sprintf(MaxClientIdKey, slot, connId, sessionId)
}

func getMsgTimerLockKey(sessionId, msgId uint64) string {
	return fmt.Sprintf("%d_%d", sessionId, msgId)
}

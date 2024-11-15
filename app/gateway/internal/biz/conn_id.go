package biz

import (
	"errors"
	"sync"
	"time"
)

const (
	version      = uint64(0)
	sequenceBits = uint64(16)

	maxSequence = int64(-1) ^ (int64(-1) << sequenceBits) //计算一个最大序列值
	timeLeft    = uint8(16)                               // timeLeft = sequenceBits // 时间戳向左偏移量
	versionLeft = uint8(63)                               // 左移动到最高位
	// 2020-05-20 08:00:00 +0800 CST
	twepoch = int64(1589923200000) //常量时间戳(毫秒)
)

var node *ConnIdGenerator

func init() {
	node = &ConnIdGenerator{}
}

type ConnIdGenerator struct {
	mu        sync.Mutex
	LastStamp int64 //记录上一次ID的时间戳
	Sequence  int64 //当前毫秒已经生成的ID序列号（从0开始累加）1毫秒内最多生成2^16个ID
}

// NextId 这里的锁会自旋，不会多么影响性能，主要是临界区小
func (w *ConnIdGenerator) NextId() (uint64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.nextId()
}

func (w *ConnIdGenerator) nextId() (uint64, error) {
	timestamp := time.Now().UnixMilli()
	if timestamp < w.LastStamp {
		return 0, errors.New("time is moving backwards,waiting until")
	}
	if timestamp == w.LastStamp {
		//在同一个毫秒内，通过增加 Sequence（序列号）来生成不同的 ID。
		//maxSequence 是序列号的最大值，通常是一个位掩码，用来限制序列号的位数。
		//序列号增加后与 maxSequence 进行按位与操作，确保序列号不会超过最大值。
		w.Sequence = (w.Sequence + 1) & maxSequence
		if w.Sequence == 0 { // 如果序列号增加后变为 0，表示在当前毫秒内已经生成了最大数量的 ID。为了防止 ID 重复，代码会等待到下一个毫秒。
			for timestamp <= w.LastStamp {
				timestamp = time.Now().UnixMilli()
			}
		}
	} else { // 如果与上次分配的时间戳不等，则为了防止可能的时钟飘移现象，就必须重新计数
		w.Sequence = 0
	}
	w.LastStamp = timestamp
	//减法可以压缩一下时间戳
	//计算 ID 的主要部分。
	//首先，将当前时间戳减去某个起始时间戳 twepoch（通常是一个固定的值，如 Unix 纪元），
	//然后左移 timeLeft 位，以确保时间戳占据 ID 的大部分位数。
	//接着，将序列号与左移后的时间戳进行按位或操作，组合成初步的 ID。
	id := ((timestamp - twepoch) << timeLeft) | w.Sequence
	connId := uint64(id) | (version << versionLeft)
	return connId, nil
}

package biz

import (
	"context"
	"sort"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
)

type SourceRepo interface {
	Watch() <-chan *Event
}

type Dispatcher struct {
	//repo SourceRepo
	log            *log.Helper
	candidateTable map[string]*Endpoint
	sync.RWMutex
}

func NewDispatcher(repo SourceRepo, logger log.Logger) *Dispatcher {
	dp := &Dispatcher{
		log:            log.NewHelper(logger),
		candidateTable: make(map[string]*Endpoint),
	}
	go func() {
		for event := range repo.Watch() {
			switch event.Type {
			case AddNodeEvent:
				dp.AddNode(event)
			case DelNodeEvent:
				dp.DelNode(event)
			}
		}
	}()
	return dp
}

func (dp *Dispatcher) GetIpList(ctx context.Context) ([]*Endpoint, error) {
	//1. 获取候选的Endpoint
	eps := dp.getCandidateEndpoint()
	//2. 逐一计算得分
	for _, ep := range eps {
		ep.CalculateScore()
	}
	//3. 全局排序 动静结合的排序策略
	sort.Slice(eps, func(i, j int) bool {
		// 优先基于活跃分数进行排序
		if eps[i].ActiveScore > eps[j].ActiveScore {
			return true
		}
		// 如果活跃分数相同，则使用静态分数排序
		if eps[i].ActiveScore == eps[j].ActiveScore {
			if eps[i].StaticScore > eps[j].StaticScore {
				return true
			}
			return false
		}
		return false
	})
	return nil, nil
}

func (dp *Dispatcher) getCandidateEndpoint() []*Endpoint {
	dp.RLock()
	defer dp.RUnlock()
	candidates := make([]*Endpoint, 0, len(dp.candidateTable))
	for _, endpoint := range dp.candidateTable {
		candidates = append(candidates, endpoint)
	}
	return candidates
}

func (dp *Dispatcher) AddNode(event *Event) {
	dp.Lock()
	defer dp.Unlock()
	var (
		ep *Endpoint
		ok bool
	)
	if ep, ok = dp.candidateTable[event.IP]; !ok {
		ep = NewEndpoint(event.IP, event.Port)
		dp.candidateTable[event.Key()] = ep
	}
	ep.UpdateState(&State{
		ConnectNum:   event.ConnectNum,
		MessageBytes: event.MessageBytes,
	})
}

func (dp *Dispatcher) DelNode(event *Event) {
	dp.Lock()
	defer dp.Unlock()
	delete(dp.candidateTable, event.Key())
}

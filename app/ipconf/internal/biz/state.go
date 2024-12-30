package biz

import "math"

// State 不同时期加入的物理机配置不同->已使用资源无法反映出实际的负载差距->使用其自身所剩余的资源指标
type State struct {
	ConnectNum   float64 // 业务上 IM Gateway 总体持有的长连接数量 的剩余值
	MessageBytes float64 // 业务上 IM Gateway 每秒收发消息的总字节数 的剩余值
}

func (s *State) Add(st *State) {
	if st == nil {
		return
	}
	s.ConnectNum += st.ConnectNum
	s.MessageBytes += st.MessageBytes
}

func (s *State) Sub(st *State) {
	if st == nil {
		return
	}
	s.ConnectNum -= st.ConnectNum
	s.MessageBytes -= st.MessageBytes
}

func (s *State) Clone() State {
	return State{
		ConnectNum:   s.ConnectNum,
		MessageBytes: s.MessageBytes,
	}
}

// CalculateActiveScore 此处的假设或者说前提是 网络带宽将是系统的瓶颈所在->哪台机器富裕的带宽资源多，哪台机器的负载就是最新的
func (s *State) CalculateActiveScore() float64 {
	return getGB(s.MessageBytes)
}

func (s *State) CalculateStaticScore() float64 {
	return s.ConnectNum
}

func (s *State) Avg(num float64) {
	s.ConnectNum /= num
	s.MessageBytes /= num
}

func getGB(m float64) float64 {
	return decimal(m / (1 << 30))
}

func decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

func min(a, b, c float64) float64 {
	m := func(k, j float64) float64 {
		if k > j {
			return j
		}
		return k
	}
	return m(a, m(b, c))
}

package biz

const (
	_windowSize = 5
)

type StateWindow struct {
	stateQueue []*State
	stateChan  chan *State
	sumState   *State
	idx        int64
}

func newStateWindow() *StateWindow {
	return &StateWindow{
		stateQueue: make([]*State, _windowSize),
		stateChan:  make(chan *State),
		sumState:   &State{},
	}
}

func (w *StateWindow) getState() *State {
	res := w.sumState.Clone()
	res.Avg(_windowSize)
	return &res
}

func (w *StateWindow) appendState(state *State) {
	curIdx := w.idx % _windowSize
	// 减去即将被删除的state
	w.sumState.Sub(w.stateQueue[curIdx])
	// 更新最新的state
	w.stateQueue[curIdx] = state
	// 计算最新的窗口和
	w.sumState.Add(state)
	w.idx++
}

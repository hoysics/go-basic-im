package biz

import (
	"sync/atomic"
	"unsafe"
)

type Endpoint struct {
	IP          string
	Port        string
	ActiveScore float64
	StaticScore float64
	State       *State
	window      *StateWindow
}

func NewEndpoint(ip string, port string) *Endpoint {
	ed := &Endpoint{IP: ip, Port: port}
	ed.window = newStateWindow()
	ed.State = ed.window.getState()
	go func() {
		for state := range ed.window.stateChan {
			ed.window.appendState(state)
			newState := ed.window.getState()
			atomic.SwapPointer((*unsafe.Pointer)((unsafe.Pointer)(ed.State)), unsafe.Pointer(newState))
		}
	}()
	return ed
}

func (e *Endpoint) UpdateState(s *State) {
	e.window.stateChan <- s
}

func (e *Endpoint) CalculateScore() {
	if e.State != nil {
		e.ActiveScore = e.State.CalculateActiveScore()
		e.StaticScore = e.State.CalculateStaticScore()
	}
}

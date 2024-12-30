package data

import (
	"github.com/hoysics/basic-im/pkg/timingwheel"
	"time"
)

var _wheel *timingwheel.TimingWheel

func InitTimer() {
	_wheel = timingwheel.NewTimingWheel(time.Millisecond, 20)
	_wheel.Start()
}
func CloseTimer() {
	_wheel.Stop()
}

func AfterFunc(d time.Duration, f func()) *timingwheel.Timer {
	t := _wheel.AfterFunc(d, f)
	return t
}

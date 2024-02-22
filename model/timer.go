package model

import (
	"time"
)

type Timer struct {
	startTime time.Time
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) Stop() int64 {
	runningTime := time.Since(t.startTime)
	return int64(runningTime.Seconds())
}

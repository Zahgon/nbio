// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package timer

import (
	"math"
	"sync"
	"time"
)

const (
	TimeForever = time.Duration(math.MaxInt64)
)

type Timer struct {
	name      string
	asyncMux  sync.Mutex
	asyncList []func()
}

//go:norace
func New(name string) *Timer { _ = "STUB: not implemented"; return nil }

// IsTimerRunning .
//
//go:norace
func (t *Timer) IsTimerRunning() bool {
	_ = "STUB: not implemented"

	// Start .
	//
	//go:norace
	return false
}

func (t *Timer) Start() {
	_ = "STUB: not implemented"

	// Stop .
	//
	//go:norace
	return
}

func (t *Timer) Stop() {
	_ = "STUB: not implemented"

	// After used as time.After.
	//
	//go:norace
	return
}

func (t *Timer) After(d time.Duration) <-chan time.Time { _ = "STUB: not implemented"; return nil }

// AfterFunc used as time.AfterFunc.
//
//go:norace
func (t *Timer) AfterFunc(timeout time.Duration, f func()) *time.Timer {
	_ = "STUB: not implemented"
	return nil
}

// Async executes f in another goroutine.
//
//go:norace
func (t *Timer) Async(f func()) { _ = "STUB: not implemented"; return }

// func (t *Timer) Async(f func()) {

// 	go func() {
// 		defer func() {
// 			err := recover()
// 			if err != nil {
// 				const size = 64 << 10
// 				buf := make([]byte, size)
// 				buf = buf[:runtime.Stack(buf, false)]
// 				logging.Error("Timer[%v] exec call failed: %v\n%v\n", t.name, err, *(*string)(unsafe.Pointer(&buf)))
// 			}
// 		}()
// 		f()
// 	}()
// }

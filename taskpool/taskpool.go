// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package taskpool

// TaskPool .
type TaskPool struct {
	concurrent    int64
	maxConcurrent int64
	chQqueue      chan func()
	chClose       chan struct{}
	caller        func(f func())
}

// fork .
//
//go:norace
func (tp *TaskPool) fork(f func()) bool { _ = "STUB: not implemented"; return false }

// Call .
//
//go:norace
func (tp *TaskPool) Call(f func()) {
	_ = "STUB: not implemented"

	// Go .
	//
	//go:norace
	return
}

func (tp *TaskPool) Go(f func()) {
	_ = "STUB: not implemented"
	// If current goroutine num is less than maxConcurrent,
	// creat a new goroutine to exec new task.
	return
}

// Else push the new task into chan/queue.

// Stop .
//
//go:norace
func (tp *TaskPool) Stop() { _ = "STUB: not implemented"; return }

// New creates and returns a TaskPool.
//
//go:norace
func New(maxConcurrent int, chQqueueSize int, v ...interface{}) *TaskPool {
	_ = "STUB: not implemented"
	return nil
}

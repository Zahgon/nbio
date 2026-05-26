package taskpool

import (
	"sync"
)

// IOTaskPool  .
type IOTaskPool struct {
	task *TaskPool
	pool sync.Pool
}

// Call .
//
//go:norace
func (tp *IOTaskPool) Call(f func(*[]byte)) { _ = "STUB: not implemented"; return }

// Go .
//
//go:norace
func (tp *IOTaskPool) Go(f func(*[]byte)) { _ = "STUB: not implemented"; return }

// Stop .
//
//go:norace
func (tp *IOTaskPool) Stop() {
	_ = "STUB: not implemented"

	// NewIO creates and returns a IOTaskPool.
	//
	//go:norace
	return
}

func NewIO(concurrent, queueSize, bufSize int, v ...interface{}) *IOTaskPool {
	_ = "STUB: not implemented"
	return nil
}

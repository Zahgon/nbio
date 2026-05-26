package mempool

import (
	"sync"
)

type sizeMap struct {
	MallocCount int64 `json:"MallocCount"`
	FreeCount   int64 `json:"FreeCount"`
	NeedFree    int64 `json:"NeedFree"`
}

type debugger struct {
	mux         sync.Mutex
	on          bool
	MallocCount int64            `json:"MallocCount"`
	FreeCount   int64            `json:"FreeCount"`
	NeedFree    int64            `json:"NeedFree"`
	SizeMap     map[int]*sizeMap `json:"SizeMap"`
}

//go:norace
func (d *debugger) SetDebug(dbg bool) {
	_ = "STUB: not implemented"

	//go:norace
	return
}

func (d *debugger) incrMalloc(pbuf *[]byte) { _ = "STUB: not implemented"; return }

//go:norace
func (d *debugger) incrMallocSlow(pbuf *[]byte) { _ = "STUB: not implemented"; return }

//go:norace
func (d *debugger) incrFree(pbuf *[]byte) { _ = "STUB: not implemented"; return }

//go:norace
func (d *debugger) incrFreeSlow(pbuf *[]byte) { _ = "STUB: not implemented"; return }

//go:norace
func (d *debugger) String() string { _ = "STUB: not implemented"; return "" }

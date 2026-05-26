package mempool

import (
	"bytes"
	"sync"
)

var (
	stackMux    = sync.Mutex{}
	stackBuf    = [1024 * 64]byte{}
	stackWriter = bytes.NewBuffer(stackBuf[:0])
	stackMap    = map[string][2]uintptr{}
	nilStackPtr = [2]uintptr{0, 0}
)

type TraceDebugger struct {
	mux       sync.Mutex
	pAlloced  map[uintptr][2]uintptr
	allocator Allocator
}

//go:norace
func NewTraceDebuger(allocator Allocator) *TraceDebugger { _ = "STUB: not implemented"; return nil }

// Malloc .
//
//go:norace
func (td *TraceDebugger) Malloc(size int) *[]byte { _ = "STUB: not implemented"; return nil }

// deprecated.
//
//go:norace
func (td *TraceDebugger) Realloc(pbuf *[]byte, size int) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// Append .
//
//go:norace
func (td *TraceDebugger) Append(pbuf *[]byte, more ...byte) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// AppendString .
//
//go:norace
func (td *TraceDebugger) AppendString(pbuf *[]byte, more string) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// Free .
//
//go:norace
func (td *TraceDebugger) Free(pbuf *[]byte) { _ = "STUB: not implemented"; return }

//go:norace
func (td *TraceDebugger) setBufferPointer(ptr uintptr) { _ = "STUB: not implemented"; return }

//go:norace
func (td *TraceDebugger) deleteBufferPointer(ptr uintptr) { _ = "STUB: not implemented"; return }

// func getStack() string {
// 	stackMux.Lock()
// 	defer stackMux.Unlock()
// 	buf := stackBuf[:runtime.Stack(stackBuf, false)]
// 	return string(buf)
// }

//go:norace
func getStackAndPtr() (string, [2]uintptr) { _ = "STUB: not implemented"; return "", nil }

//go:norace
func ptr2StackString(ptr [2]uintptr) string { _ = "STUB: not implemented"; return "" }

// func bytesToStr(b []byte) string {
// 	return *(*string)(unsafe.Pointer(&b))
// }

// func strToBytes(s string) []byte {
// 	x := (*[2]uintptr)(unsafe.Pointer(&s))
// 	h := [3]uintptr{x[0], x[1], x[1]}
// 	return *(*[]byte)(unsafe.Pointer(&h))
// }

//go:norace
func printStack(info string, preStackPtr [2]uintptr) { _ = "STUB: not implemented"; return }

// os.Exit(-1)

//go:norace
func bytesPointer(pbuf *[]byte) uintptr { _ = "STUB: not implemented"; return 0 }

// func stringPointer(s *string) uintptr {
// 	ptr := (*uintptr)(unsafe.Pointer(s))
// 	return (uintptr)(unsafe.Pointer(&ptr))
// }

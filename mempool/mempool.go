// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mempool

import (
	"sync"
)

// MemPool .
type MemPool struct {
	*debugger
	bufSize  int
	freeSize int
	pool     *sync.Pool
}

// New .
func New(bufSize, freeSize int) Allocator { _ = "STUB: not implemented"; return *new(Allocator) }

// Debug:       true,

// Malloc .
func (mp *MemPool) Malloc(size int) *[]byte { _ = "STUB: not implemented"; return nil }

// Realloc .
func (mp *MemPool) Realloc(pbuf *[]byte, size int) *[]byte { _ = "STUB: not implemented"; return nil }

// Append .
func (mp *MemPool) Append(pbuf *[]byte, more ...byte) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// AppendString .
func (mp *MemPool) AppendString(pbuf *[]byte, more string) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// Free .
func (mp *MemPool) Free(pbuf *[]byte) { _ = "STUB: not implemented"; return }

package mempool

import (
	"sync"
)

var (
	alignedPools   [alignedPoolBucketNum]sync.Pool
	alignedIndexes [maxAlignedBufferSize + 1]byte
)

const (
	minAlignedBufferSizeBits = 5
	maxAlignedBufferSizeBits = 15
	minAlignedBufferSize     = 1 << minAlignedBufferSizeBits                           // 32
	minAlignedBufferSizeMask = minAlignedBufferSize - 1                                // 31
	maxAlignedBufferSize     = 1 << maxAlignedBufferSizeBits                           // 32k
	alignedPoolBucketNum     = maxAlignedBufferSizeBits - minAlignedBufferSizeBits + 1 // 12
)

//go:norace
func init() {
	var poolSizes [alignedPoolBucketNum]int
	for i := range alignedPools {
		size := 1 << (i + minAlignedBufferSizeBits)
		poolSizes[i] = size
		alignedPools[i].New = func() interface{} {
			b := make([]byte, size)
			return &b
		}
	}

	getPoolBySize := func(size int) byte {
		for i, n := range poolSizes {
			if size <= n {
				return byte(i)
			}
		}
		return 0xFF
	}

	for i := range alignedIndexes {
		alignedIndexes[i] = getPoolBySize(i)
	}
}

// NewAligned .
//
//go:norace
func NewAligned() Allocator { _ = "STUB: not implemented"; return *new(Allocator) }

// AlignedAllocator .
type AlignedAllocator struct {
	*debugger
}

// Malloc .
//
//go:norace
func (amp *AlignedAllocator) Malloc(size int) *[]byte { _ = "STUB: not implemented"; return nil }

// Realloc .
//
//go:norace
func (amp *AlignedAllocator) Realloc(pbuf *[]byte, size int) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// Append .
//
//go:norace
func (amp *AlignedAllocator) Append(pbuf *[]byte, more ...byte) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// AppendString .
//
//go:norace
func (amp *AlignedAllocator) AppendString(pbuf *[]byte, s string) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// Free .
//
//go:norace
func (amp *AlignedAllocator) Free(pbuf *[]byte) { _ = "STUB: not implemented"; return }

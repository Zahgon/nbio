package mempool

// stdAllocator .
type stdAllocator struct {
	*debugger
}

// Malloc .
//
//go:norace
func (a *stdAllocator) Malloc(size int) *[]byte { _ = "STUB: not implemented"; return nil }

// Realloc .
//
//go:norace
func (a *stdAllocator) Realloc(pbuf *[]byte, size int) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

// Free .
//
//go:norace
func (a *stdAllocator) Free(pbuf *[]byte) {
	_ = "STUB: not implemented"

	//go:norace
	return
}

func (a *stdAllocator) Append(pbuf *[]byte, more ...byte) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

//go:norace
func (a *stdAllocator) AppendString(pbuf *[]byte, more string) *[]byte {
	_ = "STUB: not implemented"
	return nil
}

//go:norace
func NewSTD() Allocator { _ = "STUB: not implemented"; return *new(Allocator) }

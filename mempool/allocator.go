package mempool

// DefaultMemPool .
var DefaultMemPool = New(1024, 1024*1024*1024)

type Allocator interface {
	Malloc(size int) *[]byte
	Realloc(buf *[]byte, size int) *[]byte // deprecated.
	Append(buf *[]byte, more ...byte) *[]byte
	AppendString(buf *[]byte, more string) *[]byte
	Free(buf *[]byte)
}

type DebugAllocator interface {
	Allocator
	String() string
	SetDebug(bool)
}

//go:norace
func Malloc(size int) *[]byte { _ = "STUB: not implemented"; return nil }

//go:norace
func Realloc(pbuf *[]byte, size int) *[]byte { _ = "STUB: not implemented"; return nil }

//go:norace
func Append(pbuf *[]byte, more ...byte) *[]byte { _ = "STUB: not implemented"; return nil }

//go:norace
func AppendString(pbuf *[]byte, more string) *[]byte { _ = "STUB: not implemented"; return nil }

//go:norace
func Free(pbuf *[]byte) { _ = "STUB: not implemented"; return }

// func Init(bufSize, freeSize int) {
// 	DefaultMemPool = New(bufSize, freeSize)
// }

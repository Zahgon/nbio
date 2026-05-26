package websocket

import (
	"compress/flate"
	"io"
	"sync"
)

const (
	minCompressionLevel     = -2
	maxCompressionLevel     = flate.BestCompression
	defaultCompressionLevel = 1

	flateReaderTail = "\x00\x00\xff\xff" + "\x01\x00\x00\xff\xff"
)

var (
	flateWriterPools [maxCompressionLevel - minCompressionLevel + 1]sync.Pool
	flateReaderPool  = sync.Pool{New: func() interface{} {
		return flate.NewReader(nil)
	}}
)

//go:norace
func isValidCompressionLevel(level int) bool { _ = "STUB: not implemented"; return false }

//go:norace
func decompressReader(r io.Reader) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

type flateReadWrapper struct {
	fr io.ReadCloser
}

//go:norace
func (r *flateReadWrapper) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Preemptively place the reader back in the pool. This helps with
// scenarios where the application does not call NextReader() soon after
// this final read.

//go:norace
func (r *flateReadWrapper) Close() error { _ = "STUB: not implemented"; return nil }

//go:norace
func compressWriter(w io.WriteCloser, level int) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

type truncWriter struct {
	w io.WriteCloser
	n int
	p [4]byte
}

//go:norace
func (w *truncWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type flateWriteWrapper struct {
	fw *flate.Writer
	p  *sync.Pool
}

//go:norace
func (w *flateWriteWrapper) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		//go:norace
		nil
}

func (w *flateWriteWrapper) Close() error { _ = "STUB: not implemented"; return nil }

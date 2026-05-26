// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbhttp

import (
	"sync"
)

var (
	emptyBodyReader = BodyReader{}
	bodyReaderPool  = sync.Pool{
		New: func() interface{} {
			return &BodyReader{}
		},
	}
)

// BodyReader implements io.ReadCloser and is to be used as HTTP body.
type BodyReader struct {
	index   int       // first buffer read index
	left    int       // num of byte left
	buffers []*[]byte // buffers that storage HTTP body
	engine  *Engine   // allocator that manages buffers
	closed  bool
}

// Read reads body bytes to p, returns the num of bytes read and error.
//
//go:norace
func (br *BodyReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close frees buffers and resets itself to empty value.
//
//go:norace
func (br *BodyReader) Close() error { _ = "STUB: not implemented"; return nil }

// *br = emptyBodyReader
// bodyReaderPool.Put(br)

// Index returns current head buffer's reading index.
//
//go:norace
func (br *BodyReader) Index() int {
	_ = "STUB: not implemented"

	// Left returns how many bytes are left for reading.
	//
	//go:norace
	return 0
}

func (br *BodyReader) Left() int {
	_ = "STUB: not implemented"

	// Buffers returns the underlayer buffers that store the HTTP Body.
	//
	//go:norace
	return 0
}

func (br *BodyReader) Buffers() []*[]byte {
	_ = "STUB: not implemented"

	// RawBodyBuffers returns a reference of BodyReader's current buffers.
	// The buffers returned will be closed(released automatically when closed)
	// HTTP Handler is called, users should not free the buffers and should
	// not hold it any longer after the HTTP Handler is called.
	//
	//go:norace
	return nil
}

func (br *BodyReader) RawBodyBuffers() [][]byte { _ = "STUB: not implemented"; return nil }

// Engine returns Engine that creates this HTTP Body.
//
//go:norace
func (br *BodyReader) Engine() *Engine {
	_ = "STUB: not implemented"

	// append appends data to buffers.
	//
	//go:norace
	return nil
}

func (br *BodyReader) append(data []byte) error { _ = "STUB: not implemented"; return nil }

// NewBodyReader creates a BodyReader.
//
//go:norace
func NewBodyReader(engine *Engine) *BodyReader { _ = "STUB: not implemented"; return nil }

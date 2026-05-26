// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbhttp

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

// Response represents the server side of an HTTP response.
type Response struct {
	Parser *Parser

	request *http.Request // request for this response.

	status     string
	statusCode int // status code passed to WriteHeader.

	header      http.Header
	trailer     map[string]string
	trailerSize int

	buffer       *[]byte
	bodyBuffer   *[]byte
	contentLen   int
	bodyWritten  int
	intFormatBuf [10]byte

	chunked      bool
	chunkChecked bool
	headEncoded  bool
	hasBody      bool
	hijacked     bool
}

// Hijack .
//
//go:norace
func (res *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

// Header .
//
//go:norace
func (res *Response) Header() http.Header {
	_ = "STUB: not implemented"

	// WriteHeader .
	//
	//go:norace
	return *new(http.Header)
}

func (res *Response) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

const maxPacketSize = 65536

// WriteString .
//
//go:norace
func (res *Response) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Write .
//
//go:norace
func (res *Response) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Header has been sent, no cached head buffer,
// append the data to body buffer.

// If has header buffer and total size <  maxPacketSize,
// set the header buffer as the body buffer and process
// the new body data.
// Else, send header buffer first, then process the data.

// If "Content-Length" has been set,
// and no cached buffer,
// and the data size >= maxPacketSize,
// send the data directly.

// Prepare a new buffer for caching the data.

// If "Content-Length" has been set,
// has cached buffer, and
// the data total size >= maxPacketSize,
// send the cached buffer first.

// If the new data size >= maxPacketSize,
// send the new data directly.

// Append the data to the body buffer cache.

// writeChunk .
//
//go:norace
func (res *Response) writeChunk(conn net.Conn, data []byte, l int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If total size < maxPacketSize, append the data to the cache buffer,
// then return and wait for new data.

// When total size >= maxPacketSize:
// 1. If has cache buffer, send the cache buffer and length string first.

// Reset the cache buffer.

// 2. Append length string to the new buffer.

// 3. Append data and tail to the buffer and send the buffer.

func (res *Response) contentLength() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFrom .
//
//go:norace
func (res *Response) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Push implements the http.Pusher interface.
// It is not supported by nbhttp, so it always returns an error.
//
//go:norace
func (res *Response) Push(target string, opts *http.PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Flush implements http.Flusher. It sends any buffered data to the client.
//
//go:norace
func (res *Response) Flush() { _ = "STUB: not implemented"; return }

// checkChunked .
//
//go:norace
func (res *Response) checkChunked() { _ = "STUB: not implemented"; return }

// 1. See if chunking is already set

// 2. See if we should fall back to chunking

// Don't chunk for responses that are forbidden from having a body

// 3. See if we need to chunk for trailers

// flush .
//
//go:norace
func (res *Response) eoncodeHead() { _ = "STUB: not implemented"; return }

//go:norace
func (res *Response) flush(conn io.Writer) error { _ = "STUB: not implemented"; return nil }

var numMap = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

//go:norace
func (res *Response) formatInt(n int, base int) string { _ = "STUB: not implemented"; return "" }

// NewResponse .
//
//go:norace
func NewResponse(parser *Parser, request *http.Request) *Response {
	_ = "STUB: not implemented"
	return nil
}

/*"Server": []string{"nbio"}*/

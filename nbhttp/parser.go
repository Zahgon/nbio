// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbhttp

import (
	"net"
	"net/http"
	"sync"
)

const (
	transferEncodingHeader = "Transfer-Encoding"
	trailerHeader          = "Trailer"
	contentLengthHeader    = "Content-Length"

	// MaxUint .
	MaxUint = ^uint(0)
	// MaxInt .
	MaxInt = int64(int(MaxUint >> 1))
)

type ParserCloser interface {
	UnderlayerConn() net.Conn
	Parse(data []byte) error
	CloseAndClean(err error)
}

// Parser .
type Parser struct {
	mux sync.Mutex

	// bytesCached for half packet.
	bytesCached *[]byte

	// errClose error

	onClose func(p *Parser, err error)

	ParserCloser ParserCloser

	Engine *Engine

	// Underlayer Conn.
	Conn net.Conn

	// used to call message handler when got a full Request/Response.
	Execute func(f func()) bool

	Processor Processor

	// http fields
	proto         string
	statusCode    int
	status        string
	headerKey     string
	headerValue   string
	header        http.Header
	trailer       http.Header
	contentLength int
	chunkSize     int

	state        int8
	chunked      bool
	isClient     bool
	headerExists bool
}

//go:norace
func (p *Parser) UnderlayerConn() net.Conn {
	_ = "STUB: not implemented"

	//go:norace
	return *new(net.Conn)
}

func (p *Parser) nextState(state int8) { _ = "STUB: not implemented"; return }

// OnClose registers callback for closing.
//
//go:norace
func (p *Parser) OnClose(h func(p *Parser, err error)) {
	_ = "STUB: not implemented"

	// CloseAndClean closes the underlayer connection and cleans up related.
	//
	//go:norace
	return
}

func (p *Parser) CloseAndClean(err error) { _ = "STUB: not implemented"; return }

// p.errClose = err

// if p.ReadCloser != nil {
// 	p.ReadCloser.CloseWithError(p.errClose)
// }

//go:norace
func parseAndValidateChunkSize(originalStr string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Parse parses data bytes and calls HTTP handler when full request received.
// If the connection is upgraded, it passes the data bytes to the ParserCloser
// and doesn't parse them itself any more.
//
//go:norace
func (p *Parser) Parse(data []byte) error { _ = "STUB: not implemented"; return nil }

// if !isToken(c) {
// 	return ErrInvalidCharInHeader
// }

// chunk size is 0

// read trailer headers

// read tail cr lf

// all trailer header readed

// if !isToken(c) {
// 	return ErrInvalidCharInHeader
// }

// if !isToken(c) {
// 	return ErrInvalidCharInHeader
// }

//go:norace
func (p *Parser) parseTransferEncoding() error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *Parser) parseContentLength() (err error) { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *Parser) parseTrailer() error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *Parser) handleMessage() { _ = "STUB: not implemented"; return }

// NewParser creates an HTTP parser.
//
//go:norace
func NewParser(conn net.Conn, engine *Engine, processor Processor, isClient bool, executor func(f func()) bool) *Parser {
	_ = "STUB: not implemented"
	return nil
}

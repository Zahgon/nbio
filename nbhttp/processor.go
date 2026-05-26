// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbhttp

import (
	"net/http"
	"sync"
)

var (
	// used to reset a http.Request to empty value.
	emptyRequest = http.Request{}
	// used to reset a Response to empty value.
	emptyResponse = Response{}
	// used to reset a http.Response to empty value.
	emptyClientResponse = http.Response{}

	requestPool = sync.Pool{
		New: func() interface{} {
			return &http.Request{}
		},
	}

	responsePool = sync.Pool{
		New: func() interface{} {
			return &Response{}
		},
	}

	clientResponsePool = sync.Pool{
		New: func() interface{} {
			return &http.Response{}
		},
	}
)

//go:norace
func releaseRequest(req *http.Request, retainHTTPBody bool) { _ = "STUB: not implemented"; return }

// do not release the body

// fast gc for fields

//go:norace
func releaseResponse(res *Response) { _ = "STUB: not implemented"; return }

//go:norace
func releaseClientResponse(res *http.Response) { _ = "STUB: not implemented"; return }

// Processor .
type Processor interface {
	OnMethod(parser *Parser, method string)
	OnURL(parser *Parser, uri string) error
	OnProto(parser *Parser, proto string) error
	OnStatus(parser *Parser, code int, status string)
	OnHeader(parser *Parser, key, value string)
	OnContentLength(parser *Parser, contentLength int)
	OnBody(parser *Parser, data []byte) error
	OnTrailerHeader(parser *Parser, key, value string)
	OnComplete(parser *Parser)
	Close(parser *Parser, err error)
	Clean(parser *Parser)
}

var (
	emptyServerProcessor = ServerProcessor{}
	emptyClientProcessor = ClientProcessor{}
)

// ServerProcessor is used for server side connection.
type ServerProcessor struct {
	request *http.Request
}

// OnMethod .
//
//go:norace
func (p *ServerProcessor) OnMethod(parser *Parser, method string) {
	_ = "STUB: not implemented"
	return
}

// OnURL .
//
//go:norace
func (p *ServerProcessor) OnURL(parser *Parser, rawurl string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnProto .
//
//go:norace
func (p *ServerProcessor) OnProto(parser *Parser, proto string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnStatus .
//
//go:norace
func (p *ServerProcessor) OnStatus(parser *Parser, code int, status string) {
	_ = "STUB: not implemented"

	// OnHeader .
	//
	//go:norace
	return
}

func (p *ServerProcessor) OnHeader(parser *Parser, key, value string) {
	_ = "STUB: not implemented"
	return
}

// p.isUpgrade = (key == "Connection" && value == "upgrade")

// OnContentLength .
//
//go:norace
func (p *ServerProcessor) OnContentLength(parser *Parser, contentLength int) {
	_ = "STUB: not implemented"
	return
}

// OnBody .
//
//go:norace
func (p *ServerProcessor) OnBody(parser *Parser, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTrailerHeader .
//
//go:norace
func (p *ServerProcessor) OnTrailerHeader(parser *Parser, key, value string) {
	_ = "STUB: not implemented"
	return
}

// OnComplete .
//
//go:norace
func (p *ServerProcessor) OnComplete(parser *Parser) { _ = "STUB: not implemented"; return }

// http 2.0
// if request.Method == "PRI" && len(request.Header) == 0 && request.URL.Path == "*" && request.Proto == "HTTP/2.0" {
// 	p.isUpgrade = true
// 	p.parser.Upgrader = &Http2Upgrader{}
// 	return
// }

//go:norace
func (p *ServerProcessor) flushResponse(parser *Parser, res *Response) {
	_ = "STUB: not implemented"
	return
}

// the data may still in the send queue

// Clean .
//
//go:norace
func (p *ServerProcessor) Clean(parser *Parser) { _ = "STUB: not implemented"; return }

// Close .
//
//go:norace
func (p *ServerProcessor) Close(parser *Parser, err error) {
	_ = "STUB: not implemented"

	// NewServerProcessor .
	//
	//go:norace
	return
}

func NewServerProcessor() Processor {
	_ = "STUB: not implemented"
	return *

	// ClientProcessor is used for client side connection.
	new(Processor)
}

type ClientProcessor struct {
	conn     *ClientConn
	response *http.Response
	handler  func(res *http.Response, err error)
}

// OnMethod .
//
//go:norace
func (p *ClientProcessor) OnMethod(parser *Parser, method string) {
	_ = "STUB: not implemented"

	// OnURL .
	//
	//go:norace
	return
}

func (p *ClientProcessor) OnURL(parser *Parser, uri string) error {
	_ = "STUB: not implemented"

	// OnProto .
	//
	//go:norace
	return nil
}

func (p *ClientProcessor) OnProto(parser *Parser, proto string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnStatus .
//
//go:norace
func (p *ClientProcessor) OnStatus(parser *Parser, code int, status string) {
	_ = "STUB: not implemented"
	return
}

// OnHeader .
//
//go:norace
func (p *ClientProcessor) OnHeader(parser *Parser, key, value string) {
	_ = "STUB: not implemented"
	return
}

// OnContentLength .
//
//go:norace
func (p *ClientProcessor) OnContentLength(parser *Parser, contentLength int) {
	_ = "STUB: not implemented"
	return
}

// OnBody .
//
//go:norace
func (p *ClientProcessor) OnBody(parser *Parser, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTrailerHeader .
//
//go:norace
func (p *ClientProcessor) OnTrailerHeader(parser *Parser, key, value string) {
	_ = "STUB: not implemented"
	return
}

// OnComplete .
//
//go:norace
func (p *ClientProcessor) OnComplete(parser *Parser) { _ = "STUB: not implemented"; return }

// Fix #225
// Handle upgrade handshake response in the io goroutine to avoid concurrent issue:
// 1. when the server may send a message together with handshake response
// 2. we handle the handshake response in another goroutine
// 3. poller continue reading data using http parser(the upgrader reader hasn't been set before 2)
// then we got parsing errors or panic.

// Clean .
//
//go:norace
func (p *ClientProcessor) Clean(parser *Parser) { _ = "STUB: not implemented"; return }

// Close .
//
//go:norace
func (p *ClientProcessor) Close(parser *Parser, err error) { _ = "STUB: not implemented"; return }

// NewClientProcessor .
//
//go:norace
func NewClientProcessor(conn *ClientConn, handler func(res *http.Response, err error)) Processor {
	_ = "STUB: not implemented"
	return *new(Processor)
}

// EmptyProcessor .
type EmptyProcessor struct{}

// OnMethod .
//
//go:norace
func (p *EmptyProcessor) OnMethod(parser *Parser, method string) {
	_ = "STUB: not implemented"

	// OnURL .
	//
	//go:norace
	return
}

func (p *EmptyProcessor) OnURL(parser *Parser, uri string) error {
	_ = "STUB: not implemented"

	// OnProto .
	//
	//go:norace
	return nil
}

func (p *EmptyProcessor) OnProto(parser *Parser, proto string) error {
	_ = "STUB: not implemented"

	// OnStatus .
	//
	//go:norace
	return nil
}

func (p *EmptyProcessor) OnStatus(parser *Parser, code int, status string) {
	_ = "STUB: not implemented"

	// OnHeader .
	//
	//go:norace
	return
}

func (p *EmptyProcessor) OnHeader(parser *Parser, key, value string) {
	_ = "STUB: not implemented"

	// OnContentLength .
	//
	//go:norace
	return
}

func (p *EmptyProcessor) OnContentLength(parser *Parser, contentLength int) {
	_ = "STUB: not implemented"

	// OnBody .
	//
	//go:norace
	return
}

func (p *EmptyProcessor) OnBody(parser *Parser, data []byte) error {
	_ = "STUB: not implemented"

	// OnTrailerHeader .
	//
	//go:norace
	return nil
}

func (p *EmptyProcessor) OnTrailerHeader(parser *Parser, key, value string) {
	_ = "STUB: not implemented"

	// OnComplete .
	//
	//go:norace
	return
}

func (p *EmptyProcessor) OnComplete(parser *Parser) {
	_ = "STUB: not implemented"

	// Clean .
	//
	//go:norace
	return
}

func (p *EmptyProcessor) Clean(parser *Parser) {
	_ = "STUB: not implemented"

	// Close .
	//
	//go:norace
	return
}

func (p *EmptyProcessor) Close(parser *Parser, err error) {
	_ = "STUB: not implemented"

	// NewEmptyProcessor .
	//
	//go:norace
	return
}

func NewEmptyProcessor() Processor { _ = "STUB: not implemented"; return *new(Processor) }

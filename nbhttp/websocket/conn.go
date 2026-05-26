// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package websocket

import (
	"context"
	"io"
	"net"
	"sync"
	"time"

	"github.com/lesismal/nbio/mempool"
	"github.com/lesismal/nbio/nbhttp"
)

const (
	maxControlFramePayloadSize = 125
)

// MessageType .
type MessageType int8

// The message types are defined in RFC 6455, section 11.8.t .
const (
	// FragmentMessage .
	FragmentMessage MessageType = 0 // Must be preceded by Text or Binary message
	// TextMessage .
	TextMessage MessageType = 1
	// BinaryMessage .
	BinaryMessage MessageType = 2
	// CloseMessage .
	CloseMessage MessageType = 8
	// PingMessage .
	PingMessage MessageType = 9
	// PongMessage .
	PongMessage MessageType = 10
)

const (
	maskBit = 1 << 7
)

// Conn .
type Conn struct {
	*commonFields
	net.Conn

	mux sync.Mutex

	closeErr error

	chSessionInited chan struct{}
	session         interface{}

	subprotocol string

	compressionLevel int
	onClose          func(c *Conn, err error)

	sendQueue                []*[]byte
	sendQueueSize            uint16
	closed                   bool
	isClient                 bool
	enableCompression        bool
	remoteCompressionEnabled bool
	enableWriteCompression   bool
	isBlockingMod            bool
	isReadingByParser        bool
	isInReadingLoop          bool
	expectingFragments       bool
	compress                 bool
	releasePayload           bool
	msgType                  MessageType
	message                  *[]byte
	bytesCached              *[]byte

	Engine  *nbhttp.Engine
	Execute func(f func()) bool
}

//go:norace
func (c *Conn) UnderlayerConn() net.Conn {
	_ = "STUB: not implemented"

	// IsClient .
	//
	//go:norace
	return *new(net.Conn)
}

func (c *Conn) IsClient() bool {
	_ = "STUB: not implemented"

	// SetClient .
	//
	//go:norace
	return false
}

func (c *Conn) SetClient(isClient bool) { _ = "STUB: not implemented"; return }

// IsBlockingMod .
//
//go:norace
func (c *Conn) IsBlockingMod() bool { _ = "STUB: not implemented"; return false }

// IsAsyncWrite .
//
//go:norace
func (c *Conn) IsAsyncWrite() bool { _ = "STUB: not implemented"; return false }

// Close .
//
//go:norace
func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithError .
//
//go:norace
func (c *Conn) CloseWithError(err error) { _ = "STUB: not implemented"; return }

// SetCloseError .
//
//go:norace
func (c *Conn) SetCloseError(err error) { _ = "STUB: not implemented"; return }

// CompressionEnabled .
//
//go:norace
func (c *Conn) CompressionEnabled() bool {
	_ = "STUB: not implemented"

	//go:norace
	return false
}

func (c *Conn) safeBufferPointer(pbody *[]byte) *[]byte { _ = "STUB: not implemented"; return nil }

//go:norace
func (c *Conn) handleDataFrame(opcode MessageType, fin bool, pbody *[]byte) {
	_ = "STUB: not implemented"
	return
}

//go:norace
func (c *Conn) handleMessage(opcode MessageType, pbody *[]byte) { _ = "STUB: not implemented"; return }

//go:norace
func (c *Conn) handleProtocolMessage(opcode MessageType, pbody *[]byte) {
	_ = "STUB: not implemented"
	return
}

//go:norace
func (c *Conn) handleWsMessage(opcode MessageType, pData *[]byte) {
	_ = "STUB: not implemented"
	return
}

// no status

// protocol_error

//go:norace
func (c *Conn) nextFrame() (int, MessageType, []byte, bool, bool, bool, error) {
	_ = "STUB: not implemented"
	return 0, *new(MessageType), nil, false, false, false, nil
}

// Read .
//
//go:norace
func (c *Conn) Parse(data []byte) (retErr error) { _ = "STUB: not implemented"; return nil }

// if compressed, should check utf8 after decompressed the whole message.
// if c.msgType == TextMessage && len(frame) > 0 && !c.Engine.CheckUtf8(frame) {
// 	c.Conn.Close()
// 	err = ErrInvalidUtf8
// 	return
// }

// need more data

// OnMessage .
//
//go:norace
func (c *Conn) OnMessage(h func(*Conn, MessageType, []byte)) { _ = "STUB: not implemented"; return }

// OnMessagePtr .
//
//go:norace
func (c *Conn) OnMessagePtr(h func(*Conn, MessageType, *[]byte)) { _ = "STUB: not implemented"; return }

// OnDataFrame .
//
//go:norace
func (c *Conn) OnDataFrame(h func(*Conn, MessageType, bool, []byte)) {
	_ = "STUB: not implemented"
	return
}

// OnDataFramePtr .
//
//go:norace
func (c *Conn) OnDataFramePtr(h func(*Conn, MessageType, bool, *[]byte)) {
	_ = "STUB: not implemented"
	return
}

// EnableCompression .
//
//go:norace
func (c *Conn) EnableCompression(enable bool) { _ = "STUB: not implemented"; return }

//go:norace
func (c *Conn) OnClose(h func(*Conn, error)) {
	_ = "STUB: not implemented"

	// WriteClose .
	//
	//go:norace
	return
}

func (c *Conn) WriteClose(code int, reason string) error { _ = "STUB: not implemented"; return nil }

// WriteMessage .
//
//go:norace
func (c *Conn) WriteMessage(messageType MessageType, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Keepalive .
//
//go:norace
func (c *Conn) Keepalive(d time.Duration) *time.Timer { _ = "STUB: not implemented"; return nil }

// Session returns user session.
//
//go:norace
func (c *Conn) Session() interface{} { _ = "STUB: not implemented"; return nil }

// SessionWithLock returns user session with lock, returns as soon as the session has been seted.
//
//go:norace
func (c *Conn) SessionWithLock() interface{} { _ = "STUB: not implemented"; return nil }

// SessionWithContext returns user session, returns as soon as the session has been seted or
// waits until the context is done.
//
//go:norace
func (c *Conn) SessionWithContext(ctx context.Context) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// SetSession sets user session.
//
//go:norace
func (c *Conn) SetSession(session interface{}) { _ = "STUB: not implemented"; return }

type writeBuffer struct {
	pbuf      *[]byte
	allocator mempool.Allocator
}

// Write .
//
//go:norace
func (w *writeBuffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close .
//
//go:norace
func (w *writeBuffer) Close() error { _ = "STUB: not implemented"; return nil }

// CloseAndClean .
//
//go:norace
func (c *Conn) CloseAndClean(err error) {
	_ = "STUB: not implemented"
	// c.WriteClose(1000, "normal close")
	return
}

// WriteFrame .
//
//go:norace
func (c *Conn) WriteFrame(messageType MessageType, sendOpcode, fin bool, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//go:norace
func (c *Conn) writeFrame(messageType MessageType, sendOpcode, fin bool, data []byte, compress bool) error {
	_ = "STUB: not implemented"
	return nil
}

// opcode

// fin

// Write overwrites nbio.Conn.Write.
//
//go:norace
func (c *Conn) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// EnableWriteCompression .
//
//go:norace
func (c *Conn) EnableWriteCompression(enable bool) { _ = "STUB: not implemented"; return }

// Subprotocol returns the negotiated websocket subprotocol.
//
//go:norace
func (c *Conn) Subprotocol() string { _ = "STUB: not implemented"; return "" }

//go:norace
func NewClientConn(opt *Options, c net.Conn, subprotocol string, remoteCompressionEnabled bool, asyncWrite bool) *Conn {
	_ = "STUB: not implemented"
	return nil
}

//go:norace
func NewServerConn(u *Upgrader, c net.Conn, subprotocol string, remoteCompressionEnabled bool, asyncWrite bool) *Conn {
	_ = "STUB: not implemented"
	return nil
}

//go:norace
func newConn(u *Upgrader, c net.Conn, subprotocol string, remoteCompressionEnabled bool, asyncWrite bool, isClient bool) *Conn {
	_ = "STUB: not implemented"
	return nil
}

// HandleRead .
//
//go:norace
func (c *Conn) HandleRead(bufSize int) { _ = "STUB: not implemented"; return }

// return false if length is ok.
//
//go:norace
func (c *Conn) isMessageTooLarge(len int) bool {
	_ = "STUB: not implemented"
	// <=0 means unlimitted size
	return false
}

//go:norace
func (c *Conn) validFrame(opcode MessageType, fin, res1, res2, res3, expectingFragments bool) error {
	_ = "STUB: not implemented"
	return nil
}

//go:norace
func (c *Conn) readAll(r io.Reader, size int) (*[]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// can not extend more bytes.

// extend to the limit size at most.

//go:norace
func validCloseCode(code int) bool { _ = "STUB: not implemented"; return false }

//| Normal Closure  | hybi@ietf.org | RFC 6455  |

//      | Going Away      | hybi@ietf.org | RFC 6455  |

//   | Protocol error  | hybi@ietf.org | RFC 6455  |

//     | Unsupported Data| hybi@ietf.org | RFC 6455  |

//     | ---Reserved---- | hybi@ietf.org | RFC 6455  |

//      | No Status Rcvd  | hybi@ietf.org | RFC 6455  |

//      | Abnormal Closure| hybi@ietf.org | RFC 6455  |

//      | Invalid frame   | hybi@ietf.org | RFC 6455  |
//      |            | payload data    |               |           |

//     | Policy Violation| hybi@ietf.org | RFC 6455  |

//       | Message Too Big | hybi@ietf.org | RFC 6455  |

//       | Mandatory Ext.  | hybi@ietf.org | RFC 6455  |

//       | Internal Server | hybi@ietf.org | RFC 6455  |
//     |            | Error           |               |           |

//  | TLS handshake   | hybi@ietf.org | RFC 6455

// IANA registration policy and should be granted in the range 3000-3999.
// The range of status codes from 4000-4999 is designated for Private

//go:norace
func maskXOR(b, key []byte) { _ = "STUB: not implemented"; return }

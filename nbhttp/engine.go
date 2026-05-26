// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbhttp

import (
	"context"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/lesismal/llib/std/crypto/tls"
	"github.com/lesismal/nbio"
	"github.com/lesismal/nbio/lmux"
	"github.com/lesismal/nbio/mempool"
)

const (
	// IOModNonBlocking represents that the server serve all the connections by nbio poller goroutines to handle io events.
	IOModNonBlocking = 0
	// IOModBlocking represents that the server serve each connection with one goroutine at least to handle reading.
	IOModBlocking = 1
	// IOModMixed represents that the server creates listener mux to handle different connections, 1 listener will be dispatch to two ChanListener:
	// If ChanListener A's online is less than its max online num, the new connection will be dispatch to this listener A and served by single goroutine;
	// Else the new connection will be dispatch to ChanListener B and served by nbio poller.
	IOModMixed = 2

	// DefaultIOMod represents the default IO Mod used by nbhttp.Engine.
	DefaultIOMod = IOModNonBlocking
	// DefaultMaxBlockingOnline represents the default num of connections that will be dispatched to ChanListner A.
	DefaultMaxBlockingOnline = 10000
)

const (
	// DefaultMaxLoad .
	DefaultMaxLoad = 1024 * 1024

	// DefaultHTTPReadLimit .
	DefaultHTTPReadLimit = 1024 * 1024 * 64

	// DefaultMaxWebsocketFramePayloadSize .
	DefaultMaxWebsocketFramePayloadSize = 1024 * 32

	// DefaultKeepaliveTime .
	DefaultKeepaliveTime = time.Second * 120

	// DefaultBlockingReadBufferSize sets to 4k.
	DefaultBlockingReadBufferSize = 1024 * 4
)

const defaultNetwork = "tcp"

// ConfAddr .
type ConfAddr struct {
	Network   string
	Addr      string
	NListener int
	TLSConfig *tls.Config
	pAddr     *string
}

// Config .
type Config struct {
	// Name describes your gopher name for logging, it's set to "NB" by default.
	Name string

	// Network is the global listening protocol, used with Addrs toghter.
	// tcp* supported only by now, there's no plan for other protocol such as udp,
	// because it's too easy to write udp server/client.
	Network string

	// TLSConfig is the global tls config for all tls addrs.
	TLSConfig *tls.Config

	// Addrs is the non-tls listening addr list for an Engine.
	// if it is empty, no listener created, then the Engine is used for client by default.
	Addrs []string

	// AddrsTLS is the tls listening addr list for an Engine.
	// Engine will create listeners by AddrsTLS if it's not empty.
	AddrsTLS []string

	// AddrConfigs is the non-tls listening addr details list for an Engine.
	AddrConfigs []ConfAddr

	// AddrConfigsTLS is the tls listening addr details list for an Engine.
	AddrConfigsTLS []ConfAddr

	// Listen is used to create listener for Engine.
	Listen func(network, addr string) (net.Listener, error)

	// ListenUDP is used to create udp listener for Engine.
	ListenUDP func(network string, laddr *net.UDPAddr) (*net.UDPConn, error)

	// MaxLoad represents the max online num, it's set to 10k by default.
	MaxLoad int

	// NListener represents listner goroutine num for each ConfAddr, it's set to 1 by default.
	NListener int

	// NPoller represents poller goroutine num.
	NPoller int

	// ReadLimit represents the max size for parser reading, it's set to 64M by default.
	ReadLimit int

	// MaxHTTPBodySize represents the max size of HTTP body for parser reading.
	MaxHTTPBodySize int

	// ReadBufferSize represents buffer size for reading, it's set to 64k by default.
	ReadBufferSize int

	// MaxWriteBufferSize represents max write buffer size for Conn, 0 by default, represents no limit for writeBuffer
	// if MaxWriteBufferSize is set greater than to 0, and the connection's Send-Q is full and the data cached by nbio is
	// more than MaxWriteBufferSize, the connection would be closed by nbio.
	MaxWriteBufferSize int

	// MaxWebsocketFramePayloadSize represents max payload size of websocket frame.
	MaxWebsocketFramePayloadSize int

	// MessageHandlerPoolSize represents max http server's task pool goroutine num, it's set to runtime.NumCPU() * 256 by default.
	MessageHandlerPoolSize int

	// MessageHandlerTaskIdleTime represents idle time for task pool's goroutine, it's set to 60s by default.
	// MessageHandlerTaskIdleTime time.Duration

	// WriteTimeout represents Conn's write time out when response to a HTTP request.
	WriteTimeout time.Duration

	// KeepaliveTime represents Conn's ReadDeadline when waiting for a new request, it's set to 120s by default.
	KeepaliveTime time.Duration

	// LockListener represents listener's goroutine to lock thread or not, it's set to false by default.
	LockListener bool

	// LockPoller represents poller's goroutine to lock thread or not, it's set to false by default.
	LockPoller bool

	// DisableSendfile .
	DisableSendfile bool

	// ReleaseWebsocketPayload automatically release data buffer after function each call to websocket OnMessage or OnDataFrame.
	ReleaseWebsocketPayload bool

	// RetainHTTPBody represents whether to automatically release HTTP body's buffer after calling HTTP handler.
	RetainHTTPBody bool

	// MaxConnReadTimesPerEventLoop represents max read times in one poller loop for one fd.
	MaxConnReadTimesPerEventLoop int

	// OnAcceptError is called when accept error.
	OnAcceptError func(err error)

	// Handler sets HTTP handler for Engine.
	Handler http.Handler

	// `OnRequest` sets HTTP handler which will be called before `Handler.ServeHTTP`.
	//
	// A `Request` is pushed into a task queue and waits to be executed by the goroutine pool by default, which means the `Request`
	// may not be executed at once and may wait for long to be executed: if the client-side supports `pipeline` and the previous
	// `Requests` are handled for long. In some scenarios, we need to know when the `Request` is received, then we can control and
	// customize whether we should drop the task or record the real processing time for the `Request`. That is what this func should
	// be used for.
	OnRequest http.HandlerFunc

	// ServerExecutor sets the executor for data reading callbacks.
	ServerExecutor func(f func())

	// ClientExecutor sets the executor for client callbacks.
	ClientExecutor func(f func())

	// TLSAllocator sets the buffer allocator for TLS.
	TLSAllocator tls.Allocator

	// BodyAllocator sets the buffer allocator for HTTP.
	BodyAllocator mempool.Allocator

	// Context sets common context for Engine.
	Context context.Context

	// Cancel sets the cancel func for common context.
	Cancel func()

	// SupportServerOnly .
	SupportServerOnly bool

	// IOMod represents io mod, it is set to IOModNonBlocking by default.
	IOMod int
	// MaxBlockingOnline represents max blocking conn's online num.
	MaxBlockingOnline int
	// BlockingReadBufferSize represents read buffer size of blocking mod.
	BlockingReadBufferSize int

	// EpollMod .
	EpollMod uint32
	// EPOLLONESHOT .
	EPOLLONESHOT uint32

	// ReadBufferPool .
	ReadBufferPool mempool.Allocator

	// Deprecated.
	// WebsocketCompressor .
	WebsocketCompressor func(w io.WriteCloser, level int) io.WriteCloser

	// Deprecated.
	// WebsocketDecompressor .
	WebsocketDecompressor func(r io.Reader) io.ReadCloser

	// AsyncReadInPoller represents how the reading events and reading are handled
	// by epoll goroutine:
	// true : epoll goroutine handles the reading events only, another goroutine
	//        pool will handles the reading.
	// false: epoll goroutine handles both the reading events and the reading.
	//        false is by defalt.
	AsyncReadInPoller bool
	// IOExecute is used to handle the aysnc reading, users can customize it.
	IOExecute func(f func(*[]byte))
}

// Engine .
type Engine struct {
	*nbio.Engine
	Config

	CheckUtf8 func(data []byte) bool

	shutdown bool

	listenerMux *lmux.ListenerMux
	listeners   []net.Listener

	_onAcceptError func(err error)
	_onOpen        func(c net.Conn)
	_onClose       func(c net.Conn, err error)
	_onStop        func()

	mux         sync.Mutex
	conns       map[connValue]struct{}
	dialerConns map[connValue]struct{}

	// tlsBuffers [][]byte
	// getTLSBuffer func(c *nbio.Conn) []byte

	emptyRequest *http.Request
	BaseCtx      context.Context
	Cancel       func()

	SyncCall      func(f func())
	ExecuteClient func(f func())

	// isOneshot bool
}

// OnAcceptError is called when accept error.
//
//go:norace
func (e *Engine) OnAcceptError(h func(err error)) { _ = "STUB: not implemented"; return }

// OnOpen registers callback for new connection.
//
//go:norace
func (e *Engine) OnOpen(h func(c net.Conn)) {
	_ = "STUB: not implemented"

	// OnClose registers callback for disconnected.
	//
	//go:norace
	return
}

func (e *Engine) OnClose(h func(c net.Conn, err error)) {
	_ = "STUB: not implemented"

	// OnStop registers callback before Engine is stopped.
	//
	//go:norace
	return
}

func (e *Engine) OnStop(h func()) {
	_ = "STUB: not implemented"

	// Online .
	//
	//go:norace
	return
}

func (e *Engine) Online() int { _ = "STUB: not implemented"; return 0 }

// DialerOnline .
//
//go:norace
func (e *Engine) DialerOnline() int { _ = "STUB: not implemented"; return 0 }

//go:norace
func (e *Engine) closeAllConns() { _ = "STUB: not implemented"; return }

type Conn struct {
	net.Conn
	Parser    *Parser
	Trasfered bool
}

//go:norace
func (e *Engine) listen(ln net.Listener, tlsConfig *tls.Config, addConn func(*Conn, *tls.Config, func()), decrease func()) {
	_ = "STUB: not implemented"
	return
}

// ln.Close()

//go:norace
func (e *Engine) startListeners() error { _ = "STUB: not implemented"; return nil }

//go:norace
func (e *Engine) stopListeners() { _ = "STUB: not implemented"; return }

// SetETAsyncRead .
//
//go:norace
func (e *Engine) SetETAsyncRead() { _ = "STUB: not implemented"; return }

// SetLTSyncRead .
//
//go:norace
func (e *Engine) SetLTSyncRead() { _ = "STUB: not implemented"; return }

// Start .
//
//go:norace
func (e *Engine) Start() error { _ = "STUB: not implemented"; return nil }

// Stop .
//
//go:norace
func (e *Engine) Stop() { _ = "STUB: not implemented"; return }

// Shutdown .
//
//go:norace
func (e *Engine) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// DataHandler .
//
//go:norace
func (e *Engine) DataHandler(c *nbio.Conn, data []byte) { _ = "STUB: not implemented"; return }

// TLSDataHandler .
//
//go:norace
func (e *Engine) TLSDataHandler(c *nbio.Conn, data []byte) { _ = "STUB: not implemented"; return }

// c.SetReadDeadline(time.Now().Add(conf.KeepaliveTime))

// AddTransferredConn .
//
//go:norace
func (engine *Engine) AddTransferredConn(nbc *nbio.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// AddConnNonTLSNonBlocking .
//
//go:norace
func (engine *Engine) AddConnNonTLSNonBlocking(conn *Conn, tlsConfig *tls.Config, decrease func()) {
	_ = "STUB: not implemented"
	return
}

// if engine.isOneshot {
// 	parser.Execute = SyncExecutor
// }

// AddConnNonTLSBlocking .
//
//go:norace
func (engine *Engine) AddConnNonTLSBlocking(conn *Conn, tlsConfig *tls.Config, decrease func()) {
	_ = "STUB: not implemented"
	return
}

// AddConnTLSNonBlocking .
//
//go:norace
func (engine *Engine) AddConnTLSNonBlocking(conn *Conn, tlsConfig *tls.Config, decrease func()) {
	_ = "STUB: not implemented"
	return
}

// if engine.isOneshot {
// 	parser.Execute = SyncExecutor
// }

// AddConnTLSBlocking .
//
//go:norace
func (engine *Engine) AddConnTLSBlocking(conn *Conn, tlsConfig *tls.Config, decrease func()) {
	_ = "STUB: not implemented"
	return
}

//go:norace
func (engine *Engine) readConnBlocking(conn *Conn, parser *Parser, decrease func()) {
	_ = "STUB: not implemented"
	return
}

// }()

//go:norace
func (engine *Engine) readTLSConnBlocking(conn *Conn, rconn net.Conn, tlsConn *tls.Conn, parser *Parser, decrease func()) {
	_ = "STUB: not implemented"
	return
}

// NewEngine .
//
//go:norace
func NewEngine(conf Config) *Engine { _ = "STUB: not implemented"; return nil }

// avoid deadlock

// init non-tls addr configs

// init tls addr configs

// shouldSupportTLS := !conf.SupportServerOnly || len(conf.AddrsTLS) > 0
// if shouldSupportTLS {
// 	engine.InitTLSBuffers()
// }

// g.OnOpen(engine.ServerOnOpen)

// engine.isOneshot = (conf.EpollMod == nbio.EPOLLET && conf.EPOLLONESHOT == nbio.EPOLLONESHOT && runtime.GOOS == "linux")
// if engine.isOneshot {
// 	readBufferPool := conf.ReadBufferPool
// 	if readBufferPool == nil {
// 		readBufferPool = getReadBufferPool(conf.ReadBufferSize)
// 	}

// 	g.OnRead(func(c *nbio.Conn) {
// 		serverExecutor(func() {
// 			buf := readBufferPool.Malloc(conf.ReadBufferSize)
// 			defer func() {
// 				readBufferPool.Free(buf)
// 				c.ResetPollerEvent()
// 			}()
// 			for {
// 				n, err := c.Read(buf)
// 				if n > 0 && c.DataHandler != nil {
// 					c.DataHandler(c, buf[:n])
// 				}
// 				if errors.Is(err, syscall.EINTR) {
// 					continue
// 				}
// 				if errors.Is(err, syscall.EAGAIN) {
// 					break
// 				}
// 				if err != nil {
// 					c.CloseWithError(err)
// 				}
// 				if n < len(buf) {
// 					return
// 				}
// 			}
// 		})
// 	})
// }

var ReadBufferPools = &sync.Map{}

//go:norace
func getReadBufferPool(size int) mempool.Allocator {
	_ = "STUB: not implemented"
	return *new(mempool.Allocator)
}

//go:norace
func SyncExecutor(f func()) bool { _ = "STUB: not implemented"; return false }

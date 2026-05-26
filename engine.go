// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbio

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/lesismal/nbio/mempool"
	"github.com/lesismal/nbio/taskpool"
	"github.com/lesismal/nbio/timer"
)

const (
	// DefaultReadBufferSize .
	DefaultReadBufferSize = 1024 * 64

	// DefaultMaxWriteBufferSize .
	DefaultMaxWriteBufferSize = 0

	// DefaultMaxConnReadTimesPerEventLoop .
	DefaultMaxConnReadTimesPerEventLoop = 3

	// DefaultUDPReadTimeout .
	DefaultUDPReadTimeout = 120 * time.Second
)

const (
	NETWORK_TCP        = "tcp"
	NETWORK_TCP4       = "tcp4"
	NETWORK_TCP6       = "tcp6"
	NETWORK_UDP        = "udp"
	NETWORK_UDP4       = "udp4"
	NETWORK_UDP6       = "udp6"
	NETWORK_UNIX       = "unix"
	NETWORK_UNIXGRAM   = "unixgram"
	NETWORK_UNIXPACKET = "unixpacket"
)

var (
	// MaxOpenFiles .
	MaxOpenFiles = 1024 * 1024 * 2
)

// Config Of Engine.
type Config struct {
	// Name describes your gopher name for logging, it's set to "NB" by default.
	Name string

	// Network is the listening protocol, used with Addrs together.
	Network string

	// Addrs is the listening addr list for a nbio server.
	// if it is empty, no listener created, then the Engine is used for client by default.
	Addrs []string

	// NPoller represents poller goroutine num.
	NPoller int

	// ReadBufferSize represents buffer size for reading, it's set to 64k by default.
	ReadBufferSize int

	// MaxWriteBufferSize represents max write buffer size for Conn, 0 by default, represents no limit for writeBuffer
	// if MaxWriteBufferSize is set greater than to 0, and the connection's Send-Q is full and the data cached by nbio is
	// more than MaxWriteBufferSize, the connection would be closed by nbio.
	MaxWriteBufferSize int

	// MaxConnReadTimesPerEventLoop represents max read times in one poller loop for one fd
	MaxConnReadTimesPerEventLoop int

	// LockListener represents whether to lock thread for listener's goroutine, false by default.
	LockListener bool

	// LockPoller represents whether to lock thread for poller's goroutine, false by default.
	LockPoller bool

	// EpollMod sets the epoll mod, EPOLLLT by default.
	EpollMod uint32

	// EPOLLONESHOT sets EPOLLONESHOT, 0 by default.
	EPOLLONESHOT uint32

	// UDPReadTimeout sets the timeout for udp sessions.
	UDPReadTimeout time.Duration

	// Listen is used to create listener for Engine.
	// Users can set this func to customize listener, such as reuseport.
	Listen func(network, addr string) (net.Listener, error)

	// ListenUDP is used to create udp listener for Engine.
	ListenUDP func(network string, laddr *net.UDPAddr) (*net.UDPConn, error)

	// AsyncReadInPoller represents how the reading events and reading are handled
	// by epoll goroutine:
	// true : epoll goroutine handles the reading events only, another goroutine
	//        pool will handle the reading.
	// false: epoll goroutine handles both the reading events and the reading.
	AsyncReadInPoller bool
	// IOExecute is used to handle the aysnc reading, users can customize it.
	IOExecute func(f func(*[]byte))

	// BodyAllocator sets the buffer allocator for write cache.
	BodyAllocator mempool.Allocator
}

// Gopher keeps old type to compatible with new name Engine.
type Gopher = Engine

//go:norace
func NewGopher(conf Config) *Gopher { _ = "STUB: not implemented"; return nil }

// Engine is a manager of poller.
type Engine struct {
	Config
	*timer.Timer
	sync.WaitGroup

	Execute func(f func())
	mux     sync.Mutex

	isOneshot bool

	wgConn sync.WaitGroup

	// store std connections, for Windows only.
	connsStd map[*Conn]struct{}

	// store *nix connections.
	connsUnix []*Conn

	// listeners.
	listeners []*poller
	pollers   []*poller

	// onAcceptError is called when accept error.
	onAcceptError func(err error)

	// onUDPListen for udp listener created.
	onUDPListen func(c *Conn)
	// callback for new connection connected.
	onOpen func(c *Conn)
	// callback for connection closed.
	onClose func(c *Conn, err error)
	// callback for reading event.
	onRead func(c *Conn)
	// callback for coming data.
	onDataPtr func(c *Conn, pdata *[]byte)
	// callback for writing data size caculation.
	onWrittenSize func(c *Conn, b []byte, n int)
	// callback for allocationg the reading buffer.
	onReadBufferAlloc func(c *Conn) *[]byte
	// callback for freeing the reading buffer.
	onReadBufferFree func(c *Conn, pbuf *[]byte)

	// depreacated.
	// beforeRead  func(c *Conn)
	// afterRead   func(c *Conn)
	// beforeWrite func(c *Conn)

	// callback for Engine stop.
	onStop func()

	ioTaskPool *taskpool.IOTaskPool
}

// SetETAsyncRead .
//
//go:norace
func (e *Engine) SetETAsyncRead() { _ = "STUB: not implemented"; return }

// SetLTSyncRead .
//
//go:norace
func (e *Engine) SetLTSyncRead() { _ = "STUB: not implemented"; return }

// Stop closes listeners/pollers/conns/timer.
//
//go:norace
func (g *Engine) Stop() { _ = "STUB: not implemented"; return }

// Shutdown stops Engine gracefully with context.
//
//go:norace
func (g *Engine) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// AddConn adds conn to a poller.
//
//go:norace
func (g *Engine) AddConn(conn net.Conn) (*Conn, error) { _ = "STUB: not implemented"; return nil, nil }

//go:norace
func (g *Engine) addDialer(c *Conn) (*Conn, error) { _ = "STUB: not implemented"; return nil, nil }

// OnAcceptError is called when accept error.
//
//go:norace
func (g *Engine) OnAcceptError(h func(err error)) { _ = "STUB: not implemented"; return }

// OnOpen registers callback for new connection.
//
//go:norace
func (g *Engine) OnUDPListen(h func(c *Conn)) { _ = "STUB: not implemented"; return }

// OnOpen registers callback for new connection.
//
//go:norace
func (g *Engine) OnOpen(h func(c *Conn)) { _ = "STUB: not implemented"; return }

// OnClose registers callback for disconnected.
//
//go:norace
func (g *Engine) OnClose(h func(c *Conn, err error)) { _ = "STUB: not implemented"; return }

// OnRead registers callback for reading event.
//
//go:norace
func (g *Engine) OnRead(h func(c *Conn)) {
	_ = "STUB: not implemented"

	// OnData registers callback for data.
	//
	//go:norace
	return
}

func (g *Engine) OnData(h func(c *Conn, data []byte)) { _ = "STUB: not implemented"; return }

// OnDataPtr registers callback for data ptr.
//
//go:norace
func (g *Engine) OnDataPtr(h func(c *Conn, pdata *[]byte)) { _ = "STUB: not implemented"; return }

// OnWrittenSize registers callback for written size.
// If len(b) is bigger than 0, it represents that it's writing a buffer,
// else it's operating by Sendfile.
//
//go:norace
func (g *Engine) OnWrittenSize(h func(c *Conn, b []byte, n int)) { _ = "STUB: not implemented"; return }

// OnReadBufferAlloc registers callback for memory allocating.
//
//go:norace
func (g *Engine) OnReadBufferAlloc(h func(c *Conn) *[]byte) { _ = "STUB: not implemented"; return }

// OnReadBufferFree registers callback for memory release.
//
//go:norace
func (g *Engine) OnReadBufferFree(h func(c *Conn, pbuf *[]byte)) { _ = "STUB: not implemented"; return }

// Depracated .
// OnWriteBufferRelease registers callback for write buffer memory release.
// func (g *Engine) OnWriteBufferRelease(h func(c *Conn, b []byte)) {
// 	if h == nil {
// 		panic("invalid handler: nil")
// 	}
// 	g.onWriteBufferFree = h
// }

// BeforeRead registers callback before syscall.Read
// the handler would be called on windows.
// func (g *Engine) BeforeRead(h func(c *Conn)) {
// 	if h == nil {
// 		panic("invalid handler: nil")
// 	}
// 	g.beforeRead = h
// }

// Depracated .
// AfterRead registers callback after syscall.Read
// the handler would be called on *nix.
// func (g *Engine) AfterRead(h func(c *Conn)) {
// 	if h == nil {
// 		panic("invalid handler: nil")
// 	}
// 	g.afterRead = h
// }

// Depracated .
// BeforeWrite registers callback befor syscall.Write and syscall.Writev
// the handler would be called on windows.
// func (g *Engine) BeforeWrite(h func(c *Conn)) {
// 	if h == nil {
// 		panic("invalid handler: nil")
// 	}
// 	g.beforeWrite = h
// }

// OnStop registers callback before Engine is stopped.
//
//go:norace
func (g *Engine) OnStop(h func()) { _ = "STUB: not implemented"; return }

// PollerBuffer returns Poller's buffer by Conn, can be used on linux/bsd.
//
//go:norace
func (g *Engine) PollerBuffer(c *Conn) []byte { _ = "STUB: not implemented"; return nil }

// PollerBufferPtr returns Poller's buffer by Conn, can be used on linux/bsd.
//
//go:norace
func (g *Engine) PollerBufferPtr(c *Conn) *[]byte { _ = "STUB: not implemented"; return nil }

//go:norace
func (g *Engine) initHandlers() { _ = "STUB: not implemented"; return }

// g.OnRead(func(c *Conn, b []byte) ([]byte, error) {
// 	n, err := c.Read(b)
// 	if n > 0 {
// 		return b[:n], err
// 	}
// 	return nil, err
// })

// g.OnWriteBufferRelease(func(c *Conn, buffer []byte) {})
// g.BeforeRead(func(c *Conn) {})
// g.AfterRead(func(c *Conn) {})
// g.BeforeWrite(func(c *Conn) {})

//go:norace
func (g *Engine) borrow(c *Conn) *[]byte { _ = "STUB: not implemented"; return nil }

//go:norace
func (g *Engine) payback(c *Conn, pbuf *[]byte) { _ = "STUB: not implemented"; return }

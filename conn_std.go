// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package nbio

import (
	"bytes"
	"net"
	"sync"
	"syscall"
	"time"
)

// Conn wraps net.Conn.
type Conn struct {
	p *poller

	hash int

	mux sync.Mutex

	conn    net.Conn
	connUDP *udpConn

	rTimer *time.Timer

	typ      ConnType
	closed   bool
	closeErr error

	ReadBuffer []byte

	// user session.
	session interface{}

	jobList []func()

	cache *bytes.Buffer

	dataHandler func(c *Conn, data []byte)

	onConnected func(c *Conn, err error)
}

// Hash returns a hashcode.
//
//go:norace
func (c *Conn) Hash() int {
	_ = "STUB: not implemented"

	// Read wraps net.Conn.Read.
	//
	//go:norace
	return 0
}

func (c *Conn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//go:norace
func (c *Conn) read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//go:norace
func (c *Conn) readTCP(b []byte) (int, error) {
	_ = "STUB: not implemented"

	// g.beforeRead(c)
	return 0, nil
}

//go:norace
func (c *Conn) readUDP(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Write wraps net.Conn.Write.
//
//go:norace
func (c *Conn) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//go:norace
func (c *Conn) writeTCP(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// c.p.g.beforeWrite(c)
	return 0, nil
}

//go:norace
func (c *Conn) writeUDPClientFromDial(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//go:norace
func (c *Conn) writeUDPClientFromRead(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Writev wraps buffers.WriteTo/syscall.Writev.
//
//go:norace
func (c *Conn) Writev(in [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close wraps net.Conn.Close.
//
//go:norace
func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithError .
//
//go:norace
func (c *Conn) CloseWithError(err error) error { _ = "STUB: not implemented"; return nil }

// LocalAddr wraps net.Conn.LocalAddr.
//
//go:norace
func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// RemoteAddr wraps net.Conn.RemoteAddr.
//
//go:norace
func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// SetDeadline wraps net.Conn.SetDeadline.
//
//go:norace
func (c *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline wraps net.Conn.SetReadDeadline.
//
//go:norace
func (c *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline wraps net.Conn.SetWriteDeadline.
//
//go:norace
func (c *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetNoDelay wraps net.Conn.SetNoDelay.
//
//go:norace
func (c *Conn) SetNoDelay(nodelay bool) error { _ = "STUB: not implemented"; return nil }

// SetReadBuffer wraps net.Conn.SetReadBuffer.
//
//go:norace
func (c *Conn) SetReadBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetWriteBuffer wraps net.Conn.SetWriteBuffer.
//
//go:norace
func (c *Conn) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetKeepAlive wraps net.Conn.SetKeepAlive.
//
//go:norace
func (c *Conn) SetKeepAlive(keepalive bool) error { _ = "STUB: not implemented"; return nil }

// SetKeepAlivePeriod wraps net.Conn.SetKeepAlivePeriod.
//
//go:norace
func (c *Conn) SetKeepAlivePeriod(d time.Duration) error { _ = "STUB: not implemented"; return nil }

// SetLinger wraps net.Conn.SetLinger.
//
//go:norace
func (c *Conn) SetLinger(onoff int32, linger int32) error { _ = "STUB: not implemented"; return nil }

//go:norace
func newConn(conn net.Conn) *Conn { _ = "STUB: not implemented"; return nil }

// NBConn converts net.Conn to *Conn.
//
//go:norace
func NBConn(conn net.Conn) (*Conn, error) { _ = "STUB: not implemented"; return nil, nil }

type udpConn struct {
	*net.UDPConn
	rAddr *net.UDPAddr

	mux    sync.RWMutex
	parent *udpConn
	conns  map[string]*Conn
}

//go:norace
func (u *udpConn) Close() error { _ = "STUB: not implemented"; return nil }

//go:norace
func (u *udpConn) getConn(p *poller, rAddr *net.UDPAddr) (*Conn, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//go:norace
func (c *Conn) SyscallConn() (syscall.RawConn, error) {
	_ = "STUB: not implemented"
	return *new(syscall.RawConn), nil
}

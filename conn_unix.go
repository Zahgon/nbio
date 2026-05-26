// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build linux || darwin || netbsd || freebsd || openbsd || dragonfly
// +build linux darwin netbsd freebsd openbsd dragonfly

package nbio

import (
	"net"
	"sync"
	"syscall"
	"time"
)

// var (
// used to reset toWrite struct to empty value.
// emptyToWrite = toWrite{}

//	poolToWrite = sync.Pool{
//		New: func() interface{} {
//			return &toWrite{}
//		},
//	}
// )

//go:norace
func (c *Conn) newToWriteBuf(buf []byte) { _ = "STUB: not implemented"; return }

// poolToWrite.New().(*toWrite)

//go:norace
func (c *Conn) newToWriteFile(fd int, offset, remain int64) {
	_ = "STUB: not implemented"
	// poolToWrite.New().(*toWrite)
	return
}

//go:norace
func (c *Conn) releaseToWrite(t *toWrite) { _ = "STUB: not implemented"; return }

// *t = emptyToWrite
// poolToWrite.Put(t)

const maxWriteCacheOrFlushSize = 1024 * 64

type toWrite struct {
	fd     int     // file descriptor, used for sendfile
	buf    *[]byte // buffer to write
	offset int64   // buffer or file offset
	remain int64   // buffer or file remain bytes
}

// Conn implements net.Conn with non-blocking interfaces.
type Conn struct {
	mux sync.Mutex

	// the poller that handles io events for this connection.
	p *poller

	// file descriptor.
	fd int

	connUDP *udpConn

	// used for read deadline.
	rTimer *time.Timer
	// used for write deadline.
	wTimer *time.Timer

	// how many bytes are cached and wait to be written.
	left int
	// cache for buffers or files to be sent.
	writeList []*toWrite

	typ    ConnType
	closed bool

	// whether the writing event has been set in the poller.
	isWAdded bool
	// the first closing error.
	closeErr error

	// local addr.
	lAddr net.Addr
	// remote addr.
	rAddr net.Addr

	// user session.
	session interface{}

	// job list.
	jobList []func()

	readEvents int32

	dataHandler func(c *Conn, data []byte)

	onConnected func(c *Conn, err error)
}

// Hash returns a hash code of this connection.
//
//go:norace
func (c *Conn) Hash() int {
	_ = "STUB: not implemented"

	// AsyncReadInPoller is used for reading data async.
	//
	//go:norace
	return 0
}

func (c *Conn) AsyncRead() {
	_ = "STUB: not implemented"

	// If is EPOLLONESHOT, run the read job directly, because the reading event wouldn't
	// be re-dispatched before this reading event has been handled and set again.
	return
}

// If is not EPOLLONESHOT, the reading event may be re-dispatched for more than
// once, here we reduce the duplicate reading events.

// Only handle it when it's the first reading event.

// try to read all the data available.

// Read .
// Depracated .
// It was used to customize users' reading implementation, but better to use
// `ReadAndGetConn` instead, which can handle different types of connection and
// returns the consistent connection instance for UDP.
// Notice: non-blocking interface, should not be used as you use std.
//
//go:norace
func (c *Conn) Read(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// When the connection is closed and the fd is reused on Unix,
	// new connection maybe hold the same fd.
	// Use lock to prevent data confusion.
	return 0, nil
}

// if err == nil {
// 	c.p.g.afterRead(c)
// }

// ReadAndGetConn handles reading for different types of connection.
// It returns the real connection:
//  1. For Non-UDP connection, it returns the Conn itself.
//  2. For UDP connection, it may be a UDP Server fd, then it returns consistent
//     Conn for the same socket which has the same local addr and remote addr.
//
// Notice: non-blocking interface, should not be used as you use std.
//
//go:norace
func (c *Conn) ReadAndGetConn(pdata *[]byte) (*Conn, int, error) {
	_ = "STUB: not implemented"
	// When the connection is closed and the fd is reused on Unix,
	// new connection maybe hold the same fd.
	// Use lock to prevent data confusion.
	return nil, 0, nil
}

// if err == nil {
// 	c.p.g.afterRead(c)
// }

//go:norace
func (c *Conn) doRead(b []byte) (*Conn, int, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// no need to read for this type of connection,
// it's handled when reading ConnTypeUDPServer.

// read from TCP/Unix socket.
//
//go:norace
func (c *Conn) readStream(b []byte) (*Conn, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// read from UDP socket.
//
//go:norace
func (c *Conn) readUDP(b []byte) (*Conn, int, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// get or create and cache the consistent connection for the socket
// that has the same local addr and remote addr.

// Write writes data to the connection.
// Notice:
//  1. This is a non-blocking interface, but you can use it as you use std.
//  2. When it can't write all the data now, the connection will cache the data
//     left to be written and wait for the writing event then try to flush it.
//
//go:norace
func (c *Conn) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// c.p.g.beforeWrite(c)
	return 0, nil
}

// no data left to be written, clear write deadline timer.

// has data left to be written, set writing event.

// Writev does similar things as Write, but with [][]byte input arg.
// Notice: doesn't support UDP if more than 1 []byte.
//
//go:norace
func (c *Conn) Writev(in [][]byte) (int, error) {
	_ = "STUB: not implemented"
	// c.p.g.beforeWrite(c)
	return 0, nil
}

// no data left to be written, clear write deadline timer.

// has data left to be written, set writing event.

// write to TCP/Unix socket.
//
//go:norace
func (c *Conn) writeStream(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// write to UDP dialer.
//
//go:norace
func (c *Conn) writeUDPClientFromDial(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// write to UDP connection which is from server reading.
//
//go:norace
func (c *Conn) writeUDPClientFromRead(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements closes connection.
//
//go:norace
func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithError closes connection with user specified error.
//
//go:norace
func (c *Conn) CloseWithError(err error) error { _ = "STUB: not implemented"; return nil }

// LocalAddr returns the local network address, if known.
//
//go:norace
func (c *Conn) LocalAddr() net.Addr {
	_ = "STUB: not implemented"

	// RemoteAddr returns the remote network address, if known.
	//
	//go:norace
	return *new(net.Addr)
}

func (c *Conn) RemoteAddr() net.Addr {
	_ = "STUB: not implemented"

	// SetDeadline sets deadline for both read and write.
	// If it is time.Zero, SetDeadline will clear the deadlines.
	//
	//go:norace
	return *new(net.Addr)
}

func (c *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (c *Conn) setDeadline(timer **time.Timer, errClose error, t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// SetReadDeadline sets the deadline for future Read calls.
// When the user doesn't update the deadline and the deadline exceeds,
// the connection will be closed.
// If it is time.Zero, SetReadDeadline will clear the deadline.
//
// Notice:
//  1. Users should update the read deadline in time.
//  2. For example, call SetReadDeadline whenever a new WebSocket message
//     is received.
//
//go:norace
func (c *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline sets the deadline for future data writing.
// If it is time.Zero, SetReadDeadline will clear the deadline.
//
// If the next Write call writes all the data successfully and there's no data
// left to bewritten, the deadline timer will be cleared automatically;
// Else when the user doesn't update the deadline and the deadline exceeds,
// the connection will be closed.
//
//go:norace
func (c *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetNoDelay controls whether the operating system should delay
// packet transmission in hopes of sending fewer packets (Nagle's
// algorithm).  The default is true (no delay), meaning that data is
// sent as soon as possible after a Write.
//
//go:norace
func (c *Conn) SetNoDelay(nodelay bool) error { _ = "STUB: not implemented"; return nil }

// SetReadBuffer sets the size of the operating system's
// receive buffer associated with the connection.
//
//go:norace
func (c *Conn) SetReadBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetWriteBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
//
//go:norace
func (c *Conn) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetKeepAlive sets whether the operating system should send
// keep-alive messages on the connection.
//
//go:norace
func (c *Conn) SetKeepAlive(keepalive bool) error { _ = "STUB: not implemented"; return nil }

// SetKeepAlivePeriod sets period between keep-alives.
//
//go:norace
func (c *Conn) SetKeepAlivePeriod(d time.Duration) error { _ = "STUB: not implemented"; return nil }

// SetLinger .
//
//go:norace
func (c *Conn) SetLinger(onoff int32, linger int32) error { _ = "STUB: not implemented"; return nil }

// 1
// 0

// sets writing event.
//
//go:norace
func (c *Conn) modWrite() { _ = "STUB: not implemented"; return }

// reset io event to read only.
//
//go:norace
func (c *Conn) resetRead() { _ = "STUB: not implemented"; return }

//go:norace
func (c *Conn) write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// c.appendWrite(t)

// c.appendWrite(t)

//go:norace
func (c *Conn) writev(in [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// c.appendWrite(t)

// c.appendWrite(t)

// c.appendWrite(t)

// func (c *Conn) appendWrite(t *toWrite) {
// c.writeList = append(c.writeList, t)
// if t.buf != nil {
// 	c.left += len(t.buf)
// }
// }

// flush cached data to the fd when writing available.
//
//go:norace
func (c *Conn) flush() error { _ = "STUB: not implemented"; return nil }

// iovc := make([][]byte, 4)[0:0]
// writeBuffers := func() error {
// 	var (
// 		n    int
// 		err  error
// 		head *toWrite
// 	)

// 	if len(c.writeList) == 1 {
// 		head = c.writeList[0]
// 		buf := head.buf[head.offset:]
// 		for len(buf) > 0 && err == nil {
// 			n, err = syscall.Write(c.fd, buf)
// 			if n > 0 {
// 				if c.p.g.onWrittenSize != nil {
// 					c.p.g.onWrittenSize(c, buf[:n], n)
// 				}
// 				c.left -= n
// 				head.offset += int64(n)
// 				buf = buf[n:]
// 				if len(buf) == 0 {
// 					c.releaseToWrite(head)
// 					c.writeList = nil
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 		return err
// 	}

// 	writevSize := maxWriteCacheOrFlushSize
// 	iovc = iovc[0:0]
// 	for i := 0; i < len(c.writeList) && i < 1024; i++ {
// 		head = c.writeList[i]
// 		if head.buf != nil {
// 			b := head.buf[head.offset:]
// 			writevSize -= len(b)
// 			if writevSize < 0 {
// 				break
// 			}
// 			iovc = append(iovc, b)
// 		}
// 	}

// 	for len(iovc) > 0 && err == nil {
// 		n, err = writev(c, iovc)
// 		if n > 0 {
// 			c.left -= n
// 			for n > 0 {
// 				head = c.writeList[0]
// 				headLeft := len(head.buf) - int(head.offset)
// 				if n < headLeft {
// 					if onWrittenSize != nil {
// 						onWrittenSize(c, head.buf[head.offset:head.offset+int64(n)], n)
// 					}
// 					head.offset += int64(n)
// 					iovc[0] = iovc[0][n:]
// 					break
// 				} else {
// 					if onWrittenSize != nil {
// 						onWrittenSize(c, head.buf[head.offset:], headLeft)
// 					}
// 					c.releaseToWrite(head)
// 					c.writeList = c.writeList[1:]
// 					if len(c.writeList) == 0 {
// 						c.writeList = nil
// 					}
// 					iovc = iovc[1:]
// 					n -= headLeft
// 				}
// 			}
// 		} else {
// 			break
// 		}
// 	}
// 	return err
// }

// c.modWrite()

//go:norace
func (c *Conn) doWrite(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//go:norace
func (c *Conn) overflow(n int) bool { _ = "STUB: not implemented"; return false }

//go:norace
func (c *Conn) closeWithError(err error) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (c *Conn) closeWithErrorWithoutLock(err error) error { _ = "STUB: not implemented"; return nil }

// NBConn converts net.Conn to *Conn.
//
//go:norace
func NBConn(conn net.Conn) (*Conn, error) { _ = "STUB: not implemented"; return nil, nil }

type udpConn struct {
	parent *Conn

	rAddr    syscall.Sockaddr
	rAddrKey udpAddrKey

	mux   sync.RWMutex
	conns map[udpAddrKey]*Conn
}

//go:norace
func (u *udpConn) Close() error { _ = "STUB: not implemented"; return nil }

// This connection is created by reading from a UDP server,
// need to clear itself from the UDP server.

// This connection is a UDP server or dialer, need to close itself
// and close all children if this is a server.

//go:norace
func (u *udpConn) getConn(p *poller, fd int, rsa syscall.Sockaddr) (*Conn, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// new connection, create it.

// storage the consistent connection for the same remote addr.

type udpAddrKey [22]byte

//go:norace
func getUDPNetAddrKey(sa syscall.Sockaddr) udpAddrKey {
	_ = "STUB: not implemented"
	return *new(udpAddrKey)
}

//go:norace
func getUDPNetAddr(sa syscall.Sockaddr) *net.UDPAddr { _ = "STUB: not implemented"; return nil }

//go:norace
func (c *Conn) SyscallConn() (syscall.RawConn, error) {
	_ = "STUB: not implemented"
	return *new(syscall.RawConn), nil
}

type rawConn struct {
	fd int
}

//go:norace
func (c *rawConn) Control(f func(fd uintptr)) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (c *rawConn) Read(f func(fd uintptr) (done bool)) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (c *rawConn) Write(f func(fd uintptr) (done bool)) error {
	_ = "STUB: not implemented"
	return nil
}

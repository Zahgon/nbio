// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

package nbio

import (
	"net"
	"syscall"
)

const (
	// EPOLLLT .
	EPOLLLT = 0

	// EPOLLET .
	EPOLLET = 0x80000000

	// EPOLLONESHOT .
	EPOLLONESHOT = syscall.EPOLLONESHOT
)

const (
	epollEventsRead  = syscall.EPOLLPRI | syscall.EPOLLIN
	epollEventsWrite = syscall.EPOLLOUT
	epollEventsError = syscall.EPOLLERR | syscall.EPOLLHUP | syscall.EPOLLRDHUP
)

const (
	IPPROTO_TCP   = syscall.IPPROTO_TCP
	TCP_KEEPINTVL = syscall.TCP_KEEPINTVL
	TCP_KEEPIDLE  = syscall.TCP_KEEPIDLE
)

type poller struct {
	g *Engine // parent engine

	epfd  int // epoll fd
	evtfd int // event fd for trigger

	index int // poller index in engine

	pollType string // listener or io poller

	shutdown bool // state

	// whether poller is used for listener.
	isListener bool
	// listener.
	listener net.Listener
	// if poller is used as UnixConn listener,
	// store the addr and remove it when exit.
	unixSockAddr string

	ReadBuffer []byte // default reading buffer
}

// add the connection to poller and handle its io events.
//
//go:norace
func (p *poller) addConn(c *Conn) error { _ = "STUB: not implemented"; return nil }

// add the connection to poller and handle its io events.
//
//go:norace
func (p *poller) addDialer(c *Conn) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) getConn(fd int) *Conn { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) deleteConn(c *Conn) { _ = "STUB: not implemented"; return }

// p.deleteEvent(fd)

//go:norace
func (p *poller) start() { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) acceptorLoop() { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) readWriteLoop() { _ = "STUB: not implemented"; return }

// triggered by stop, exit event loop

// for socket connections

//go:norace
func (p *poller) stop() { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) addRead(fd int) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) resetRead(fd int) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) setRead(op int, fd int) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) modWrite(fd int) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) addReadWrite(fd int) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) setReadWrite(op int, fd int) error { _ = "STUB: not implemented"; return nil }

// func (p *poller) deleteEvent(fd int) error {
// 	return syscall.EpollCtl(
// 		p.epfd,
// 		syscall.EPOLL_CTL_DEL,
// 		fd,
// 		&syscall.EpollEvent{Fd: int32(fd)},
// 	)
// }

//go:norace
func newPoller(g *Engine, isListener bool, index int) (*poller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//go:norace
func (c *Conn) ResetPollerEvent() { _ = "STUB: not implemented"; return }

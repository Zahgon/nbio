// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build darwin || netbsd || freebsd || openbsd || dragonfly
// +build darwin netbsd freebsd openbsd dragonfly

package nbio

import (
	"net"
	"sync"
	"syscall"
)

const (
	// EPOLLLT .
	EPOLLLT = 0

	// EPOLLET .
	EPOLLET = 1

	// EPOLLONESHOT .
	EPOLLONESHOT = 0
)

const (
	IPPROTO_TCP   = 0
	TCP_KEEPINTVL = 0
	TCP_KEEPIDLE  = 0
)

type poller struct {
	mux sync.Mutex

	g *Engine

	kfd   int
	evtfd int

	index int

	shutdown bool

	listener     net.Listener
	isListener   bool
	unixSockAddr string

	ReadBuffer []byte

	pollType string

	eventList []syscall.Kevent_t
}

//go:norace
func (p *poller) addConn(c *Conn) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) addDialer(c *Conn) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) getConn(fd int) *Conn { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) deleteConn(c *Conn) { _ = "STUB: not implemented"; return }

// p.deleteEvent(fd)

//go:norace
func (p *poller) trigger() error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) addRead(fd int) { _ = "STUB: not implemented"; return }

// p.eventList = append(p.eventList, syscall.Kevent_t{Ident: uint64(fd), Flags: syscall.EV_ADD, Filter: syscall.EVFILT_WRITE})

//go:norace
func (p *poller) resetRead(fd int) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) modWrite(fd int) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) addReadWrite(fd int) { _ = "STUB: not implemented"; return }

// func (p *poller) deleteEvent(fd int) {
// 	p.mux.Lock()
// 	p.eventList = append(p.eventList,
// 		syscall.Kevent_t{Ident: uint64(fd), Flags: syscall.EV_DELETE, Filter: syscall.EVFILT_READ},
// 		syscall.Kevent_t{Ident: uint64(fd), Flags: syscall.EV_DELETE, Filter: syscall.EVFILT_WRITE})
// 	p.mux.Unlock()
// 	p.trigger()
// }

//go:norace
func (p *poller) readWrite(ev *syscall.Kevent_t) { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) start() { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) acceptorLoop() { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) readWriteLoop() { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) stop() { _ = "STUB: not implemented"; return }

//go:norace
func newPoller(g *Engine, isListener bool, index int) (*poller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//go:norace
func (c *Conn) ResetPollerEvent() { _ = "STUB: not implemented"; return }

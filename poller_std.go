// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package nbio

import (
	"net"
)

const (
	// EPOLLLT .
	EPOLLLT = 0

	// EPOLLET .
	EPOLLET = 1

	// EPOLLONESHOT .
	EPOLLONESHOT = 0
)

type poller struct {
	g *Engine

	index int

	ReadBuffer []byte

	pollType   string
	isListener bool
	listener   net.Listener
	shutdown   bool

	chStop chan struct{}
}

//go:norace
func (p *poller) accept() error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) readConn(c *Conn) { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) addConn(c *Conn) error { _ = "STUB: not implemented"; return nil }

// should not call onOpen for udp server conn

// should not read udp client from reading udp server conn

//go:norace
func (p *poller) addDialer(c *Conn) error { _ = "STUB: not implemented"; return nil }

//go:norace
func (p *poller) deleteConn(c *Conn) { _ = "STUB: not implemented"; return }

// should not call onClose for udp server conn

//go:norace
func (p *poller) start() { _ = "STUB: not implemented"; return }

//go:norace
func (p *poller) stop() { _ = "STUB: not implemented"; return }

//go:norace
func newPoller(g *Engine, isListener bool, index int) (*poller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//go:norace
func (c *Conn) ResetPollerEvent() { _ = "STUB: not implemented"; return }

// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package nbio

import (
	"time"
)

// Start inits and starts pollers.
//
//go:norace
func (g *Engine) Start() error {
	_ = "STUB: not implemented"
	// Create listener pollers.
	return nil
}

// Create IO pollers.

// Start IO pollers.

// Start TCP/Unix listener pollers.

// Start UDP listener pollers.

// g.Timer.Start()

// NewEngine creates an Engine and init default configurations.
//
//go:norace
func NewEngine(conf Config) *Engine { _ = "STUB: not implemented"; return nil }

// DialAsync connects asynchrony to the address on the named network.
//
//go:norace
func (engine *Engine) DialAsync(network, addr string, onConnected func(*Conn, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// DialAsync connects asynchrony to the address on the named network with timeout.
//
//go:norace
func (engine *Engine) DialAsyncTimeout(network, addr string, timeout time.Duration, onConnected func(*Conn, error)) error {
	_ = "STUB: not implemented"
	return nil
}

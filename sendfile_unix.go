// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build linux || darwin || netbsd || freebsd || openbsd || dragonfly
// +build linux darwin netbsd freebsd openbsd dragonfly

package nbio

import (
	"os"
)

const maxSendfileSize = 4 << 20

// Sendfile .
//
//go:norace
func (c *Conn) Sendfile(f *os.File, remain int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// f.Fd() will set the fd to blocking mod.
// We need to set the fd to non-blocking mod again.

// If c.writeList is not empty, the socket is not writable now.
// We push this File to writeList and wait to send it when writable.

// After this Sendfile func returns, fs will be closed by the caller.
// So we need to dup the fd and close it when we don't need it any more.

// c.appendWrite(t)

// c.p.g.beforeWrite(c)

// After this Sendfile func returns, fs will be closed by the caller.
// So we need to dup the fd and close it when we don't need it any more.

// c.appendWrite(t)

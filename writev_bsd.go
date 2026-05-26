// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build darwin || netbsd || freebsd || openbsd || dragonfly
// +build darwin netbsd freebsd openbsd dragonfly

package nbio

//go:norace
func writev(c *Conn, iovs [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

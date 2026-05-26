// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

package nbio

//go:norace
func writev(c *Conn, bs [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package nbio

import (
	"os"
)

// Sendfile .
//
//go:norace
func (c *Conn) Sendfile(f *os.File, remain int64) (written int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

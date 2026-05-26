// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build linux || darwin || netbsd || freebsd || openbsd || dragonfly
// +build linux darwin netbsd freebsd openbsd dragonfly

package nbio

import (
	"net"
	"syscall"
)

//go:norace
func init() {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err == nil {
		if n := int(limit.Max); n > 0 && n < MaxOpenFiles {
			MaxOpenFiles = n
		}
	}
}

//go:norace
func dupStdConn(conn net.Conn) (*Conn, error) { _ = "STUB: not implemented"; return nil, nil }

// err = syscall.SetNonblock(newFd, true)
// if err != nil {
// 	syscall.Close(newFd)
// 	return nil, err
// }

// no remote addr, this is a listener

// has remote addr, this is a dialer

//go:norace
func parseDomainAndType(network, addr string) (int, int, syscall.Sockaddr, net.Addr, ConnType, error) {
	_ = "STUB: not implemented"
	return 0, 0, *new(syscall.Sockaddr), *new(net.Addr), *new(ConnType), nil
}

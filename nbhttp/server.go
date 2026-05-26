// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package nbhttp

// Server .
type Server struct {
	*Engine
}

// NewServer .
//
//go:norace
func NewServer(conf Config, v ...interface{}) *Server { _ = "STUB: not implemented"; return nil }

// NewServerTLS .
//
//go:norace
func NewServerTLS(conf Config, v ...interface{}) *Server { _ = "STUB: not implemented"; return nil }

package tls

// deprecated.

import (
	"github.com/lesismal/llib/std/crypto/tls"
	"github.com/lesismal/nbio"
)

// Conn .
type Conn = tls.Conn

// Config .
type Config = tls.Config

// Dial returns a net.Conn to be added to a Engine.
//
//go:norace
func Dial(network, addr string, config *Config) (*tls.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WrapOpen returns an opening handler of nbio.Engine.
//
//go:norace
func WrapOpen(tlsConfig *Config, isClient bool, h func(c *nbio.Conn, tlsConn *Conn)) func(c *nbio.Conn) {
	_ = "STUB: not implemented"
	return nil
}

// WrapClose returns an closing handler of nbio.Engine.
//
//go:norace
func WrapClose(h func(c *nbio.Conn, tlsConn *Conn, err error)) func(c *nbio.Conn, err error) {
	_ = "STUB: not implemented"
	return nil
}

// WrapData returns a data handler of nbio.Engine.
//
//go:norace
func WrapData(h func(c *nbio.Conn, tlsConn *Conn, data []byte), args ...interface{}) func(c *nbio.Conn, data []byte) {
	_ = "STUB: not implemented"
	return nil
}

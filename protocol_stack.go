package nbio

import (
	"net"
)

type Protocol interface {
	Parse(c net.Conn, b []byte, ps *ProtocolStack) (net.Conn, []byte, error)
	Write(b []byte) (int, error)
}

type ProtocolStack struct {
	stack []Protocol
}

//go:norace
func (ps *ProtocolStack) Add(p Protocol) { _ = "STUB: not implemented"; return }

//go:norace
func (ps *ProtocolStack) Delete(p Protocol) { _ = "STUB: not implemented"; return }

//go:norace
func (ps *ProtocolStack) Parse(c net.Conn, b []byte, ps_ ProtocolStack) (net.Conn, []byte, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

//go:norace
func (ps *ProtocolStack) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

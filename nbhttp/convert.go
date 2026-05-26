package nbhttp

import (
	"net"
	"unsafe"
)

const (
	uintptrSize   = int(unsafe.Sizeof(uintptr(0)))
	connValueSize = uintptrSize + 1
)

const (
	connTypNONE byte = 0
	connTypNBIO byte = 1
	connTypTCP  byte = 2
	connTypUNIX byte = 3
	connTypTLS  byte = 4
	connTypLTLS byte = 5
)

// We can use this array-value as map key to reduce gc cost.
// Ref: https://github.com/lesismal/nbio/pull/304#issuecomment-1583880587
type connValue [connValueSize]byte

// Convert net.Conn to array value.
//
//go:norace
func conn2Array(conn net.Conn) (connValue, error) {
	_ = "STUB: not implemented"
	return *new(connValue), nil
}

// Convert array value to net.Conn.
//
//go:norace
func array2Conn(b connValue) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

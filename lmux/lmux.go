package lmux

import (
	"net"
)

type event struct {
	err  error
	conn net.Conn
}

type listenerAB struct {
	a, b *ChanListener
}

// New returns a ListenerMux.
//
//go:norace
func New(maxOnlineA int) *ListenerMux { _ = "STUB: not implemented"; return nil }

// ListenerMux manages listeners and handle the connection dispatching logic.
type ListenerMux struct {
	shutdown   bool
	listeners  map[net.Listener]listenerAB
	chClose    chan struct{}
	onlineA    int32
	maxOnlineA int32
}

// Mux creates and returns ChanListener A and B:
// If the online num of A is less than ListenerMux. maxOnlineA, the new connection will be dispatched to A;
// Else the new connection will be dispatched to B.
//
//go:norace
func (lm *ListenerMux) Mux(l net.Listener) (*ChanListener, *ChanListener) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start starts to accept and dispatch the connections to ChanListener A or B.
//
//go:norace
func (lm *ListenerMux) Start() { _ = "STUB: not implemented"; return }

// Exit the loop after a non recoverable error

// Stop stops all the listeners.
//
//go:norace
func (lm *ListenerMux) Stop() { _ = "STUB: not implemented"; return }

// DecreaseOnlineA decreases the online num of ChanListener A.
//
//go:norace
func (lm *ListenerMux) DecreaseOnlineA() { _ = "STUB: not implemented"; return }

// ChanListener .
type ChanListener struct {
	addr     net.Addr
	chEvent  chan event
	chClose  chan struct{}
	decrease func()
}

// Accept accepts a connection.
//
//go:norace
func (l *ChanListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Close does nothing but implementing net.Conn.Close.
// User should call ListenerMux.Close to close it automatically.
//
//go:norace
func (l *ChanListener) Close() error {
	_ = "STUB: not implemented"

	// Addr returns the listener's network address.
	//
	//go:norace
	return nil
}

func (l *ChanListener) Addr() net.Addr {
	_ = "STUB: not implemented"

	// Decrease decreases the online num if it's A.
	//
	//go:norace
	return *new(net.Addr)
}

func (l *ChanListener) Decrease() { _ = "STUB: not implemented"; return }

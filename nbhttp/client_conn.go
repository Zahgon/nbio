package nbhttp

import (
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/lesismal/llib/std/crypto/tls"
)

type resHandler struct {
	c net.Conn
	t time.Time
	h func(res *http.Response, conn net.Conn, err error)
}

// ClientConn .
type ClientConn struct {
	mux      sync.Mutex
	conn     net.Conn
	handlers []resHandler

	closed bool

	onClose func()

	Engine *Engine

	Jar http.CookieJar

	Timeout time.Duration

	IdleConnTimeout time.Duration

	TLSClientConfig *tls.Config

	Dial func(network, addr string) (net.Conn, error)

	Proxy func(*http.Request) (*url.URL, error)

	CheckRedirect func(req *http.Request, via []*http.Request) error
}

// Reset resets itself as new created.
//
//go:norace
func (c *ClientConn) Reset() { _ = "STUB: not implemented"; return }

// OnClose registers a callback for closing.
//
//go:norace
func (c *ClientConn) OnClose(h func()) {
	_ = "STUB: not implemented"

	// Close closes underlayer connection with EOF.
	//
	//go:norace
	return
}

func (c *ClientConn) Close() { _ = "STUB: not implemented"; return }

// CloseWithError closes underlayer connection with error.
//
//go:norace
func (c *ClientConn) CloseWithError(err error) { _ = "STUB: not implemented"; return }

//go:norace
func (c *ClientConn) closeWithErrorWithoutLock(err error) { _ = "STUB: not implemented"; return }

//go:norace
func (c *ClientConn) onResponse(res *http.Response, err error) { _ = "STUB: not implemented"; return }

// Do sends an HTTP request and returns an HTTP response.
// Notice:
//  1. It's blocking when Dial to the server;
//  2. It's non-blocking for waiting for the response;
//  3. It calls the handler when the response is received
//     or other errors occur, such as timeout.
//
//go:norace
func (c *ClientConn) Do(req *http.Request, handler func(res *http.Response, conn net.Conn, err error)) {
	_ = "STUB: not implemented"
	return
}

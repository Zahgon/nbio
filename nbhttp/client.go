package nbhttp

import (
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/lesismal/llib/std/crypto/tls"
)

//go:norace
func newHostConns(cli *Client) *hostConns { _ = "STUB: not implemented"; return nil }

// 1024 by default

type hostConns struct {
	mux        sync.Mutex
	cli        *Client
	conns      map[*ClientConn]struct{}
	connNum    int32
	maxConnNum int32
	chConnss   chan *ClientConn
}

//go:norace
func (hcs *hostConns) closeWithError(err error) { _ = "STUB: not implemented"; return }

//go:norace
func (hcs *hostConns) getConn() (*hostConns, *ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// 1. Get an existing free connection.

// 2. Try to create a new connection if the num of existing
//    connections is smaller than maxConnNum

// 3. Wait for an existed working connection to be free.

//go:norace
func (hcs *hostConns) releaseConn(hc *ClientConn) {
	_ = "STUB: not implemented"

	// Client implements the similar functions with std http.Client.
	return
}

type Client struct {
	mux    sync.Mutex
	closed bool

	connsMux     sync.RWMutex
	connsOfHosts map[string]*hostConns

	Engine *Engine

	Jar http.CookieJar

	Timeout time.Duration

	MaxConnsPerHost int32
	IdleConnTimeout time.Duration

	TLSClientConfig *tls.Config

	Dial func(network, addr string) (net.Conn, error)

	Proxy func(*http.Request) (*url.URL, error)

	CheckRedirect func(req *http.Request, via []*http.Request) error
}

// Close closes all underlayer connections with EOF.
//
//go:norace
func (c *Client) Close() { _ = "STUB: not implemented"; return }

// CloseWithError closes all underlayer connections with error.
//
//go:norace
func (c *Client) CloseWithError(err error) { _ = "STUB: not implemented"; return }

//go:norace
func (c *Client) getConn(host string) (*hostConns, *ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Do sends an HTTP request and returns an HTTP response.
// Notice:
//  1. It's blocking when Dial to the server;
//  2. It's non-blocking for waiting for the response;
//  3. It calls the handler when the response is received
//     or other errors occur, such as timeout.
//
//go:norace
func (c *Client) Do(req *http.Request, handler func(res *http.Response, conn net.Conn, err error)) {
	_ = "STUB: not implemented"
	return
}

type netDialerFunc func(network, addr string) (net.Conn, error)

//go:norace
func (fn netDialerFunc) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

var proxySchemes map[string]func(*url.URL, proxyDialer) (proxyDialer, error)

//go:norace
func proxyRegisterDialerType(scheme string, f func(*url.URL, proxyDialer) (proxyDialer, error)) {
	_ = "STUB: not implemented"
	return
}

//go:norace
func proxyFromURL(u *url.URL, forward proxyDialer) (proxyDialer, error) {
	_ = "STUB: not implemented"
	return *new(proxyDialer), nil
}

//go:norace
func hostPortNoPort(u *url.URL) (hostPort, hostNoPort string) {
	_ = "STUB: not implemented"
	return "", ""
}

type proxyDialer interface {
	Dial(network, addr string) (c net.Conn, err error)
}

type httpProxyDialer struct {
	proxyURL    *url.URL
	forwardDial func(network, addr string) (net.Conn, error)
}

//go:norace
func (hpd *httpProxyDialer) Dial(network string, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type proxyAuth struct {
	User, Password string
}

//go:norace
func proxySOCKS5(network, addr string, auth *proxyAuth, forward proxyDialer) (proxyDialer, error) {
	_ = "STUB: not implemented"
	return *new(proxyDialer), nil
}

type proxySocks5 struct {
	user, password string
	network, addr  string
	forward        proxyDialer
}

const proxySocks5Version = 5

const (
	proxySocks5AuthNone     = 0
	proxySocks5AuthPassword = 2
)

const proxySocks5Connect = 1

const (
	proxySocks5IP4    = 1
	proxySocks5Domain = 3
	proxySocks5IP6    = 4
)

var proxySocks5Errors = []string{
	"",
	"general failure",
	"connection forbidden",
	"network unreachable",
	"host unreachable",
	"connection refused",
	"TTL expired",
	"command not supported",
	"address type not supported",
}

//go:norace
func (s *proxySocks5) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

const errProxyAtSocks5Prefix = "proxy: SOCKS5 proxy at "

//go:norace
func (s *proxySocks5) connect(conn net.Conn, target string) error {
	_ = "STUB: not implemented"
	return nil
}

// the size here is just an estimate

/* num auth methods */

/* num auth methods */

// See RFC 1929

/* password protocol version */

/* reserved */

//go:norace
func init() {
	proxyRegisterDialerType("http", func(proxyURL *url.URL, forwardDialer proxyDialer) (proxyDialer, error) {
		return &httpProxyDialer{proxyURL: proxyURL, forwardDial: forwardDialer.Dial}, nil
	})
}

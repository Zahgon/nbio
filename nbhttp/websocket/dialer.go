package websocket

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/lesismal/llib/std/crypto/tls"
	"github.com/lesismal/nbio/nbhttp"
)

const (
	hostHeaderField                = "Host"
	upgradeHeaderField             = "Upgrade"
	connectionHeaderField          = "Connection"
	secWebsocketKeyHeaderField     = "Sec-Websocket-Key"
	secWebsocketVersionHeaderField = "Sec-Websocket-Version"
	secWebsocketExtHeaderField     = "Sec-Websocket-Extensions"
	secWebsocketProtoHeaderField   = "Sec-Websocket-Protocol"
)

// Dialer .
type Dialer struct {
	Engine *nbhttp.Engine

	Options  *Options
	Upgrader *Upgrader

	Jar http.CookieJar

	DialTimeout time.Duration

	TLSClientConfig *tls.Config

	Proxy func(*http.Request) (*url.URL, error)

	CheckRedirect func(req *http.Request, via []*http.Request) error

	Subprotocols []string

	EnableCompression bool

	Cancel context.CancelFunc
}

// Dial .
//
//go:norace
func (d *Dialer) Dial(urlStr string, requestHeader http.Header, v ...interface{}) (*Conn, *http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DialContext .
//
//go:norace
func (d *Dialer) DialContext(ctx context.Context, urlStr string, requestHeader http.Header, v ...interface{}) (*Conn, *http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

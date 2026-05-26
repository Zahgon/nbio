package websocket

import (
	"io"
	"net"
	"net/http"
	"time"

	"github.com/lesismal/nbio/nbhttp"
)

var (
	// DefaultBlockingReadBufferSize .
	DefaultBlockingReadBufferSize = 1024 * 4

	// DefaultBlockingModAsyncWrite .
	DefaultBlockingModAsyncWrite = true

	// DefaultBlockingModHandleRead .
	DefaultBlockingModHandleRead = true

	// DefaultBlockingModTransferConnToPoller .
	DefaultBlockingModTransferConnToPoller = false

	// DefaultBlockingModSendQueueInitSize .
	DefaultBlockingModSendQueueInitSize = 4

	// DefaultBlockingModSendQueueMaxSize .
	DefaultBlockingModSendQueueMaxSize uint16 = 0

	// DefaultMessageLengthLimit .
	DefaultMessageLengthLimit = 1024 * 1024 * 4

	// DefaultBlockingModAsyncCloseDelay .
	DefaultBlockingModAsyncCloseDelay = time.Second / 10

	// DefaultEngine will be set to a Upgrader.Engine to handle details such as buffers.
	DefaultEngine = nbhttp.NewEngine(nbhttp.Config{
		ReleaseWebsocketPayload: true,
	})
)

type commonFields struct {
	KeepaliveTime              time.Duration
	MessageLengthLimit         int
	BlockingModAsyncCloseDelay time.Duration

	ReleasePayload        bool
	WebsocketCompressor   func(c *Conn, w io.WriteCloser, level int) io.WriteCloser
	WebsocketDecompressor func(c *Conn, r io.Reader) io.ReadCloser

	pingMessageHandler  func(c *Conn, appData string)
	pongMessageHandler  func(c *Conn, appData string)
	closeMessageHandler func(c *Conn, code int, text string)

	openHandler      func(*Conn)
	messageHandler   func(c *Conn, messageType MessageType, messagePtr *[]byte)
	dataFrameHandler func(c *Conn, messageType MessageType, fin bool, framePtr *[]byte)
}

type Options = Upgrader

//go:norace
func NewOptions() *Options { _ = "STUB: not implemented"; return nil }

// Upgrader .
type Upgrader struct {
	commonFields

	// Engine .
	Engine *nbhttp.Engine

	// Subprotocols .
	Subprotocols []string

	// CheckOrigin .
	CheckOrigin func(r *http.Request) bool

	// HandshakeTimeout represents the timeout duration during websocket handshake.
	HandshakeTimeout time.Duration

	// BlockingModReadBufferSize represents the read buffer size of a Conn if it's in blocking mod.
	BlockingModReadBufferSize int

	// BlockingModAsyncWrite represents whether use a goroutine to handle writing:
	// true: use dynamic goroutine to handle writing.
	// false: write buffer to the conn directely.
	BlockingModAsyncWrite bool

	// BlockingModHandleRead represents whether start a goroutine to handle reading automatically during `Upgrade``:
	// true: start a new goroutine to handle reading.
	// false: use the current goroutine to handle reading.
	//
	// Notice:
	// If we start a goroutine to handle read during `Upgrade`, we may receive a new websocket message.
	// before we have left the http.Handler for the `Websocket Handshake`.
	// Then if we have the logic of `websocket.Conn.SetSession` in the http.Handler, it's possible that when we receive
	// and are handling a websocket message and call `websocket.Conn.Session()`, we get nil.
	//
	// To fix this nil session problem, can use `websocket.Conn.SessionWithLock()`.
	//
	// For other concurrent problems(including the nil session problem), we can:
	// 1st: set this `BlockingModHandleRead = false`
	// 2nd: `go wsConn.HandleRead(YourBufSize)` after `Upgrade` and finished initialization.
	// Then the websocket message wouldn't come before the http.Handler for `Websocket Handshake` has done.
	BlockingModHandleRead bool

	// BlockingModTrasferConnToPoller represents whether try to transfer a blocking connection to nonblocking and add to `Engine``.
	// true: try to transfer.
	// false: don't try to transfer.
	//
	// Notice:
	// Only `net.TCPConn` and `llib's blocking tls.Conn` can be transferred to nonblocking.
	BlockingModTrasferConnToPoller bool

	// BlockingModSendQueueInitSize represents the init size of a Conn's send queue,
	// only takes effect when `BlockingModAsyncWrite` is true.
	BlockingModSendQueueInitSize int

	// BlockingModSendQueueInitSize represents the max size of a Conn's send queue,
	// only takes effect when `BlockingModAsyncWrite` is true.
	BlockingModSendQueueMaxSize uint16

	enableCompression bool
	compressionLevel  int
	onClose           func(c *Conn, err error)
}

// NewUpgrader .
//
//go:norace
func NewUpgrader() *Upgrader { _ = "STUB: not implemented"; return nil }

// EnableCompression .
//
//go:norace
func (u *Upgrader) EnableCompression(enable bool) { _ = "STUB: not implemented"; return }

// SetCompressionLevel .
//
//go:norace
func (u *Upgrader) SetCompressionLevel(level int) error { _ = "STUB: not implemented"; return nil }

// SetCloseHandler .
//
//go:norace
func (u *Upgrader) SetCloseHandler(h func(*Conn, int, string)) { _ = "STUB: not implemented"; return }

// SetPingHandler .
//
//go:norace
func (u *Upgrader) SetPingHandler(h func(*Conn, string)) { _ = "STUB: not implemented"; return }

// SetPongHandler .
//
//go:norace
func (u *Upgrader) SetPongHandler(h func(*Conn, string)) { _ = "STUB: not implemented"; return }

// OnOpen .
//
//go:norace
func (u *Upgrader) OnOpen(h func(*Conn)) {
	_ = "STUB: not implemented"

	// OnMessage .
	//
	//go:norace
	return
}

func (u *Upgrader) OnMessage(h func(*Conn, MessageType, []byte)) { _ = "STUB: not implemented"; return }

// OnMessage .
//
//go:norace
func (u *Upgrader) OnMessagePtr(h func(*Conn, MessageType, *[]byte)) {
	_ = "STUB: not implemented"
	return
}

// OnDataFrame .
//
//go:norace
func (u *Upgrader) OnDataFrame(h func(*Conn, MessageType, bool, []byte)) {
	_ = "STUB: not implemented"
	return
}

// OnDataFramePtr .
//
//go:norace
func (u *Upgrader) OnDataFramePtr(h func(*Conn, MessageType, bool, *[]byte)) {
	_ = "STUB: not implemented"
	return
}

// OnClose .
//
//go:norace
func (u *Upgrader) OnClose(h func(*Conn, error)) {
	_ = "STUB: not implemented"

	// Upgrade .
	//
	//go:norace
	return
}

func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header, args ...interface{}) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Scenario 1: *nbio.Conn, handled by nbhttp.Engine.

// Scenario 2: llib's *tls.Conn.

// 2.1 The conn may be from std's http.Server.Serve(llib's tls.Listener),
//     or from nbhttp.Engine's IOModBlocking/Mixed(blocking part).

// 2.1.1 Transfer the conn to poller.

// 2.1.2 Don't transfer the conn to poller.

// 2.2 The conn is from nbio poller.

// Scenario 3: std's *net.TCPConn.

// 3.1 Transfer the conn to poller.

// 3.2 Don't transfer the conn to poller.

// Scenario 4: Unknown conn type, mostly is std *tls.Conn, from std http.Server.

// if parser != nil {
// 	parser.ReadCloser = wsc
// 	wsc.Execute = parser.Execute
// }

// If upgrader.ReleasePayload is false, which maybe by default,
// then set it to engine.ReleaseWebsocketPayload.
// Considering compatibility with old versions, should set it to
// upgrader.ReleasePayload only when upgrader.ReleasePayload is true.

//go:norace
func (u *Upgrader) UpgradeAndTransferConnToPoller(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//go:norace
func (u *Upgrader) UpgradeWithoutHandlingReadForConnFromSTDServer(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*Conn, error) {
	_ = "STUB: not implemented"
	// handle std server's conn, no need transfer conn to nbio Engine
	return nil, nil
}

//go:norace
func (u *Upgrader) commCheck(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

// Negotiate PMCE

//go:norace
func (u *Upgrader) commResponse(conn net.Conn, responseHeader http.Header, challengeKey, subprotocol string, compress bool) error {
	_ = "STUB: not implemented"
	return nil
}

// prevent response splitting.

//go:norace
func (u *Upgrader) returnError(w http.ResponseWriter, _ *http.Request, status int, err error) error {
	_ = "STUB: not implemented"
	return nil
}

//go:norace
func (u *Upgrader) selectSubprotocol(r *http.Request, responseHeader http.Header) string {
	_ = "STUB: not implemented"
	return ""
}

//go:norace
func subprotocols(r *http.Request) []string { _ = "STUB: not implemented"; return nil }

var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

//go:norace
func acceptKeyString(challengeKey string) string {
	_ = "STUB: not implemented"
	//nolint:gosec // per websocket protocol spec
	return ""
}

//go:norace
func acceptKeyBytes(challengeKey string) []byte {
	_ = "STUB: not implemented"
	//nolint:gosec // per websocket protocol spec
	return nil
}

//go:norace
func challengeKey() (string, error) { _ = "STUB: not implemented"; return "", nil }

//go:norace
func checkSameOrigin(r *http.Request) bool { _ = "STUB: not implemented"; return false }

//go:norace
func headerContains(header http.Header, name string, value string) bool {
	_ = "STUB: not implemented"
	return false
}

//go:norace
func equalASCIIFold(s, t string) bool { _ = "STUB: not implemented"; return false }

//go:norace
func parseExtensions(header http.Header) []map[string]string { _ = "STUB: not implemented"; return nil }

var isTokenOctet = [256]bool{
	'!':  true,
	'#':  true,
	'$':  true,
	'%':  true,
	'&':  true,
	'\'': true,
	'*':  true,
	'+':  true,
	'-':  true,
	'.':  true,
	'0':  true,
	'1':  true,
	'2':  true,
	'3':  true,
	'4':  true,
	'5':  true,
	'6':  true,
	'7':  true,
	'8':  true,
	'9':  true,
	'A':  true,
	'B':  true,
	'C':  true,
	'D':  true,
	'E':  true,
	'F':  true,
	'G':  true,
	'H':  true,
	'I':  true,
	'J':  true,
	'K':  true,
	'L':  true,
	'M':  true,
	'N':  true,
	'O':  true,
	'P':  true,
	'Q':  true,
	'R':  true,
	'S':  true,
	'T':  true,
	'U':  true,
	'W':  true,
	'V':  true,
	'X':  true,
	'Y':  true,
	'Z':  true,
	'^':  true,
	'_':  true,
	'`':  true,
	'a':  true,
	'b':  true,
	'c':  true,
	'd':  true,
	'e':  true,
	'f':  true,
	'g':  true,
	'h':  true,
	'i':  true,
	'j':  true,
	'k':  true,
	'l':  true,
	'm':  true,
	'n':  true,
	'o':  true,
	'p':  true,
	'q':  true,
	'r':  true,
	's':  true,
	't':  true,
	'u':  true,
	'v':  true,
	'w':  true,
	'x':  true,
	'y':  true,
	'z':  true,
	'|':  true,
	'~':  true,
}

//go:norace
func skipSpace(s string) (rest string) { _ = "STUB: not implemented"; return "" }

//go:norace
func nextToken(s string) (token, rest string) { _ = "STUB: not implemented"; return "", "" }

//go:norace
func nextTokenOrQuoted(s string) (value string, rest string) {
	_ = "STUB: not implemented"
	return "", ""
}

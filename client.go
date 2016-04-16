package abreuvoir

import (
	"net"

	"github.com/HowardStark/abreuvoir/entry"
	"github.com/HowardStark/abreuvoir/message"
	"github.com/HowardStark/abreuvoir/util"
)

// ClientStatus is the enum type to represent the different
// states/statuses the client could have
type ClientStatus int

const (
	// ClientDisconnected indicates that the client cannot reach
	// the server
	ClientDisconnected ClientStatus = iota
	// ClientConnected indicates that the client has connected to
	// the server but has not began actual communication
	ClientConnected
	// ClientSentHello indicates that the client has sent the hello
	// packets and is waiting for a response from the server
	ClientSentHello
	// ClientInSync indicates that the client is completely in sync
	// with the server and has all the correct values.
	ClientInSync
)

var (
	lastPacket message.Adapter
	lastSent   int
)

// Client is the NetworkTables Client
type Client struct {
	conn    net.Conn
	entries map[string]entry.Adapter
	status  ClientStatus
}

func newClient(connAddr, connPort string) (*Client, error) {
	tcpConn, err := net.Dial("tcp", util.ConcatAddress(connAddr, connPort))
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:    tcpConn,
		entries: map[string]entry.Adapter{},
	}
}

// GetBoolean fetches a boolean at the specified key
func (client *Client) GetBoolean(key string) bool {
	key = util.SanitizeKey(key)
	_ = key
	return true
}

// keepAlive should be run in a Go routine. It sends a
// the provided packet after the provided time (ms) have
// passed between the last packet.
func (client *Client) keepAlive(packet message.Adapter, timeout int) {

}

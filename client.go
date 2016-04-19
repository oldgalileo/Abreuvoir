package abreuvoir

import (
	"errors"
	"io"
	"net"
	"time"

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
	// ClientStartingSync indicates that the client has received the
	// server hello and is beginning to synchronize values with the
	// server.
	ClientStartingSync
	// ClientInSync indicates that the client is completely in sync
	// with the server and has all the correct values.
	ClientInSync
	// keepAliveTime is the amount of time (seconds) between packets
	// that the client waits before it sends a KeepAlive message.
	// It is advised to never have this lower than one second so as to
	// prevent overloading the server.
	keepAliveTime int64 = 1
)

var (
	lastPacket message.Adapter
	lastSent   int64
)

// Client is the NetworkTables Client
type Client struct {
	handler ClientHandler
	conn    net.Conn
	entries map[string]entry.Adapter
	status  ClientStatus
}

func newClient(connAddr, connPort string) (*Client, error) {
	tcpConn, err := net.Dial("tcp", util.ConcatAddress(connAddr, connPort))
	if err != nil {
		return &Client{
			handler: nil,
			conn:    nil,
			entries: map[string]entry.Adapter{},
			status:  ClientDisconnected,
		}, err
	}
	client := Client{
		handler: ClientHandler{
			client,
		},
		conn:    tcpConn,
		entries: map[string]entry.Adapter{},
		status:  ClientConnected,
	}
	defer client.startHandshake()
	return &client, nil
}

// Close disconnects and closes the client from the server.
func (client *Client) Close() error {
	if client.status == ClientDisconnected {
		return errors.New("client: Already disconnected")
	}
	client.status = ClientDisconnected
	client.conn.Close()
	return nil
}

// GetBoolean fetches a boolean at the specified key
func (client *Client) GetBoolean(key string) bool {
	key = util.SanitizeKey(key)
	_ = key
	return true
}

func (client *Client) startHandshake() {
	go client.keepAlive(message.KeepAliveFromItems())
}

// sendMessage
func (client *Client) sendMessage(message message.Adapter) error {
	if client.status != ClientDisconnected {
		client.conn.Write(message.CompressToBytes())
		defer updateLastSent()
		return nil
	}
	return errors.New("client: server could not be reached")
}

// readMessage
func (client *Client) receiveIncoming() {
	var potentialMessage [1]byte
	for client.status != ClientDisconnected {
		_, ioError := io.ReadFull(client.conn, potentialMessage[:])
		if ioError != nil {
			if ioError == io.EOF {
				continue
			}
			panic(ioError)
		}
		tempPacket, messageError := message.BuildFromReader(potentialMessage, client.conn)
		if messageError != nil {
			panic(messageError)
		}
		lastPacket = tempPacket
	}
}

func updateLastSent() {
	currentTime := time.Now()
	lastSent = currentTime.Unix()
}

// keepAlive should be run in a Go routine. It sends a
// the provided packet after the provided time (seconds) have
// passed between the last packet.
func (client *Client) keepAlive(packet message.Adapter) {
	for client.status != ClientDisconnected {
		currentTime := time.Now()
		currentSeconds := currentTime.Unix()
		if (currentSeconds - lastSent) >= keepAliveTime {
			go client.sendMessage(packet)
		}
	}
}

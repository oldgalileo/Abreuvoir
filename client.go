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
	lastPacket          message.Adapter
	lastSent            int64
	messageOutgoingChan = make(chan message.Adapter)
)

// Client is the NetworkTables Client
type Client struct {
	handler ClientMessageHandler
	conn    net.Conn
	entries map[string]entry.Adapter
	status  ClientStatus
}

func newClient(connAddr, connPort string) (*Client, error) {
	tcpConn, err := net.Dial("tcp", util.ConcatAddress(connAddr, connPort))
	if err != nil {
		return &Client{
			conn:    nil,
			entries: map[string]entry.Adapter{},
			status:  ClientDisconnected,
		}, err
	}
	client := Client{
		handler: ClientMessageHandler{
			&client,
		},
		conn:    tcpConn,
		entries: map[string]entry.Adapter{},
		status:  ClientConnected,
	}
	defer client.connect()
	return &client, nil
}

func (client *Client) connect() {
	go client.sendOutgoing()
	go client.startHandshake()
	go client.receiveIncoming()
}

func (client *Client) startHandshake() {
	clientName := []byte(IDENTITY)
	clientLength := util.EncodeULeb128(uint32(len(clientName)))
	clientName = append(clientLength, clientName...)
	helloMessage := message.ClientHelloFromItems(VERSION, clientName)
	client.QueueMessage(helloMessage)
	client.status = ClientSentHello
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

// QueueMessage prepares the message that has been provided for
// sending.
func (client *Client) QueueMessage(message message.Adapter) error {
	if client.status != ClientDisconnected {
		messageOutgoingChan <- message
		return nil
	}
	return errors.New("client: server could not be reached")
}

func (client *Client) sendOutgoing() error {
	for client.status != ClientDisconnected {
		sending := <-messageOutgoingChan
		client.conn.Write(sending.CompressToBytes())
		defer updateLastSent()
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
			client.Close()
		}
		tempPacket, messageError := message.BuildFromReader(potentialMessage, client.conn)
		if messageError != nil {
			client.Close()
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
	for client.status == ClientInSync {
		currentTime := time.Now()
		currentSeconds := currentTime.Unix()
		if (currentSeconds - lastSent) >= keepAliveTime {
			go client.QueueMessage(packet)
		}
	}
}

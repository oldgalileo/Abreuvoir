package abreuvoir

import "github.com/HowardStark/Abreuvoir/message"

// ClientMessageHandler handles incoming messages for the client.
type ClientMessageHandler struct {
	client *Client
}

func (handler *ClientMessageHandler) handleMessage(message message.Adapter) {
	switch message.GetType() {
	case message.TypeClientHello:
		handleClientHello(message)
	case message.TypeProtoUnsupported:
		handleProtoUnsupported(message)
	case message.TypeServerHelloComplete:
		handleServerHelloComplete(message)
	case message.TypeServerHello:
		handleServerHello(message)
	}
}

// handleClientHello should never be triggered on a client, and the client
// should disconnect immediately if it is triggered.
func (handler *ClientMessageHandler) handleClientHello(message message.ClientHello) {
	handler.client.Close()
}

// handleProtoUnsupported should trigger a reconnect using the returned
// version if possible, or close if the version is unsupported by the
// client.
func (handler *ClientMessageHandler) handleProtoUnsupported(message message.ProtoUnsupported) {
	// Abreuvoir doesn't support any versions of NetworkTables earlier than
	// 3.0, so any protocol unsupported messages should cause an immediate
	// disconnect. This should be updated when Abreuvoir implements more versions.
	handler.client.Close()
}

func (handler *ClientMessageHandler) handleServerHelloComplete(message message.ServerHelloComplete) {
	if handler.client.status == ClientStartingSync {
		handler.client.status = ClientInSync
	} else {
		handler.client.Close()
	}
}

func (handler *ClientMessageHandler) handleServerHello(message message.ServerHello) {
	if handler.client.status == ClientSentHello {
		handler.client.status = ClientStartingSync
	} else {
		handler.client.Close()
	}
}

// handleClientHelloComplete should never be triggered on a client, and the client
// should disconnect immediately if it is triggered.
func (handler *ClientMessageHandler) handleClientHelloComplete(message message.ClientHelloComplete) {
	handler.client.Close()
}

func (handler *ClientMessageHandler) handleEntryAssign(message message.EntryAssign) {

}

func (handler *ClientMessageHandler) handleEntryUpdate(message message.EntryUpdate) {

}

func (handler *ClientMessageHandler) handleEntryFlagUpdate(message message.EntryFlagUpdate) {

}

func (handler *ClientMessageHandler) handleEntryDelete(message message.EntryDelete) {

}

func (handler *ClientMessageHandler) handleClearAllEntries() {

}

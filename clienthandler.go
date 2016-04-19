package abreuvoir

import "github.com/HowardStark/Abreuvoir/message"

// ClientMessageHandler handles incoming messages for the client.
type ClientMessageHandler struct {
	client *Client
}

func (handler *ClientMessageHandler) handleMessage(incoming message.Adapter) {
	switch incoming.(type) {
	case *message.ClientHello:
		handler.handleClientHello(*incoming.(*message.ClientHello))
	case *message.ProtoUnsupported:
		handler.handleProtoUnsupported(*incoming.(*message.ProtoUnsupported))
	case *message.ServerHelloComplete:
		handler.handleServerHelloComplete(*incoming.(*message.ServerHelloComplete))
	case *message.ServerHello:
		handler.handleServerHello(*incoming.(*message.ServerHello))
	}
}

// handleClientHello should never be triggered on a client, and the client
// should disconnect immediately if it is triggered.
func (handler *ClientMessageHandler) handleClientHello(incoming message.ClientHello) {
	handler.client.Close()
}

// handleProtoUnsupported should trigger a reconnect using the returned
// version if possible, or close if the version is unsupported by the
// client.
func (handler *ClientMessageHandler) handleProtoUnsupported(incoming message.ProtoUnsupported) {
	// Abreuvoir doesn't support any versions of NetworkTables earlier than
	// 3.0, so any protocol unsupported messages should cause an immediate
	// disconnect. This should be updated when Abreuvoir implements more versions.
	handler.client.Close()
}

func (handler *ClientMessageHandler) handleServerHelloComplete(incoming message.ServerHelloComplete) {
	if handler.client.status == ClientStartingSync {
		handler.client.status = ClientInSync
		handler.client.keepAlive(message.KeepAliveFromItems())
	} else {
		handler.client.Close()
	}
}

func (handler *ClientMessageHandler) handleServerHello(incoming message.ServerHello) {
	if handler.client.status == ClientSentHello {
		handler.client.status = ClientStartingSync
	} else {
		handler.client.Close()
	}
}

// handleClientHelloComplete should never be triggered on a client, and the client
// should disconnect immediately if it is triggered.
func (handler *ClientMessageHandler) handleClientHelloComplete(incoming message.ClientHelloComplete) {
	handler.client.Close()
}

func (handler *ClientMessageHandler) handleEntryAssign(incoming message.EntryAssign) {
	return
}

func (handler *ClientMessageHandler) handleEntryUpdate(incoming message.EntryUpdate) {
	return
}

func (handler *ClientMessageHandler) handleEntryFlagUpdate(incoming message.EntryFlagUpdate) {
	return
}

func (handler *ClientMessageHandler) handleEntryDelete(incoming message.EntryDelete) {
	return
}

func (handler *ClientMessageHandler) handleClearAllEntries() {
	return
}

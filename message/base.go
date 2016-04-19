package message

import (
	"errors"
	"io"
)

const (
	// TypeKeepAlive is the message type for the Keep Alive message
	TypeKeepAlive byte = 0x00
	// TypeClientHello is the message type for the Client Hello message
	TypeClientHello byte = 0x01
	// TypeProtoUnsupported is the message type for the Protocol Unsupported message
	TypeProtoUnsupported byte = 0x02
	// TypeServerHelloComplete is the message type for the Server Hello Complete message
	TypeServerHelloComplete byte = 0x03
	// TypeServerHello is the message type for the Server Hello message
	TypeServerHello byte = 0x04
	// TypeClientHelloComplete is the message type for the Client Hello Complete message
	TypeClientHelloComplete byte = 0x05
	// TypeEntryAssign is the message type for the Entry Assign message
	TypeEntryAssign byte = 0x10
	// TypeEntryUpdate is the message type for the Entry Update message
	TypeEntryUpdate byte = 0x11
	// TypeEntryFlagUpdate is the message type for the Entry Flag Update message
	TypeEntryFlagUpdate byte = 0x12
	// TypeEntryDelete is the message type for the Entry Delete message
	TypeEntryDelete byte = 0x13
	// TypeClearAllEntries is the message type for the Clear All Entries message
	TypeClearAllEntries byte = 0x14
	// TypeRPCExec is the message type for the Remote Procedure Call Execute message
	TypeRPCExec byte = 0x20
	// TypeRPCResponse is the message type for the RPC Response message
	TypeRPCResponse byte = 0x21

	lsbFirstConnect byte = 0x00
	lsbReconnect    byte = 0x01
)

// Base is the base struct for Messages
type Base struct {
	mType byte
	mData []byte
}

// BuildFromReader identifies and builds the message with the same type as the
// message type passed in
func BuildFromReader(messageType [1]byte, reader io.Reader) (Adapter, error) {
	switch messageType[0] {
	case TypeKeepAlive:
		return KeepAliveFromReader(), nil
	case TypeClientHello:
		return ClientHelloFromReader(reader)
	case TypeProtoUnsupported:
		return ProtoUnsupportedFromReader(reader)
	case TypeServerHelloComplete:
		return ServerHelloCompleteFromReader(), nil
	case TypeServerHello:
		return ServerHelloFromReader(reader)
	case TypeClientHelloComplete:
		return ClientHelloCompleteFromReader(), nil
	case TypeEntryAssign:
		return EntryAssignFromReader(reader)
	case TypeEntryUpdate:
		return EntryUpdateFromReader(reader)
	case TypeEntryFlagUpdate:
		return EntryFlagUpdateFromReader(reader)
	case TypeEntryDelete:
		return EntryDeleteFromReader(reader)
	case TypeClearAllEntries:
	case TypeRPCExec:
	case TypeRPCResponse:
	default:
		return nil, errors.New("message: Unknown message type")
	}
	return nil, errors.New("message: Unknown message type")
}

// compressToBytes remakes the original byte slice to represent this entry
func (base *Base) compressToBytes() []byte {
	output := []byte{}
	output = append(output, base.mType)
	output = append(output, base.mData...)
	return output
}

package message

import (
	"errors"
	"io"
)

const (
	typeKeepAlive           byte = 0x00
	typeClientHello         byte = 0x01
	typeProtoUnsupported    byte = 0x02
	typeServerHelloComplete byte = 0x03
	typeServerHello         byte = 0x04
	typeClientHelloComplete byte = 0x05
	typeEntryAssign         byte = 0x10
	typeEntryUpdate         byte = 0x11
	typeEntryFlagUpdate     byte = 0x12
	typeEntryDelete         byte = 0x13
	typeClearAllEntries     byte = 0x14
	typeRPCExec             byte = 0x20
	typeRPCResponse         byte = 0x21

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
	case typeKeepAlive:
		return KeepAliveFromReader(), nil
	case typeClientHello:
		return ClientHelloFromReader(reader)
	case typeProtoUnsupported:
		return ProtoUnsupportedFromReader(reader)
	case typeServerHelloComplete:
		return ServerHelloCompleteFromReader(), nil
	case typeServerHello:
		return ServerHelloFromReader(reader)
	case typeClientHelloComplete:
		return ClientHelloCompleteFromReader(), nil
	case typeEntryAssign:
		return EntryAssignFromReader(reader)
	case typeEntryUpdate:
		return EntryUpdateFromReader(reader)
	case typeEntryFlagUpdate:
		return EntryFlagUpdateFromReader(reader)
	case typeEntryDelete:
		return EntryDeleteFromReader(reader)
	case typeClearAllEntries:
	case typeRPCExec:
	case typeRPCResponse:
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

package message

import (
	"bytes"

	"github.com/HowardStark/abreuvoir/util"
)

// ServerHello message
type ServerHello struct {
	Base
	firstConnection bool
	serverIdentity  string
}

// ServerHelloFromItems builds a new ServerHello message using the provided parameters
func ServerHelloFromItems(flags byte, identity []byte) *ServerHello {
	identityStrLen, identitySizeLen := util.ReadULeb128(bytes.NewBuffer(identity))
	identityStr := string(identity[identitySizeLen:identityStrLen])
	flagsLSB := flags & 1
	firstConn := (flagsLSB == lsbFirstConnect)
	totalData := append([]byte{flags}, identity...)
	return &ServerHello{
		firstConnection: firstConn,
		serverIdentity:  identityStr,
		Base: Base{
			mType: typeServerHello,
			mData: totalData,
		},
	}
}

// IsFirstConnection returns if this is the first connection the client has made to the server
func (serverHello *ServerHello) IsFirstConnection() bool {
	return serverHello.firstConnection
}

// GetServerIdentity returns the server identity
func (serverHello *ServerHello) GetServerIdentity() string {
	return serverHello.serverIdentity
}

// CompressToBytes returns the message in its byte array form
func (serverHello *ServerHello) CompressToBytes() []byte {
	return serverHello.Base.compressToBytes()
}

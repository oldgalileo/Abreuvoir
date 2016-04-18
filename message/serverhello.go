package message

import (
	"bytes"
	"io"

	"github.com/HowardStark/abreuvoir/util"
)

// ServerHello message
type ServerHello struct {
	Base
	firstConnection bool
	serverIdentity  string
}

// ServerHelloFromReader builds a new ServerHello message using the provided reader
func ServerHelloFromReader(reader io.Reader) (*ServerHello, error) {
	var flags [1]byte
	_, flagErr := io.ReadFull(reader, flags[:])
	if flagErr != nil {
		return nil, flagErr
	}
	firstConn := ((flags[0] & 1) == lsbFirstConnect)
	identityLength, identitySizeData := util.PeekULeb128(reader)
	identityData := make([]byte, identityLength)
	_, identityErr := io.ReadFull(reader, identityData[:])
	if identityErr != nil {
		return nil, identityErr
	}
	identity := string(identityData[:])
	var totalData []byte
	totalData = append(totalData, flags[:]...)
	totalData = append(totalData, identitySizeData...)
	totalData = append(totalData, identityData[:]...)
	return &ServerHello{
		firstConnection: firstConn,
		serverIdentity:  identity,
		Base: Base{
			mType: typeServerHello,
			mData: totalData,
		},
	}, nil
}

// ServerHelloFromItems builds a new ServerHello message using the provided parameters
func ServerHelloFromItems(flags byte, identity []byte) *ServerHello {
	identityStrLen, identitySizeLen := util.ReadULeb128(bytes.NewBuffer(identity))
	identityStr := string(identity[identitySizeLen:identityStrLen])
	firstConn := ((flags & 1) == lsbFirstConnect)
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

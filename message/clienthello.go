package message

import (
	"bytes"

	"github.com/HowardStark/abreuvoir/util"
)

// ClientHello message
type ClientHello struct {
	Base
	protoRev [2]byte
	identity string
}

// ClientHelloFromItems builds a new ClientHello message using the provided parameters
func ClientHelloFromItems(protocolRev [2]byte, nameData []byte) *ClientHello {
	nameLen, sizeLen := util.ReadULeb128(bytes.NewBuffer(nameData))
	name := string(nameData[sizeLen:nameLen])
	totalData := []byte{}
	totalData = append(totalData, protocolRev[:2]...)
	totalData = append(totalData, nameData[:]...)
	return &ClientHello{
		identity: name,
		protoRev: protocolRev,
		Base: Base{
			mType: typeClientHello,
			mData: totalData,
		},
	}
}

// GetProtoRev returns the client's NetworkTable protocol revision
func (clientHello *ClientHello) GetProtoRev() [2]byte {
	return clientHello.protoRev
}

// GetIdentity returns the client's identity
func (clientHello *ClientHello) GetIdentity() string {
	return clientHello.identity
}

// CompressToBytes returns the message in its byte array form
func (clientHello *ClientHello) CompressToBytes() []byte {
	return clientHello.Base.compressToBytes()
}

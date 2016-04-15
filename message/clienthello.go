package message

import (
	"bytes"

	"github.com/HowardStark/Abreuvoir/util"
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
	totalData := []byte{protocolRev[:2], nameData}
	return &ClientHello{
		identity: name,
		proto:    protoRev,
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

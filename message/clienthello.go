package message

import (
	"bytes"

	"github.com/HowardStark/Abreuvoir/util"
)

// ClientHello message
type ClientHello struct {
	Base
	protocolRev [2]byte
	identity    string
}

// ClientHelloFromItems builds a new ClientHello message using the provided parameters
func ClientHelloFromItems(data []byte) *ClientHello {
	var protoRev [2]byte
	copy(protoRev[:], data[:2])
	nameData := data[2:]
	nameLen, sizeLen := util.ReadULeb128(bytes.NewBuffer(nameData))
	name := string(nameData[sizeLen:nameLen])
	return &ClientHello{
		identity:    name,
		protocolRev: protoRev,
		Base: Base{
			mType: typeClientHello,
			mData: data,
		},
	}
}

package message

import (
	"bytes"

	"github.com/HowardStark/Abreuvoir/util"
)

// ClientHello message
type ClientHello struct {
	Base
	identity string
}

// ClientHelloFromItems builds a new ClientHello message using the provided parameters
func ClientHelloFromItems(protoRev [2]byte, clientIdentity []byte) *ClientHello {
	nameLen, sizeLen := util.ReadULeb128(bytes.NewBuffer(clientIdentity))
	name := string(clientIdentity[sizeLen:nameLen])
	return &ClientHello{
		identity: name,
		Base: Base{
			mType: typeClientHello,
			mData: clientIdentity,
		},
	}
}

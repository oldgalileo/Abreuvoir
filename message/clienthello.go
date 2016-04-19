package message

import (
	"bytes"
	"io"

	"github.com/HowardStark/abreuvoir/util"
)

// ClientHello message
type ClientHello struct {
	Base
	protoRev [2]byte
	identity string
}

// ClientHelloFromReader builds a new ClientHello message using the provided io.Reader
func ClientHelloFromReader(reader io.Reader) (Adapter, error) {
	var protocolRev [2]byte
	_, err := io.ReadFull(reader, protocolRev[:])
	if err != nil {
		return nil, err
	}
	nameLen, sizeData := util.PeekULeb128(reader)
	nameData := make([]byte, nameLen)
	_, err = io.ReadFull(reader, nameData[:])
	if err != nil {
		return nil, err
	}
	name := string(nameData[:])
	var totalData []byte
	totalData = append(totalData, protocolRev[:]...)
	totalData = append(totalData, sizeData...)
	totalData = append(totalData, nameData[:]...)
	return &ClientHello{
		identity: name,
		protoRev: protocolRev,
		Base: Base{
			mType: TypeClientHello,
			mData: totalData,
		},
	}, nil
}

// ClientHelloFromItems builds a new ClientHello message using the provided parameters
func ClientHelloFromItems(protocolRev [2]byte, nameData []byte) *ClientHello {
	nameLen, sizeLen := util.ReadULeb128(bytes.NewBuffer(nameData))
	name := string(nameData[sizeLen:nameLen])
	var totalData []byte
	totalData = append(totalData, protocolRev[:]...)
	totalData = append(totalData, nameData[:]...)
	return &ClientHello{
		identity: name,
		protoRev: protocolRev,
		Base: Base{
			mType: TypeClientHello,
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

// GetType returns the message's type
func (clientHello *ClientHello) GetType() byte {
	return TypeClientHello
}

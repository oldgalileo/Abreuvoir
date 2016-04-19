package message

import "io"

// ProtoUnsupported message
type ProtoUnsupported struct {
	supportedProto [2]byte
	Base
}

// ProtoUnsupportedFromReader builds a new ProtoUnsupported message
func ProtoUnsupportedFromReader(reader io.Reader) (*ProtoUnsupported, error) {
	var supportedVersion [2]byte
	_, err := io.ReadFull(reader, supportedVersion[:])
	if err != nil {
		return nil, err
	}
	return &ProtoUnsupported{
		supportedProto: supportedVersion,
		Base: Base{
			mType: TypeProtoUnsupported,
			mData: supportedVersion[:],
		},
	}, nil
}

// ProtoUnsupportedFromItems builds a new ProtoUnsupported message
func ProtoUnsupportedFromItems(data [2]byte) *ProtoUnsupported {
	return &ProtoUnsupported{
		supportedProto: data,
		Base: Base{
			mType: TypeProtoUnsupported,
			mData: data[:],
		},
	}
}

// GetSupportedProto returns the NetworkTables protocol revision that the server supports
func (protoUnsupported *ProtoUnsupported) GetSupportedProto() [2]byte {
	return protoUnsupported.supportedProto
}

// CompressToBytes returns the message in its byte array form
func (protoUnsupported *ProtoUnsupported) CompressToBytes() []byte {
	return protoUnsupported.Base.compressToBytes()
}

// GetType returns the message's type
func (protoUnsupported *ProtoUnsupported) GetType() byte {
	return TypeProtoUnsupported
}

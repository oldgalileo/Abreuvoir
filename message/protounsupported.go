package message

// ProtoUnsupported message
type ProtoUnsupported struct {
	supportedProto [2]byte
	Base
}

// ProtoUnsupportedFromItems builds a new ProtoUnsupported message
func ProtoUnsupportedFromItems(data [2]byte) *ProtoUnsupported {
	return &ProtoUnsupported{
		supportedProto: data,
		Base: Base{
			mType: typeProtoUnsupported,
			mData: data[:],
		},
	}
}

// GetSupportedProto returns the NetworkTables protocol revision that the server supports
func (protoUnsupported *ProtoUnsupported) GetSupportedProto() [2]byte {
	return protoUnsupported.supportedProto
}

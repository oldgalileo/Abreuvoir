package messages

// ProtoUnsupported message
type ProtoUnsupported struct {
	supportedProto [2]byte
	Base
}

// ProtoUnsupportedFromItems builds a new ProtoUnsupported message
func ProtoUnsupportedFromItems(data []byte) *ProtoUnsupported{
    var serverVersion [2]byte
    copy(serverVersion, data[:2])
    return &ProtoUnsupported{
        supportedProto: serverVersion,
        Base: Base{
            mType: typeProtoUnsupported,
            mData: data[:2],
        }
    }
}

func (protoUnsupported *ProtoUnsupported) GetValue() [2]byte {
    return protoUnsupported.supportedProto
}

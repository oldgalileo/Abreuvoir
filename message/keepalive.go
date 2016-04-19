package message

// KeepAlive message
type KeepAlive struct {
	Base
}

// KeepAliveFromReader builds a new KeepAlive message
func KeepAliveFromReader() *KeepAlive {
	return &KeepAlive{
		Base: Base{
			mType: TypeKeepAlive,
			// KeepAlive has no body data
			mData: []byte{},
		},
	}
}

// KeepAliveFromItems builds a new KeepAlive message
func KeepAliveFromItems() *KeepAlive {
	return &KeepAlive{
		Base: Base{
			mType: TypeKeepAlive,
			// KeepAlive has no body data
			mData: []byte{},
		},
	}
}

// CompressToBytes returns the message in its byte array form
func (keepAlive *KeepAlive) CompressToBytes() []byte {
	return keepAlive.Base.compressToBytes()
}

// GetType returns the message's type
func (keepAlive *KeepAlive) GetType() byte {
	return TypeKeepAlive
}

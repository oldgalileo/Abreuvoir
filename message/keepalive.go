package message

// KeepAlive message
type KeepAlive struct {
	Base
}

// KeepAliveFromItems builds a new KeepAlive message
func KeepAliveFromItems() *KeepAlive {
	return &KeepAlive{
		Base: Base{
			mType: typeKeepAlive,
			// KeepAlive has no body data
			mData: []byte{},
		},
	}
}

// CompressToBytes returns the message in its byte array form
func (keepAlive *KeepAlive) CompressToBytes() []byte {
	return keepAlive.Base.compressToBytes()
}

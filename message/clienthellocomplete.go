package message

// ClientHelloComplete message
type ClientHelloComplete struct {
	Base
}

// ClientHelloCompleteFromReader builds a new ClientHelloComplete message
func ClientHelloCompleteFromReader() *ClientHelloComplete {
	return &ClientHelloComplete{
		Base: Base{
			mType: typeClientHelloComplete,
			// ClientHelloComplete has no body data
			mData: []byte{},
		},
	}
}

// ClientHelloCompleteFromItems builds a new ClientHelloComplete message
func ClientHelloCompleteFromItems() *ClientHelloComplete {
	return &ClientHelloComplete{
		Base: Base{
			mType: typeClientHelloComplete,
			// ClientHelloComplete has no body data
			mData: []byte{},
		},
	}
}

// CompressToBytes returns the message in its byte array form
func (clientHelloComplete *ClientHelloComplete) CompressToBytes() []byte {
	return clientHelloComplete.Base.compressToBytes()
}

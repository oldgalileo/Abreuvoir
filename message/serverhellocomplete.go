package message

// ServerHelloComplete message
type ServerHelloComplete struct {
	Base
}

// ServerHelloCompleteFromReader builds a new ServerHelloComplete message using the provided parameters
func ServerHelloCompleteFromReader() *ServerHelloComplete {
	return &ServerHelloComplete{
		Base: Base{
			mType: typeServerHelloComplete,
			// ServerHelloComplete has no body data
			mData: []byte{},
		},
	}
}

// ServerHelloCompleteFromItems builds a new ServerHelloComplete message using the provided parameters
func ServerHelloCompleteFromItems() *ServerHelloComplete {
	return &ServerHelloComplete{
		Base: Base{
			mType: typeServerHelloComplete,
			// ServerHelloComplete has no body data
			mData: []byte{},
		},
	}
}

// CompressToBytes returns the message in its byte array form
func (serverHelloComplete *ServerHelloComplete) CompressToBytes() []byte {
	return serverHelloComplete.Base.compressToBytes()
}

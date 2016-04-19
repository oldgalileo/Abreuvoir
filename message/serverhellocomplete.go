package message

// ServerHelloComplete message
type ServerHelloComplete struct {
	Base
}

// ServerHelloCompleteFromReader builds a new ServerHelloComplete message using the provided parameters
func ServerHelloCompleteFromReader() *ServerHelloComplete {
	return &ServerHelloComplete{
		Base: Base{
			mType: TypeServerHelloComplete,
			// ServerHelloComplete has no body data
			mData: []byte{},
		},
	}
}

// ServerHelloCompleteFromItems builds a new ServerHelloComplete message using the provided parameters
func ServerHelloCompleteFromItems() *ServerHelloComplete {
	return &ServerHelloComplete{
		Base: Base{
			mType: TypeServerHelloComplete,
			// ServerHelloComplete has no body data
			mData: []byte{},
		},
	}
}

// CompressToBytes returns the message in its byte array form
func (serverHelloComplete *ServerHelloComplete) CompressToBytes() []byte {
	return serverHelloComplete.Base.compressToBytes()
}

// GetType returns the message's type
func (serverHelloComplete *ServerHelloComplete) GetType() byte {
	return TypeServerHelloComplete
}

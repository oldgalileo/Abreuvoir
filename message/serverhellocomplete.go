package message

// ServerHelloComplete message
type ServerHelloComplete struct {
	Base
}

// ServerHelloCompleteFromItems builds
func ServerHelloCompleteFromItems() *ServerHelloComplete {
	return &ServerHelloComplete{
		Base: Base{
			mType: typeServerHelloComplete,
			// ServerHelloComplete has no body data
			mData: []byte{},
		},
	}
}

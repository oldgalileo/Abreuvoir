package message

// ClientHelloComplete message
type ClientHelloComplete struct {
	Base
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

package message

type ClientHelloComplete struct {
	Base
}

// ClientHelloCompleteFromItems builds a new
func ClientHelloCompleteFromItems() *ClientHelloComplete {
	return &ClientHelloComplete{
		Base: Base{
			mType: typeClientHelloComplete,
			mData: []byte{},
		},
	}
}

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
			mData: []byte{},
		},
	}
}

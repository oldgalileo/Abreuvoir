package message

// KeepAlive message
type KeepAlive struct {
	Base
}

// BuildKeepAlive builds a new KeepAlive message
func BuildKeepAlive() *KeepAlive {
	return &KeepAlive{
		Base: Base{
			mType: typeKeepAlive,
			mData: []byte{},
		},
	}
}

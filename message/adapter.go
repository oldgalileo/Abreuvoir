package message

// Adapter is the Message interface
type Adapter interface {
	CompressToBytes() []byte
}

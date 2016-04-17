package message

// Adapter is the Message interface
type Adapter interface {
	getType() byte
	getData() []byte
	CompressToBytes() []byte
}

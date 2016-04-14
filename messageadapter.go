package abreuvoir

// MessageAdapter is the Message interface
type MessageAdapter interface {
	getType() byte
	getData() []byte
	composeMessage() []byte
}

package entryupdate

// Adapter is the entry update interface
type Adapter interface {
	CompressToBytes() []byte
}

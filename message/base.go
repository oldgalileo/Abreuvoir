package message

const (
	typeKeepAlive           byte = 0x00
	typeClientHello         byte = 0x01
	typeProtoUnsupported    byte = 0x02
	typeServerHelloComplete byte = 0x03
	typeServerHello         byte = 0x04
	typeClientComplete      byte = 0x05
	typeEntryAssign         byte = 0x10
	typeEntryUpdate         byte = 0x11
	typeEntryFlagUpdate     byte = 0x12
	typeEntryDelete         byte = 0x13
	typeClearAllEntries     byte = 0x14
	typeRPCExec             byte = 0x20
	typeRPCResponse         byte = 0x21
)

// Base is the base struct for Messages
type Base struct {
	mType byte
	mData []byte
}

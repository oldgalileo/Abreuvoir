package message

import (
	"bytes"

	"github.com/HowardStark/abreuvoir/util"
)

// ServerHello message
type ServerHello struct {
	Base
	firstConnection bool
	serverIdentity  string
}

// ServerHelloFromItems builds a new ServerHello message using the provided parameters
func ServerHelloFromItems(flags byte, identity []byte) *ServerHello {
	identityStrLen, identitySizeLen := util.ReadULeb128(bytes.NewBuffer(identity))
	identityStr := string(identity[identitySizeLen:identityStrLen])
	flagsLSB := flags & 1
	firstConn := (flagsLSB == lsbFirstConnect)
	totalData := append([]byte{flags}, identity)
	return &ServerHello{
		firstConnection: firstConn,
		serverIdentity:  identityStr,
		Base: Base{
			mType: typeServerHello,
			mData: totalData,
		},
	}
}

func (serverHello *ServerHello) IsFirstConnection() bool {
	return serverHello.firstConnection
}

func (serverHello *ServerHell) GetServerIdentity() string {
	return serverHello.serverIdentity
}

package message

import (
	"bytes"

	"github.com/HowardStark/Abreuvoir/util"
)

type ServerHello struct {
	Base
	firstConnection bool
	serverIdentity  string
}

func ServerHelloFromItems(flags byte, identity []byte) *ServerHello {
	identityStrLen, identitySizeLen := util.ReadULeb128(bytes.NewBuffer(identity))
	identityStr := string(identity[identitySizeLen:identityStrLen])
	flagsLSB := flags&1
	firstConn := (flagsLSB == lsbFirstConnect)
    return &ServerHello{
        firstConnection: firstConn,
        serverIdentity: identityStr,
        Base: Base{
            mType: typeServerHello,
            mData: []byte{flags, identity...}
        }
    }
}

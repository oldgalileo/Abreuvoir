package entry

import (
	"bytes"

	"github.com/HowardStark/abreuvoir/util"
)

const (
	typeBoolean    byte = 0x00
	typeDouble     byte = 0x01
	typeString     byte = 0x02
	typeRaw        byte = 0x03
	typeBooleanArr byte = 0x10
	typeDoubleArr  byte = 0x11
	typeStringArr  byte = 0x12
	typeRPCDef     byte = 0x20

	flagTemporary byte = 0x00
	flagPersist   byte = 0x01
	flagReserved  byte = 0xFE
)

var (
	// idSent is the required ID for an entry that is being created/sent from the client
	idSent = [2]byte{0xFF, 0xFF}
)

// Base is the base struct for entries.
type Base struct {
	eName    string
	eNameLen uint32
	eType    byte
	eID      [2]byte
	eSeq     [2]byte
	eFlag    byte
	eValue   []byte
}

// BuildFromRaw creates a Base using the data passed in.
func BuildFromRaw(data []byte) Adapter {
	nameLen, sizeLen := util.ReadULeb128(bytes.NewReader(data))
	dName := string(data[sizeLen : nameLen-1])
	dType := data[nameLen]
	dID := [2]byte{data[nameLen+1], data[nameLen+2]}
	dSeq := [2]byte{data[nameLen+3], data[nameLen+4]}
	dFlag := data[nameLen+5]
	dValue := data[nameLen+6:]
	switch dType {
	case typeBoolean:
		return BooleanFromItems(dName, dID, dSeq, dFlag, dValue)
	case typeDouble:
		return DoubleFromItems(dName, dID, dSeq, dFlag, dValue)
	case typeString:
		return StringFromItems(dName, dID, dSeq, dFlag, dValue)
	case typeRaw:
		return RawFromItems(dName, dID, dSeq, dFlag, dValue)
	case typeBooleanArr:
		return BooleanArrFromItems(dName, dID, dSeq, dFlag, dValue)
	case typeDoubleArr:
		return DoubleArrFromItems(dName, dID, dSeq, dFlag, dValue)
	case typeStringArr:
		return StringArrFromItems(dName, dID, dSeq, dFlag, dValue)
	}
}

// CompressToBytes remakes the original byte array to represent this entry
func (base *Base) CompressToBytes() []byte {
	output := []byte{}
	nameBytes := []byte(base.eName)
	output = append(output, nameBytes...)
	output = append(output, base.eType)
	output = append(output, base.eID...)
	output = append(output, base.eSeq...)
	output = append(output, base.eFlag)
	output = append(output, base.eValue...)
	return output
}

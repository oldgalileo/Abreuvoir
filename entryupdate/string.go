package entryupdate

import (
	"bytes"
	"io"

	"github.com/HowardStark/abreuvoir/util"
)

// String Entry
type String struct {
	Base
	trueValue string
}

// StringFromReader builds a string entry using the provided parameters
func StringFromReader(id [2]byte, sequence [2]byte, etype byte, reader io.Reader) (*String, error) {
	valLen, sizeData := util.PeekULeb128(reader)
	valData := make([]byte, valLen)
	_, err := io.ReadFull(reader, valData[:])
	if err != nil {
		return nil, err
	}
	val := string(valData[:])
	value := append(sizeData, valData[:]...)
	return &String{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeString,
			Value: value,
		},
	}, nil
}

// StringFromItems builds a string entry using the provided parameters
func StringFromItems(id [2]byte, sequence [2]byte, etype byte, value []byte) *String {
	valLen, sizeLen := util.ReadULeb128(bytes.NewReader(value))
	val := string(value[sizeLen : valLen-1])
	return &String{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeString,
			Value: value,
		},
	}
}

// GetValue returns the value of the String
func (stringEntry *String) GetValue() interface{} {
	return stringEntry.trueValue
}

// Clone returns an identical entry
func (stringEntry *String) Clone() *String {
	return &String{
		trueValue: stringEntry.trueValue,
		Base:      stringEntry.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the String entry
func (stringEntry *String) CompressToBytes() []byte {
	return stringEntry.Base.compressToBytes()
}

package entryupdate

import (
	"bytes"
	"io"

	"github.com/HowardStark/abreuvoir/util"
)

// Raw entry
type Raw struct {
	Base
	trueValue []byte
}

// RawFromReader builds a raw entry using the provided parameters
func RawFromReader(id [2]byte, sequence [2]byte, etype byte, reader io.Reader) (*Raw, error) {
	valLen, sizeData := util.PeekULeb128(reader)
	valData := make([]byte, valLen)
	_, err := io.ReadFull(reader, valData[:])
	if err != nil {
		return nil, err
	}
	value := append(sizeData, valData[:]...)
	return &Raw{
		trueValue: valData[:],
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeRaw,
			Value: value,
		},
	}, nil
}

// RawFromItems builds a raw entry using the provided parameters
func RawFromItems(id [2]byte, sequence [2]byte, etype byte, value []byte) *Raw {
	valLen, sizeLen := util.ReadULeb128(bytes.NewReader(value))
	val := value[sizeLen : valLen-1]
	return &Raw{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeRaw,
			Value: value,
		},
	}
}

// GetValue returns the raw value of this entry
func (raw *Raw) GetValue() interface{} {
	return raw.trueValue
}

// Clone returns an identical entry
func (raw *Raw) Clone() *Raw {
	return &Raw{
		trueValue: raw.trueValue,
		Base:      raw.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the Raw entry
func (raw *Raw) CompressToBytes() []byte {
	return raw.Base.compressToBytes()
}

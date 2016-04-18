package entryupdate

import (
	"bytes"
	"io"

	"github.com/HowardStark/abreuvoir/util"
)

// StringArr Entry
type StringArr struct {
	Base
	trueValue    []string
	isPersistant bool
}

// StringArrFromReader builds a StringArr entry using the provided parameters
func StringArrFromReader(id [2]byte, sequence [2]byte, etype byte, reader io.Reader) (*StringArr, error) {
	var value []byte
	var tempValSize [1]byte
	_, sizeErr := io.ReadFull(reader, tempValSize[:])
	if sizeErr != nil {
		return nil, sizeErr
	}
	value = append(value, tempValSize[0])
	valSize := int(tempValSize[0])
	var val []string
	for counter := 0; counter < valSize; counter++ {
		strLen, sizeData := util.PeekULeb128(reader)
		value = append(value, sizeData...)
		strData := make([]byte, strLen)
		_, strErr := io.ReadFull(reader, strData[:])
		if strErr != nil {
			return nil, strErr
		}
		value = append(value, strData[:]...)
		val = append(val, string(strData[:]))
	}
	return &StringArr{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeStringArr,
			Value: value,
		},
	}, nil
}

// StringArrFromItems builds a StringArr entry using the provided parameters
func StringArrFromItems(id [2]byte, sequence [2]byte, etype byte, value []byte) *StringArr {
	valSize := int(value[0])
	var val []string
	var previousPos uint32 = 1
	for counter := 0; counter < valSize; counter++ {
		strPos, sizePos := util.ReadULeb128(bytes.NewReader(value[previousPos:]))
		strPos += previousPos
		sizePos += previousPos
		tempVal := string(value[sizePos : strPos-1])
		val = append(val, tempVal)
		previousPos = strPos - 1
	}
	return &StringArr{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeStringArr,
			Value: value,
		},
	}
}

// GetValue returns the value of the StringArr
func (stringArr *StringArr) GetValue() interface{} {
	return stringArr.trueValue
}

// GetValueAtIndex returns the value at the specified index
func (stringArr *StringArr) GetValueAtIndex(index int) string {
	return stringArr.trueValue[index]
}

// Clone returns an identical entry
func (stringArr *StringArr) Clone() *StringArr {
	return &StringArr{
		trueValue:    stringArr.trueValue,
		isPersistant: stringArr.isPersistant,
		Base:         stringArr.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the StringArr entry
func (stringArr *StringArr) CompressToBytes() []byte {
	return stringArr.Base.compressToBytes()
}

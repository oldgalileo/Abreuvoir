package entry

import (
	"bytes"
	"encoding/binary"

	"github.com/HowardStark/Abreuvoir/util"
)

// StringArr Entry
type StringArr struct {
	Base
	trueValue    []string
	isPersistant bool
}

// StringArrFromItems builds a StringArr entry using the provided parameters
func StringArrFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *StringArr {
	valSize := binary.BigEndian.Uint64(value[0:1])
	var val []string
	previousPos := 1
	for counter := 0; counter < valSize; counter++ {
		strPos, sizePos := util.ReadULeb128(bytes.NewReader(value[previousPos:]))
		strPos += previousPos
		sizePos += previousPos
		tempVal := string(data[sizePos : strPos-1])
		val = append(val, tempVal)
		previousPos = strPos - 1
	}
	var persistant bool
	if persist == flagPersist {
		persistant = true
	} else {
		persistant = false
	}
	return &StringArr{
		trueValue:    val,
		isPersistant: persistant,
		Base: Base{
			eName:  name,
			eType:  typeStringArr,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the value of the StringArr
func (stringArr *StringArr) GetValue() []string {
	return StringArr.trueValue
}

// GetValueAtIndex returns the value at the specified index
func (stringArr *StringArr) GetValueAtIndex(index int) string {
	return stringArr.trueValue[index]
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (stringArr *StringArr) IsPersistant() bool {
	return stringArr.isPersistant
}

// Clone returns an identical entry
func (stringArr *StringArr) Clone() *StringArr {
	return &StringArr{
		trueValue:    stringArr.trueValue,
		isPersistant: stringArr.isPersistant,
		Base:         stringArr.Base.clone(),
	}
}

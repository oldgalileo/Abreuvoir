package entry

import "encoding/binary"

// BooleanArr Entry
type BooleanArr struct {
	Base
	trueValue    []bool
	isPersistant bool
}

// BooleanArrFromItems builds a boolean array entry using the provided parameters
func BooleanArrFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *BooleanArr {
	valSize := binary.BigEndian.Uint64(value[0:1])
	var val []bool
	for counter := 1; counter-1 < valSize; counter++ {
		var tempVal bool
		if value[counter] == boolTrue {
			tempVal = true
		} else {
			tempVal = false
		}
		val = append(val, tempVal)
	}
	var persistant bool
	if persist == flagPersist {
		persistant = true
	} else {
		persistant = false
	}
	return &BooleanArr{
		trueValue:    val,
		isPersistant: persistant,
		Base: Base{
			eName:  name,
			eType:  typeBooleanArr,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the entire boolean array
func (booleanArr *BooleanArr) GetValue() interface{} {
	return booleanArr.trueValue
}

// GetValueAtIndex returns the value at the specified index
func (booleanArr *BooleanArr) GetValueAtIndex(index int) bool {
	return booleaArr.trueValue[index]
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (booleanArr *BooleanArr) IsPersistant() bool {
	return booleanArr.isPersistant
}

// Clone returns an identical entry
func (booleanArr *BooleanArr) Clone() *BooleanArr {
	return &BooleanArr{
		trueValue:    booleanArr.trueValue,
		isPersistant: booleanArr.isPersistant,
		Base:         booleanArr.Base.clone(),
	}
}

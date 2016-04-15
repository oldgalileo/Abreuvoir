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
	var counter uint64
	for counter = 1; counter-1 < valSize; counter++ {
		tempVal := (value[counter] == boolTrue)
		val = append(val, tempVal)
	}
	persistant := (persist == flagPersist)
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
	return booleanArr.trueValue[index]
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
		Base:         *booleanArr.Base.clone(),
	}
}

// CompressToBytes returns a byte array representing the BooleanArr entry
func (booleanArr *BooleanArr) CompressToBytes() []byte {
	return booleanArr.Base.CompressToBytes()
}

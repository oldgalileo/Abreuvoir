package entry

const (
	boolFalse byte = 0x00
	boolTrue  byte = 0x01
)

// BooleanArr Entry
type BooleanArr struct {
	Base
	trueValue    []bool
	isPersistant bool
}

// BooleanArrFromItems builds a boolean array entry using the provided parameters
func BooleanArrFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *BooleanArr {
	var val []bool
	for counter := 1; counter < len(value)-1; counter++ {
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
func (booleanArr *BooleanArr) GetValue() []bool {
	return booleanArr.trueValue
}

// GetValueAtIndex returns the value at the specified index
func (booleanArr *BooleanArr) GetValueAtIndex(index int) bool {
	return booleaArr.trueValue[index]
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (boolean *Boolean) IsPersistant() bool {
	return boolean.isPersistant
}

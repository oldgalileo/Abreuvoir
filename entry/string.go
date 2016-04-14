package entry

import (
	"bytes"

	"github.com/HowardStark/Abreuvoir/util"
)

// String Entry
type String struct {
	Base
	trueValue    string
	isPersistant bool
}

// StringFromItems builds a string entry using the provided parameters
func StringFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *Double {
	val, _ := util.ReadULeb128(bytes.NewReader(value))
	var persistant bool
	if persist == flagPersist {
		persistant = true
	} else {
		persistant = false
	}
	return &String{
		trueValue:    val,
		isPersistant: persistant,
		Base: Base{
			eName:  name,
			eType:  typeString,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the value of the String
func (stringEntry *String) GetValue() string {
	return stringEntry.trueValue
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (stringEntry *String) IsPersistant() bool {
	return stringEntry.isPersistant
}

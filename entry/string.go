package entry

import (
	"bytes"

	"github.com/HowardStark/abreuvoir/util"
)

// String Entry
type String struct {
	Base
	trueValue    string
	isPersistant bool
}

// StringFromItems builds a string entry using the provided parameters
func StringFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *String {
	nameLen, sizeLen := util.ReadULeb128(bytes.NewReader(value))
	val := string(value[sizeLen : nameLen-1])
	persistant := (persist == flagPersist)
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
func (stringEntry *String) GetValue() interface{} {
	return stringEntry.trueValue
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (stringEntry *String) IsPersistant() bool {
	return stringEntry.isPersistant
}

// Clone returns an identical entry
func (stringEntry *String) Clone() *String {
	return &String{
		trueValue:    stringEntry.trueValue,
		isPersistant: stringEntry.isPersistant,
		Base:         *stringEntry.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the String entry
func (stringEntry *String) CompressToBytes() []byte {
	return stringEntry.Base.compressToBytes()
}

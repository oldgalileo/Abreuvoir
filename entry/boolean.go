package entry

import "io"

// Boolean Entry
type Boolean struct {
	Base
	trueValue    bool
	isPersistant bool
}

// BooleanFromReader builds a boolean entry using the provided parameters
func BooleanFromReader(name string, id [2]byte, sequence [2]byte, persist byte, reader io.Reader) (*Boolean, error) {
	var value [1]byte
	_, err := io.ReadFull(reader, value[:])
	if err != nil {
		return nil, err
	}
	return BooleanFromItems(name, id, sequence, persist, value[:]), nil
}

// BooleanFromItems builds a boolean entry using the provided parameters
func BooleanFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *Boolean {
	val := (value[0] == boolTrue)
	persistant := (persist == flagPersist)
	return &Boolean{
		trueValue:    val,
		isPersistant: persistant,
		Base: Base{
			eName:  name,
			eType:  typeBoolean,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the value of the Boolean
func (boolean *Boolean) GetValue() interface{} {
	return boolean.trueValue
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (boolean *Boolean) IsPersistant() bool {
	return boolean.isPersistant
}

// Clone returns an identical entry
func (boolean *Boolean) Clone() *Boolean {
	return &Boolean{
		trueValue:    boolean.trueValue,
		isPersistant: boolean.isPersistant,
		Base:         boolean.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the Boolean entry
func (boolean *Boolean) CompressToBytes() []byte {
	return boolean.Base.compressToBytes()
}

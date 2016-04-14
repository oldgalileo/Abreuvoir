package entry

// Boolean Entry
type Boolean struct {
	Base
	trueValue    bool
	isPersistant bool
}

// BooleanFromItems builds a boolean entry using the provided parameters
func BooleanFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *Boolean {
	var val bool
	if value[0] == boolTrue {
		val = true
	} else {
		val = false
	}
	var persistant bool
	if persist == flagPersist {
		persistant = true
	} else {
		persistant = false
	}
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
		Base:         *boolean.Base.clone(),
	}
}

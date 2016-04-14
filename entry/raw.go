package entry

// Raw entry
type Raw struct {
	Base
	trueValue    []byte
	isPersistant bool
}

// RawFromItems builds a raw entry using the provided parameters
func RawFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *Raw {
	var persistant bool
	if persist == flagPersist {
		persistant = true
	} else {
		persistant = false
	}
	return &Raw{
		trueValue:    value,
		isPersistant: persistant,
		Base: Base{
			eName:  name,
			eType:  typeRaw,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the raw value of this entry
func (raw *Raw) GetValue() []byte {
	return raw.trueValue
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (raw *Raw) IsPersistant() bool {
	return raw.isPersistant
}

// Clone returns an identical entry
func (raw *Raw) Clone() *Raw {
	return &Raw{
		trueValue:    raw.trueValue,
		isPersistant: raw.isPersistant,
		Base:         raw.Base.clone(),
	}
}

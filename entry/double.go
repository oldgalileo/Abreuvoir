package entry

import "github.com/HowardStark/Abreuvoir/util"

// Double Entry
type Double struct {
	Base
	trueValue    float64
	isPersistant bool
}

// DoubleFromItems builds a double entry using the provided parameters
func DoubleFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *Double {
	val := util.BytesToFloat64(value)
	var persistant bool
	if persist == flagPersist {
		persistant = true
	} else {
		persistant = false
	}
	return &Double{
		trueValue:    val,
		isPersistant: persistant,
		Base: Base{
			eName:  name,
			eType:  typeDouble,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the value of the Double
func (double *Double) GetValue() float64 {
	return double.trueValue
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (double *Double) IsPersistant() bool {
	return double.isPersistant
}

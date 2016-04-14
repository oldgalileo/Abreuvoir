package entry

import "github.com/HowardStark/Abreuvoir/util"

// Double Entry
type Double struct {
	Base
	trueValue    bool
	isPersistant bool
}

// DoubleFromItems builds a boolean entry using the provided parameters
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

// GetValue returns the value of the Boolean
func (double *Double) GetValue() float64 {
	return double.trueValue
}

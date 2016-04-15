package entry

import (
	"encoding/binary"

	"github.com/HowardStark/abreuvoir/util"
)

// DoubleArr Entry
type DoubleArr struct {
	Base
	trueValue    []float64
	isPersistant bool
}

// DoubleArrFromItems builds a DoubleArr entry using the provided parameters
func DoubleArrFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *DoubleArr {
	valSize := binary.BigEndian.Uint64(value[0:1])
	var val []float64
	var counter uint64
	for counter = 1; (counter-1)/8 < valSize; counter += 8 {
		tempVal := util.BytesToFloat64(value[counter : counter+8])
		val = append(val, tempVal)
	}
	persistant := (persist == flagPersist)
	return &DoubleArr{
		trueValue:    val,
		isPersistant: persistant,
		Base: Base{
			eName:  name,
			eType:  typeDoubleArr,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the value of the DoubleArr
func (doubleArr *DoubleArr) GetValue() interface{} {
	return doubleArr.trueValue
}

// GetValueAtIndex returns the value at the specified index
func (doubleArr *DoubleArr) GetValueAtIndex(index int) float64 {
	return doubleArr.trueValue[index]
}

// IsPersistant returns whether or not the entry should persist beyond restarts.
func (doubleArr *DoubleArr) IsPersistant() bool {
	return doubleArr.isPersistant
}

// Clone returns an identical entry
func (doubleArr *DoubleArr) Clone() *DoubleArr {
	return &DoubleArr{
		trueValue:    doubleArr.trueValue,
		isPersistant: doubleArr.isPersistant,
		Base:         *doubleArr.Base.clone(),
	}
}

// CompressToBytes returns a byte array representing the DoubleArr entry
func (doubleArr *DoubleArr) CompressToBytes() []byte {
	return doubleArr.Base.CompressToBytes()
}

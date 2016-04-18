package entryupdate

import (
	"io"

	"github.com/HowardStark/abreuvoir/util"
)

// DoubleArr Entry
type DoubleArr struct {
	Base
	trueValue []float64
}

// DoubleArrFromReader builds a DoubleArr entry using the provided parameters
func DoubleArrFromReader(id [2]byte, sequence [2]byte, etype byte, reader io.Reader) (*DoubleArr, error) {
	var tempValSize [1]byte
	_, sizeErr := io.ReadFull(reader, tempValSize[:])
	if sizeErr != nil {
		return nil, sizeErr
	}
	valSize := int(tempValSize[0])
	value := make([]byte, valSize*8)
	_, valErr := io.ReadFull(reader, value[:])
	if valErr != nil {
		return nil, valErr
	}
	return DoubleArrFromItems(id, sequence, etype, value), nil
}

// DoubleArrFromItems builds a DoubleArr entry using the provided parameters
func DoubleArrFromItems(id [2]byte, sequence [2]byte, etype byte, value []byte) *DoubleArr {
	valSize := int(value[0])
	var val []float64
	for counter := 1; (counter-1)/8 < valSize; counter += 8 {
		tempVal := util.BytesToFloat64(value[counter : counter+8])
		val = append(val, tempVal)
	}
	return &DoubleArr{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeDoubleArr,
			Value: value,
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

// Clone returns an identical entry
func (doubleArr *DoubleArr) Clone() *DoubleArr {
	return &DoubleArr{
		trueValue: doubleArr.trueValue,
		Base:      doubleArr.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the DoubleArr entry
func (doubleArr *DoubleArr) CompressToBytes() []byte {
	return doubleArr.Base.compressToBytes()
}

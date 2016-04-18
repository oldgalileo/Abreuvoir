package entryupdate

import (
	"io"

	"github.com/HowardStark/abreuvoir/util"
)

// Double Entry
type Double struct {
	Base
	trueValue float64
}

// DoubleFromReader builds a double entry using the provided parameters
func DoubleFromReader(id [2]byte, sequence [2]byte, etype byte, reader io.Reader) (*Double, error) {
	var value [8]byte
	_, err := io.ReadFull(reader, value[:])
	if err != nil {
		return nil, err
	}
	return DoubleFromItems(id, sequence, etype, value[:]), nil
}

// DoubleFromItems builds a double entry using the provided parameters
func DoubleFromItems(id [2]byte, sequence [2]byte, etype byte, value []byte) *Double {
	val := util.BytesToFloat64(value[:8])
	return &Double{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeDouble,
			Value: value,
		},
	}
}

// GetValue returns the value of the Double
func (double *Double) GetValue() interface{} {
	return double.trueValue
}

// Clone returns an identical entry
func (double *Double) Clone() *Double {
	return &Double{
		trueValue: double.trueValue,
		Base:      double.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the Double entry
func (double *Double) CompressToBytes() []byte {
	return double.Base.compressToBytes()
}

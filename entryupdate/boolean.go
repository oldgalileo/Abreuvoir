package entryupdate

import "io"

// Boolean Entry
type Boolean struct {
	Base
	trueValue bool
}

// BooleanFromReader builds a boolean entry using the provided parameters
func BooleanFromReader(id [2]byte, sequence [2]byte, etype byte, reader io.Reader) (*Boolean, error) {
	var value [1]byte
	_, err := io.ReadFull(reader, value[:])
	if err != nil {
		return nil, err
	}
	return BooleanFromItems(id, sequence, etype, value[:]), nil
}

// BooleanFromItems builds a boolean entry using the provided parameters
func BooleanFromItems(id [2]byte, sequence [2]byte, etype byte, value []byte) *Boolean {
	val := (value[0] == boolTrue)
	return &Boolean{
		trueValue: val,
		Base: Base{
			ID:    id,
			Seq:   sequence,
			Type:  typeBoolean,
			Value: value,
		},
	}
}

// GetValue returns the value of the Boolean
func (boolean *Boolean) GetValue() interface{} {
	return boolean.trueValue
}

// Clone returns an identical entry
func (boolean *Boolean) Clone() *Boolean {
	return &Boolean{
		trueValue: boolean.trueValue,
		Base:      boolean.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the Boolean entry
func (boolean *Boolean) CompressToBytes() []byte {
	return boolean.Base.compressToBytes()
}

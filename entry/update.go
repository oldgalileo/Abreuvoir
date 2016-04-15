package entry

import "errors"

// Update entry is a partial entry containing only certain fields of an actual entry
type Update struct {
	ID    [2]byte
	Seq   [2]byte
	Type  byte
	Value []byte
}

// UpdateFromBytes builds an Update from a byte slice
func UpdateFromBytes(data []byte) (*Update, error) {
	dID := [2]byte{data[0], data[1]}
	dSeq := [2]byte{data[2], data[3]}
	dType := data[4]
	dValue := data[5:]
	if dType == typeBoolean || dType == typeDouble || dType == typeString || dType == typeRaw || dType == typeBooleanArr || dType == typeDoubleArr || dType == typeStringArr {
		return &Update{
			ID:    dID,
			Seq:   dSeq,
			Type:  dType,
			Value: dValue,
		}, nil
	}
	return nil, errors.New("entry.Update: unknown entry type")
}

// UpdateFromItems builds an Update using the provided parameters
func UpdateFromItems(dID [2]byte, dSeq [2]byte, dType byte, dValue []byte) (*Update, error) {
	if dType == typeBoolean || dType == typeDouble || dType == typeString || dType == typeRaw || dType == typeBooleanArr || dType == typeDoubleArr || dType == typeStringArr {
		return &Update{
			ID:    dID,
			Seq:   dSeq,
			Type:  dType,
			Value: dValue,
		}, nil
	}
	return nil, errors.New("entry.Update: unknown entry type")
}

// CompressToBytes returns a byte slice representing the Update entry
func (update *Update) CompressToBytes() []byte {
	compressed := []byte{}
	compressed = append(compressed, update.ID[:]...)
	compressed = append(compressed, update.Seq[:]...)
	compressed = append(compressed, update.Type)
	compressed = append(compressed, update.Value...)
	return compressed
}

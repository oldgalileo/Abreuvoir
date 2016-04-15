package entry

// FlagUpdate entry is a partial entry containing only certain fields of an actual entry
type FlagUpdate struct {
	ID           [2]byte
	IsPersistant bool
	flags        byte
}

// FlagUpdateFromBytes builds an FlagUpdate from a byte slice
func FlagUpdateFromBytes(data []byte) *FlagUpdate {
	dID := [2]byte{data[0], data[1]}
	dFlags := data[2]
	dPersist := (dFlags == flagPersist)
	return &FlagUpdate{
		ID:           dID,
		IsPersistant: dPersist,
		flags:        dFlags,
	}
}

// FlagUpdateFromItems builds an FlagUpdate using the provided parameters
func FlagUpdateFromItems(dID [2]byte, dFlags byte) *FlagUpdate {
	dPersist := (dFlags == flagPersist)
	return &FlagUpdate{
		ID:           dID,
		IsPersistant: dPersist,
		flags:        dFlags,
	}
}

// CompressToBytes returns a byte slice representing the FlagUpdate entry
func (flagUpdate *FlagUpdate) CompressToBytes() []byte {
	compressed := []byte{}
	compressed = append(compressed, flagUpdate.ID[:]...)
	compressed = append(compressed, flagUpdate.flags)
	return compressed
}

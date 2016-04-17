package message

import "github.com/HowardStark/abreuvoir/entry"

// EntryFlagUpdate message
type EntryFlagUpdate struct {
	Base
	update entry.FlagUpdate
}

// EntryFlagUpdateFromFlagUpdate builds an EntryFlagUpdate message from a FlagUpdate
func EntryFlagUpdateFromFlagUpdate(flagUpdateData entry.FlagUpdate) *EntryFlagUpdate {
	return &EntryFlagUpdate{
		update: flagUpdateData,
		Base: Base{
			mType: typeEntryFlagUpdate,
			mData: flagUpdateData.CompressToBytes(),
		},
	}
}

// EntryFlagUpdateFromItems builds an EntryFlagUpdate message using the provided parameters
func EntryFlagUpdateFromItems(entryID [2]byte, entryFlag byte) *EntryFlagUpdate {
	tempFlagUpdate := *entry.FlagUpdateFromItems(entryID, entryFlag)
	return &EntryFlagUpdate{
		update: tempFlagUpdate,
		Base: Base{
			mType: typeEntryFlagUpdate,
			mData: tempFlagUpdate.CompressToBytes(),
		},
	}
}

// GetFlagUpdate returns the Update associated with this EntryFlagUpdate
func (entryUpdate *EntryFlagUpdate) GetFlagUpdate() entry.FlagUpdate {
	return entryUpdate.update
}

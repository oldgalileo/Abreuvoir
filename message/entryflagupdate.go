package message

import "github.com/HowardStark/Abreuvoir/entry"

// EntryFlagUpdate message
type EntryFlagUpdate struct {
	Base
	flagUpdate entry.FlagUpdate
}

// EntryFlagUpdateFromBytes builds an EntryFlagUpdate message using the provided parameters
func EntryFlagUpdateFromItems(entryId [2]byte, entryFlag byte) *EntryFlagUpdate {
	tempFlagUpdate := entry.FlagUpdateFromItems(entryId, entryFlag)
    return &EntryFlagUpdate{
        update: tempFlagUpdate,
        Base: Base{
            mType: typeEntryFlagUpdate,
            mData: tempFlagUpdate.CompressToBytes()
        }
    }
}

// GetFlagUpdate returns the Update associated with this EntryFlagUpdate
func (entryUpdate *EntryFlagUpdate) GetFlagUpdate() *entry.FlagUpdate {
    return entryUpdate.flagUpdate
}

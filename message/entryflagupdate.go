package message

import "github.com/HowardStark/Abreuvoir/entry"

// EntryFlagUpdate message
type EntryFlagUpdate struct {
	Base
	flagUpdate entry.FlagUpdate
}

// EntryFlagUpdateFromBytes builds an EntryUpdateFlag message from an entry
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

// GetUpdate returns the Update associated with this EntryUpdate
func (entryUpdate *EntryFlagUpdate) GetUpdate() *entry.FlagUpdate {
    return entryUpdate.flagUpdate
}

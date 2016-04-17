package message

import "github.com/HowardStark/abreuvoir/entry"

// EntryUpdate message
type EntryUpdate struct {
	Base
	update entry.Update
}

// EntryUpdateFromUpdate builds an EntryUpdate message from an update
func EntryUpdateFromUpdate(entryUpdate entry.Update) *EntryUpdate {
	return &EntryUpdate{
		update: entryUpdate,
		Base: Base{
			mType: typeEntryUpdate,
			mData: entryUpdate.CompressToBytes(),
		},
	}
}

// EntryUpdateFromItems builds an EntryUpdate message from the parameters provided
func EntryUpdateFromItems(entryID [2]byte, entrySeq [2]byte, entryType byte, entryData []byte) (*EntryUpdate, error) {
	tempUpdate, err := entry.UpdateFromItems(entryID, entrySeq, entryType, entryData)
	if err != nil {
		return nil, err
	}
	return &EntryUpdate{
		update: *tempUpdate,
		Base: Base{
			mType: typeEntryUpdate,
			mData: tempUpdate.CompressToBytes(),
		},
	}, nil
}

// GetUpdate returns the Update associated with this EntryUpdate
func (entryUpdate *EntryUpdate) GetUpdate() entry.Update {
	return entryUpdate.update
}

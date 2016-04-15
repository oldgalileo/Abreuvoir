package message

import "github.com/HowardStark/Abreuvoir/entry"

// EntryUpdate message
type EntryUpdate struct {
	Base
	update entry.Update
}

func EntryUpdateFromUpdate(entryUpdate entry.Update) *EntryUpdate {
	return &EntryUpdate{
		update: entryUpdate,
		Base: Base{
			mType: typeEntryUpdate,
			mData: entryUpdate.CompressToBytes(),
		},
	}
}

// EntryUpdateFromBytes builds an EntryUpdate message from an entry
func EntryUpdateFromItems(entryId [2]byte, entrySeq [2]byte, entryType byte, entryData []byte) (*EntryUpdate, error) {
	tempUpdate, err := entry.UpdateFromItems(entryId, entrySeq, entryType, entryData)
	if err != nil {
		return nil, err
	}
	return &EntryUpdate{
		update: tempUpdate,
		Base: Base{
			mType: typeEntryUpdate,
			mData: tempUpdate.CompressToBytes(),
		},
	}, nil
}

// GetUpdate returns the Update associated with this EntryUpdate
func (entryUpdate *EntryUpdate) GetUpdate() *entry.Update {
	return entryUpdate.update
}

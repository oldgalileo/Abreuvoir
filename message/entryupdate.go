package message

import (
	"io"

	"github.com/HowardStark/abreuvoir/entryupdate"
)

// EntryUpdate message
type EntryUpdate struct {
	Base
	Update entryupdate.Adapter
}

// EntryUpdateFromReader meme
func EntryUpdateFromReader(reader io.Reader) (*EntryUpdate, error) {
	tempUpdate, err := entryupdate.BuildFromReader(reader)
	if err != nil {
		return nil, err
	}
	return &EntryUpdate{
		Update: tempUpdate,
		Base: Base{
			mType: typeEntryUpdate,
			mData: tempUpdate.CompressToBytes(),
		},
	}, nil
}

// EntryUpdateFromUpdate builds an EntryUpdate message from an update
func EntryUpdateFromUpdate(entryUpdate entryupdate.Adapter) *EntryUpdate {
	return &EntryUpdate{
		Update: entryUpdate,
		Base: Base{
			mType: typeEntryUpdate,
			mData: entryUpdate.CompressToBytes(),
		},
	}
}

// GetUpdate returns the Update associated with this EntryUpdate
func (entryUpdate *EntryUpdate) GetUpdate() entryupdate.Adapter {
	return entryUpdate.Update
}

// CompressToBytes returns the message in its byte array form
func (entryUpdate *EntryUpdate) CompressToBytes() []byte {
	return entryUpdate.Base.compressToBytes()
}

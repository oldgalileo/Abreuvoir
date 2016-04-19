package message

import (
	"io"

	"github.com/HowardStark/abreuvoir/entry"
)

// EntryFlagUpdate message
type EntryFlagUpdate struct {
	Base
	update entry.FlagUpdate
}

// EntryFlagUpdateFromReader blehhh stomach hurts toooooo
func EntryFlagUpdateFromReader(reader io.Reader) (*EntryFlagUpdate, error) {
	tempFlag, err := entry.FlagUpdateFromReader(reader)
	if err != nil {
		return nil, err
	}
	return &EntryFlagUpdate{
		update: *tempFlag,
		Base: Base{
			mType: TypeEntryFlagUpdate,
			mData: tempFlag.CompressToBytes(),
		},
	}, nil
}

// EntryFlagUpdateFromFlagUpdate builds an EntryFlagUpdate message from a FlagUpdate
func EntryFlagUpdateFromFlagUpdate(flagUpdateData entry.FlagUpdate) *EntryFlagUpdate {
	return &EntryFlagUpdate{
		update: flagUpdateData,
		Base: Base{
			mType: TypeEntryFlagUpdate,
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
			mType: TypeEntryFlagUpdate,
			mData: tempFlagUpdate.CompressToBytes(),
		},
	}
}

// GetFlagUpdate returns the Update associated with this EntryFlagUpdate
func (entryFlagUpdate *EntryFlagUpdate) GetFlagUpdate() entry.FlagUpdate {
	return entryFlagUpdate.update
}

// CompressToBytes returns the message in its byte array form
func (entryFlagUpdate *EntryFlagUpdate) CompressToBytes() []byte {
	return entryFlagUpdate.Base.compressToBytes()
}

// GetType returns the message's type
func (entryFlagUpdate *EntryFlagUpdate) GetType() byte {
	return TypeEntryFlagUpdate
}

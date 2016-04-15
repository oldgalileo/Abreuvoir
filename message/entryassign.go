package message

import "github.com/HowardStark/abreuvoir/entry"

// EntryAssign message
type EntryAssign struct {
	Base
	entry entry.Adapter
}

func EntryAssignFromEntry(newEntry entry.Adapter) *EntryAssign {
	return &EntryAssign{
		entry: newEntry,
		Base: Base{
			mType: typeEntryAssign,
			mData: newEntry.CompressToBytes,
		},
	}
}

func EntryAssignFromBytes(data []byte) (*EntryAssign, error) {
	tempEntry, err := entry.BuildFromBytes(data)
	if err != nil {
		return nil, err
	}
	return &EntryAssign{
		entry: tempEntry,
		Base: Base{
			mType: typeEntryAssign,
			mData: data,
		},
	}, nil
}

package message

import "github.com/HowardStark/abreuvoir/entry"

// EntryAssign message
type EntryAssign struct {
	Base
	entry entry.Adapter
}

// EntryAssignFromEntry builds an EntryAssign message from an entry
func EntryAssignFromEntry(newEntry entry.Adapter) *EntryAssign {
	return &EntryAssign{
		entry: newEntry,
		Base: Base{
			mType: typeEntryAssign,
			mData: newEntry.CompressToBytes(),
		},
	}
}

// EntryAssignFromBytes builds an EntryAssign message and its entry from a byte slice
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

// GetEntry returns the entry associated with this EntryAssign message
func (entryAssign *EntryAssign) GetEntry() entry.Adapter {
	return entryAssign.entry
}

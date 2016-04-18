package message

import "io"

// EntryDelete message
type EntryDelete struct {
	Base
	id [2]byte
}

// EntryDeleteFromReader builds a EntryDelete message using the provided parameters
func EntryDeleteFromReader(reader io.Reader) (*EntryDelete, error) {
	var dID [2]byte
	_, idErr := io.ReadFull(reader, dID[:])
	if idErr != nil {
		return nil, idErr
	}
	return &EntryDelete{
		id: dID,
		Base: Base{
			mType: typeEntryDelete,
			mData: dID[:],
		},
	}, nil
}

// EntryDeleteFromItems builds a EntryDelete message using the provided parameters
func EntryDeleteFromItems(dID [2]byte) *EntryDelete {
	return &EntryDelete{
		id: dID,
		Base: Base{
			mType: typeEntryDelete,
			mData: dID[:],
		},
	}
}

// GetID returns this EntryDelete's ID
func (entryDelete *EntryDelete) GetID() [2]byte {
	return entryDelete.id
}

// CompressToBytes returns the message in its byte array form
func (entryDelete *EntryDelete) CompressToBytes() []byte {
	return entryDelete.Base.compressToBytes()
}

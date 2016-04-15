package message

// EntryDelete message
func EntryDelete struct {
    Base
    id [2]byte
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

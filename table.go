package abreuvoir

// Table represents a collection of data at a specified depth in the network table
type Table struct {
	path   string
	client *Client
}

// GetBoolean calls GetBoolean on the client object
func (table *Table) GetBoolean(key string) bool {
	return table.client.GetBoolean(key)
}

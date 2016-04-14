package abreuvoir

// Client is the NetworkTables Client
type Client struct {
	address string
	port    string
	entries map[string]EntryAdapter
}

func newClient(connAddr, connPort string) *Client {
	client = Client{
		address: connAddr,
		port:    connPort,
	}
	return &client
}

// GetBoolean fetches a boolean at the specified key
func (client *Client) GetBoolean(key string) bool {
	key = sanitizeKey(key)
	_ = key
	return true
}

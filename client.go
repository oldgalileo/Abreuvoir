package abreuvoir

import (
	"github.com/HowardStark/abreuvoir/entry"
	"github.com/HowardStark/abreuvoir/util"
)

// Client is the NetworkTables Client
type Client struct {
	address string
	port    string
	entries map[string]entry.Adapter
}

func newClient(connAddr, connPort string) *Client {
	return &Client{
		address: connAddr,
		port:    connPort,
	}
}

// GetBoolean fetches a boolean at the specified key
func (client *Client) GetBoolean(key string) bool {
	key = util.SanitizeKey(key)
	_ = key
	return true
}

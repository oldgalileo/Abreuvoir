package abreuvoir

import (
	"fmt"
)

var (
	address, port string = "0.0.0.0", "1735"
	client        Client
)

// SetAddress sets the address of the remote server
func SetAddress(newAddress string) {
	address = newAddress
	_ = address
}

// SetPort sets the port of the remote server
func SetPort(newPort string) {
	port = newPort
	_ = port
}

// InitClient initializes the client and the connection to the remote server.
func InitClient() {
	fmt.Println(address, port)
}

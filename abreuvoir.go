package abreuvoir

var (
	address, port string = "0.0.0.0", "1735"
	client        Client
)

// SetAddress sets the address of the remote server
func SetAddress(newAddress string) {
	address = newAddress
}

// SetPort sets the port of the remote server
func SetPort(newPort string) {
	port = newPort
}

// InitClient initializes the client and the connection to the remote server.
func InitClient() {
	client = *newClient(address, port)
}

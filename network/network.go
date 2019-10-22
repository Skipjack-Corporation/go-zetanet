package network

import "io"

// Network interface connection to outsize world.
type Network interface {
	io.Closer
	// Lister tells the network to start listening on a given address.
	Listen(interface{}) error
	// InterfaceListenAddress returns list of addresses
	InterfaceListenAddress() ([]interface{}, error)
}

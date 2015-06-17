package Tools

import (
	"net"
	"strings"
)

// Provides all IP addresses for this server.
func ReadAllIPAddresses4ThisHost() (addresses4Host []string) {
	addresses4Host = ipAddresses
	return
}

// Read all IP addreses once.
func initIPAddresses4ThisHost() {

	// Get any IP addresses:
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		// Case: Error! Use localhost.
		ipAddresses = make([]string, 1)
		ipAddresses[0] = `127.0.0.1`
		return
	}

	counter := 0
	ipAddresses = make([]string, len(addresses))

	// Loop over all addresses:
	for _, address := range addresses {
		addressText := address.String()

		// Case: CIDR notation?
		if strings.Contains(addressText, `/`) {
			// Convert the address:
			addressText = addressText[:strings.Index(addressText, `/`)]
		}

		// Parse the address to determine some meta data:
		ip := net.ParseIP(addressText)

		// Filter out any loopback and local entries:
		if !ip.IsLoopback() && !ip.IsUnspecified() && strings.ToLower(ip.String()) != `fe80::1` {
			ipAddresses[counter] = ip.String()
			counter++
		}
	}

	if counter == 0 {
		// Case: No public facing address found.
		// Use localhost instead.
		ipAddresses = make([]string, 1)
		ipAddresses[0] = `127.0.0.1`
	} else {
		ipAddresses = ipAddresses[:counter]
	}

	return
}

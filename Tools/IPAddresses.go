package Tools

import (
	"net"
	"strings"
)

func ReadAllIPAddresses4ThisHost() (addresses4Host []string) {
	addresses4Host = ipAddresses
	return
}

func initIPAddresses4ThisHost() {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		ipAddresses = make([]string, 1)
		ipAddresses[0] = `127.0.0.1`
		return
	}

	counter := 0
	ipAddresses = make([]string, len(addresses))
	for _, address := range addresses {
		addressText := address.String()
		if strings.Contains(addressText, `/`) {
			addressText = addressText[:strings.Index(addressText, `/`)]
		}

		ip := net.ParseIP(addressText)
		if !ip.IsLoopback() && !ip.IsUnspecified() && strings.ToLower(ip.String()) != `fe80::1` {
			ipAddresses[counter] = ip.String()
			counter++
		}
	}

	if counter == 0 {
		ipAddresses = make([]string, 1)
		ipAddresses[0] = `127.0.0.1`
	} else {
		ipAddresses = ipAddresses[:counter]
	}

	return
}

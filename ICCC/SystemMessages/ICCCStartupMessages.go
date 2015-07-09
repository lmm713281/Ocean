package SystemMessages

// Message type for the startup message of Ocean's servers:
type ICCCOceanStartUpMessage struct {
	PublicIPAddressPort string // The public web server's IP address and port
	AdminIPAddressPort  string // The private admin server's IP address and port
}

// Message type for a startup message for external components:
type ICCCComponentStartUpMessage struct {
	IPAddressPort string // The component's ICCC IP address and port
}

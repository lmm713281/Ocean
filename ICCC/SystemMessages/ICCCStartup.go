package SystemMessages

// Message type for the startup message:
type ICCCStartUpMessage struct {
	PublicIPAddressAndPort string // The public web server's IP address and port
	AdminIPAddressAndPort  string // The private admin server's IP address and port
}

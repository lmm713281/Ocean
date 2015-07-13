package SystemMessages

// The message to register a host to ICCC.
type ICCCRegisterHost struct {
	Hostname      string // The hostname for the end-point
	IPAddressPort string // The IP address and port for the end-point
	Kind          byte   // Ocean || Component
}

package SystemMessages

// The message to register an listener to ICCC.
type ICCCRegisterListener struct {
	Channel       string // The channel for the provided command
	Command       string // The provided command
	IsActive      bool   // Is the command active?
	IPAddressPort string // The IP address and port for the end-point
	Kind          byte   // Ocean || Component
}

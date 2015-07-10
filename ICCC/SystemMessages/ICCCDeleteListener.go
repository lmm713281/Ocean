package SystemMessages

// The message to delete an listener from ICCC.
type ICCCDeleteListener struct {
	Channel       string
	Command       string
	IPAddressPort string
}

package SystemMessages

// The message to delete a host from ICCC.
type ICCCDeleteHost struct {
	Hostname      string
	IPAddressPort string
}

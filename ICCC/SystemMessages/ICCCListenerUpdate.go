package SystemMessages

// The message to update an listener from ICCC.
type ICCCListenerUpdate struct {
	Channel       string
	Command       string
	IPAddressPort string
	IsActiveNew   bool
	Kind          byte
}

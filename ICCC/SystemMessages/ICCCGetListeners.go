package SystemMessages

// A message to request the known ICCC listeners.
type ICCCGetListeners struct {
}

// The answer to the listeners request.
type ICCCGetListenersAnswer struct {
	Channels         []string
	Commands         []string
	IPAddressesPorts []string
}

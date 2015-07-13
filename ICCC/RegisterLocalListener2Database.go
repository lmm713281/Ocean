package ICCC

// The internal function to register a local listener to ICCC.
func registerLocalListener2Database(channel, command string) {
	/*
		Cannot use here the ICCC command to register this listener.
		Because, this host is maybe the first one. In that case,
		there would be no server which can execute the ICCC command.
		Therefore, every Ocean server registers the own listeners directly.
	*/
	registerListener2Database(channel, command, correctAddressWithPort, true, KindOCEAN)
}

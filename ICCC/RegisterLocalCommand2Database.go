package ICCC

// The internal function to register a local command to ICCC.
func registerLocalCommand2Database(channel, command string) {
	/*
		Cannot use here the ICCC command to register this command.
		Because, this host is maybe the first one. In that case,
		there would be no server which can execute the ICCC command.
		Therefore, every Ocean server registers the own commans directly.
	*/
	registerCommand2Database(channel, command, correctAddressWithPort, true)
}

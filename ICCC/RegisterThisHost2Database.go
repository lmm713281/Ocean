package ICCC

import (
	"github.com/SommerEngineering/Ocean/Tools"
)

// Function to register this server to the ICCC.
func registerThisHost2Database() {
	/*
		Cannot use here the ICCC command to register this host.
		Because, this host is maybe the first one. In that case,
		there would be no server which can execute the ICCC command.
		Therefore, every Ocean server registers the own host directly.
	*/
	registerHost2Database(Tools.ThisHostname(), correctAddressWithPort)
}

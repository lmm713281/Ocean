package WebServer

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"net/http"
)

var (
	senderName              LM.Sender    = `System::WebServer` // This is the name for logging event from this package
	serverPublic            *http.Server = nil                 // The public web server
	serverAdmin             *http.Server = nil                 // The private web server
	serverPublicAddressPort string       = ""                  // The public server end-point
	serverAdminAddressPort  string       = ""                  // The private server end-point
)

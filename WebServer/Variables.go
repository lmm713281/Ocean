package WebServer

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"net/http"
)

var (
	senderName              LM.Sender    = `System::WebServer`
	serverPublic            *http.Server = nil
	serverAdmin             *http.Server = nil
	serverPublicAddressPort string       = ""
	serverAdminAddressPort  string       = ""
)

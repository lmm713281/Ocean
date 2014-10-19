package System

import (
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
)

func StartAndBlockForever() {
	ipAddressPort := Tools.LocalIPAddressAndPort()

	// Tell the whole cluster, that we are up and ready:
	data := ICCCStartUpMessage{}
	data.IPAddressAndPort = ipAddressPort
	ICCC.WriteMessage2All(ICCC.ChannelSYSTEM, `System::Start`, data)

	// Start and block:
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Web server is now listening.`, `Configuration for hostname and port.`, ipAddressPort)
	http.ListenAndServe(ipAddressPort, nil)
}

package System

import "net/http"
import "github.com/SommerEngineering/Ocean/Tools"
import "github.com/SommerEngineering/Ocean/ICCC"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

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

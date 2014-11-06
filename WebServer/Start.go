package WebServer

import (
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

func Start() {
	if serverPublic != nil {
		// Tell the whole cluster, that we are up and ready:
		data := ICCCStartUpMessage{}
		data.IPAddressAndPort = serverPublicAddressPort
		ICCC.WriteMessage2All(ICCC.ChannelSYSTEM, `System::Start`, data)

		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Web server is now listening.`, `Configuration for hostname and port.`, serverPublicAddressPort)
		go serverPublic.ListenAndServe()
	}

	if serverAdmin != nil {
		go serverAdmin.ListenAndServe()
	}
}

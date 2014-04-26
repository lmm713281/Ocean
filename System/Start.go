package System

import "net/http"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/ICCC"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func StartAndBlockForever() {
	oceanHostnameAndPort := ConfigurationDB.Read(`OceanHostnameAndPort`)

	// Init ICCC:
	ICCC.WriteMessage2All(ICCC.ChannelSYSTEM, `System::Start`, nil)

	// Start and block:
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Web server is now listening.`, `Configuration for hostname and port.`, oceanHostnameAndPort)
	http.ListenAndServe(oceanHostnameAndPort, nil)
}

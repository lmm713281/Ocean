package WebServer

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

func Start() {

	// Tell the whole cluster, that we are up and ready:
	data := SystemMessages.ICCCStartUpMessage{}

	// Start the public web server:
	if serverPublic != nil {
		data.PublicIPAddressAndPort = serverPublicAddressPort
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Public web server is now listening.`, `Configuration for hostname and port.`, serverPublicAddressPort)
		go serverPublic.ListenAndServe()
	}

	// Start the private web server:
	if serverAdmin != nil {
		data.AdminIPAddressAndPort = serverAdminAddressPort
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Admin web server is now listening.`, `Configuration for hostname and port.`, serverAdminAddressPort)
		go serverAdmin.ListenAndServe()
	}

	// Notify the whole cluster, that this server is now up and ready:
	answers := ICCC.WriteMessage2All(ICCC.ChannelSYSTEM, `System::Start`, data, SystemMessages.DefaultAnswer{})
	for n, obj := range answers {
		if obj != nil {
			answer := obj.(SystemMessages.DefaultAnswer)
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, fmt.Sprintf("An answer to the ICCC start up message: Successful=%v, Status=%d, Answer=%d/%d", answer.CommandSuccessful, answer.CommandAnswer, n+1, len(answers)))
		}
	}
}

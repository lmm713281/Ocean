package WebServer

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/System/Version"
	"strings"
)

func Start() {

	// Tell the whole cluster, that we are up and ready:
	data := SystemMessages.ICCCOceanStartUpMessage{}
	data.OceanVersion = Version.GetVersion()

	// Start the public web server:
	if serverPublic != nil {
		data.PublicIPAddressPort = serverPublicAddressPort
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Public web server is now listening.`, `Configuration for hostname and port.`, serverPublicAddressPort)

		// Is TLS configured?
		if publicTLSEnabled := ConfigurationDB.Read(`PublicWebServerUseTLS`); strings.ToLower(publicTLSEnabled) == `true` {
			go serverPublic.ListenAndServeTLS(ConfigurationDB.Read(`PublicWebServerTLSCertificateName`), ConfigurationDB.Read(`PublicWebServerTLSPrivateKey`))
		} else {
			go serverPublic.ListenAndServe()
		}

	}

	// Start the private web server:
	if serverAdmin != nil {
		data.AdminIPAddressPort = serverAdminAddressPort
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Admin web server is now listening.`, `Configuration for hostname and port.`, serverAdminAddressPort)

		// Is TLS configured?
		if adminTLSEnabled := ConfigurationDB.Read(`AdminWebServerUseTLS`); strings.ToLower(adminTLSEnabled) == `true` {
			go serverAdmin.ListenAndServeTLS(ConfigurationDB.Read(`AdminWebServerTLSCertificateName`), ConfigurationDB.Read(`AdminWebServerTLSPrivateKey`))
		} else {
			go serverAdmin.ListenAndServe()
		}
	}

	// Notify the whole cluster, that this server is now up and ready:
	answers := ICCC.WriteMessage2All(ICCC.ChannelSTARTUP, `System::OceanStart`, ICCC.KindALL, data, SystemMessages.ICCCDefaultAnswer{})
	for n, obj := range answers {
		if obj != nil {
			answer := obj.(SystemMessages.ICCCDefaultAnswer)
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, fmt.Sprintf("An answer to the ICCC start up message: Successful=%v, Status=%d, Answer=%d/%d", answer.CommandSuccessful, answer.CommandAnswer, n+1, len(answers)))
		}
	}
}

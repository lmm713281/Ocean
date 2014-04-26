package ICCC

import "fmt"
import "net/http"
import "github.com/SommerEngineering/Ocean/Tools"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func ICCCHandler(response http.ResponseWriter, request *http.Request) {
	if errParse := request.ParseForm(); errParse != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNETWORK, `Was not able to parse the HTTP form data from an ICCC message!`)
		http.NotFound(response, request)
		return
	}

	messageData := map[string][]string(request.PostForm)
	channel := messageData[`channel`][0]
	command := messageData[`command`][0]
	password := messageData[`InternalCommPassword`][0]

	if password != Tools.InternalCommPassword() {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityCritical, LM.ImpactNone, LM.MessageNamePASSWORD, `Received a ICCC message with wrong password!`, request.RemoteAddr)
		http.NotFound(response, request)
		return
	}

	key := fmt.Sprintf(`%s::%s`, channel, command)
	listener := listeners[key]
	if listener == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to find the correct listener for these ICCC message.`, `channel=`+channel, `command`+command, `hostname=`+Tools.ThisHostname())
	} else {
		listener(messageData)
	}
}

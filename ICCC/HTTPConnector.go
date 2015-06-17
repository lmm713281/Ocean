package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
)

// The HTTP handler for ICCC.
func ICCCHandler(response http.ResponseWriter, request *http.Request) {

	// Cannot parse the form?
	if errParse := request.ParseForm(); errParse != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNETWORK, `Was not able to parse the HTTP form data from an ICCC message!`)
		http.NotFound(response, request)
		return
	}

	// Read the data out of the request:
	messageData := map[string][]string(request.PostForm)

	// The data must contain at least three fields (command, channel & communication password)
	if len(messageData) < 3 {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNETWORK, `The ICCC message contains not enough data: At least the channel, command and password is required!`)
		http.NotFound(response, request)
		return
	}

	// Read the meta data:
	channel := messageData[`channel`][0]
	command := messageData[`command`][0]
	password := messageData[`InternalCommPassword`][0]

	// Check the password:
	if password != Tools.InternalCommPassword() {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityCritical, LM.ImpactNone, LM.MessageNamePASSWORD, `Received a ICCC message with wrong password!`, request.RemoteAddr)
		http.NotFound(response, request)
		return
	}

	// Build the key for the mapping of the listener cache:
	key := fmt.Sprintf(`%s::%s`, channel, command)

	// Get the matching listener
	listener := listeners[key]

	if listener == nil {
		// Case: No such listener
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to find the correct listener for these ICCC message.`, `channel=`+channel, `command`+command, `hostname=`+Tools.ThisHostname())
	} else {
		// Case: Everything is fine => deliver the message
		listener(messageData)
	}
}

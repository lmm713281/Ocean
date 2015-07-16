package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"net/url"
)

// The HTTP handler for the local ICCC listeners. Will used in case, that another server
// want to utelise an listener from this server.
func ICCCHandler(response http.ResponseWriter, request *http.Request) {

	// Cannot parse the form?
	if errParse := request.ParseForm(); errParse != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNETWORK, `Was not able to parse the HTTP form data from an ICCC message!`)
		http.NotFound(response, request)
		return
	}

	// Read the data out of the request:
	messageData := map[string][]string(request.PostForm)

	// The data must contain at least three fields (command, channel & checksum)
	if len(messageData) < 3 {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNETWORK, `The ICCC message contains not enough data: At least the channel, command and checksum is required!`)
		http.NotFound(response, request)
		return
	}

	// Read the meta data:
	channel := messageData[`channel`][0]
	command := messageData[`command`][0]
	receivedChecksum := messageData[`checksum`][0]

	// Remove the checksum as preparation for the re-hash:
	delete(messageData, `checksum`)

	// Re-hash the received message:
	receivedMessageHash := signMessage(messageData).Get(`checksum`)

	// Check the checksums:
	if receivedChecksum != receivedMessageHash {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityCritical, LM.ImpactNone, LM.MessageNamePASSWORD, `Received a ICCC message with wrong checksum!`, request.RemoteAddr, fmt.Sprintf("channel=%s", channel), fmt.Sprintf("command=%s", command))
		http.NotFound(response, request)
		return
	}

	// Build the key for the mapping of the listener cache:
	key := fmt.Sprintf(`%s::%s`, channel, command)

	// Get the matching listener
	listener := listeners[key]

	if listener == nil {
		// Case: No such listener
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to find the correct listener for these ICCC message.`, `channel=`+channel, `command=`+command, `hostname=`+Tools.ThisHostname())
		http.NotFound(response, request)
	} else {
		// Case: Everything is fine => deliver the message and read the answer:
		answersData := listener(messageData)
		if answersData != nil {
			// Convert the answer to HTTP form values:
			values := url.Values(answersData)
			answersString := values.Encode()

			// Write the answer to the other peer:
			fmt.Fprintf(response, "%s", answersString)
		}
	}
}

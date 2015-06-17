package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"net/url"
)

// Send a message to a listener.
func sendMessage(listener Scheme.Listener, data map[string][]string) {
	// Convert the data and encode it:
	valuesHTTP := url.Values(data)

	// Add the communication password:
	valuesHTTP.Add(`InternalCommPassword`, Tools.InternalCommPassword())

	// Try to deliver the message:
	if _, err := http.PostForm(`http://`+listener.IPAddressPort+`/ICCC`, valuesHTTP); err != nil {
		// Case: Was not possible to deliver.
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNETWORK, `Was not able to send the ICCC message.`, err.Error())
	}

	return
}

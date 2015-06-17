package System

import (
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The receiver function for the ICCC message, that a server is up and running.
func icccSystemStart(data map[string][]string) {

	// Converts the HTTP form data into an object:
	_, _, obj := ICCC.Data2Message(&SystemMessages.ICCCStartUpMessage{}, data)

	// Cast the object to the right type:
	messageData := obj.(*SystemMessages.ICCCStartUpMessage)

	// Provide a log entry:
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: The server is now up and ready.`, messageData.PublicIPAddressAndPort, messageData.AdminIPAddressAndPort)
}

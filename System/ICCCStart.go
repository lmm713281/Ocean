package System

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The receiver function for the ICCC message, that a server is up and running.
func icccSystemStart(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, fmt.Sprintf("Was not able to execute the ICCC server startup message. %s", err))
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := ICCC.Data2Message(SystemMessages.ICCCStartUpMessage{}, data)

	// Was it possible to convert the data?
	if obj != nil {
		// Cast the object to the right type:
		messageData := obj.(SystemMessages.ICCCStartUpMessage)

		// Provide a log entry:
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: The server is now up and ready.`, messageData.PublicIPAddressAndPort, messageData.AdminIPAddressAndPort)

		// An answer is necessary:
		return ICCC.Message2Data("", "", SystemMessages.AnswerACK)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
		fmt.Println(`[Error] ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

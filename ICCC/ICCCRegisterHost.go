package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The receiver function for the ICCC message, that registers a host.
func ICCCRegisterHost(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, fmt.Sprintf("Was not able to execute the ICCC register host message. %s", err))
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCRegisterHost{}, data)

	// Was it possible to convert the data?
	if obj != nil {
		// Cast the object to the right type:
		messageData := obj.(SystemMessages.ICCCRegisterHost)

		// Provide a log entry:
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Should register another host.`, messageData.Hostname, messageData.IPAddressPort)

		// Execute the command:
		registerHost2Database(messageData.Hostname, messageData.IPAddressPort)

		// An answer is necessary:
		return Message2Data(``, ``, SystemMessages.AnswerACK)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

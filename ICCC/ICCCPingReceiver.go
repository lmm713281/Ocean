package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The receiver function for the ICCC ping message.
func ICCCPingReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC ping message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCPing{}, data)

	// Was it possible to convert the data?
	if obj != nil {
		// An answer is necessary:
		return Message2Data("", "", SystemMessages.AnswerACK)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to convert the ping message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

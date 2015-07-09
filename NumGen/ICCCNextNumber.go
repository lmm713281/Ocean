package NumGen

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The receiver function for the ICCC message, that registers a command.
func ICCCNextNumber(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, fmt.Sprintf("Was not able to execute the ICCC next number command message. %s", err))
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := ICCC.Data2Message(SystemMessages.ICCCNumGenNext{}, data)

	// Was it possible to convert the data?
	if obj != nil {
		// Cast the object to the right type (just as check):
		_ = obj.(SystemMessages.ICCCNumGenNext)

		// Execute the command:
		nextNumber := GetUniqueID()

		// An answer is necessary:
		answer := SystemMessages.ICCCNumGenNextAnswer{}
		answer.Number = nextNumber
		return ICCC.Message2Data(``, ``, answer)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

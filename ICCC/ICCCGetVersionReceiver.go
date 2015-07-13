package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/System/Version"
	"github.com/SommerEngineering/Ocean/Tools"
)

// The receiver function for the ICCC version message.
func ICCCGetVersionReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC get version message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCGetVersion{}, data)

	// Was it possible to convert the data?
	if obj != nil {

		// Prepare the answer:
		answer := SystemMessages.ICCCGetVersionAnswer{}
		answer.Kind = KindOCEAN
		answer.Name = Tools.ThisHostname()
		answer.Version = Version.GetVersion()

		// An answer is necessary:
		return Message2Data("", "", answer)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to convert the ping message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

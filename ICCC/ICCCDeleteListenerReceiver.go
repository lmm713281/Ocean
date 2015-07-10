package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// The receiver function for the ICCC message, that deletes an listener.
func ICCCDeleteListenerReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC delete listener message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCDeleteListener{}, data)

	// Was it possible to convert the data?
	if obj != nil {

		// Convert the object:
		messageData := obj.(SystemMessages.ICCCDeleteListener)

		// The database selection:
		selectionDelete := bson.D{{`Channel`, messageData.Channel}, {`Command`, messageData.Command}, {`IPAddressPort`, messageData.IPAddressPort}}

		// Delete the entry:
		if errDelete := collectionListener.Remove(selectionDelete); errDelete != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC delete listener message.", errDelete.Error(), fmt.Sprintf("channel='%s'", messageData.Channel), fmt.Sprintf("command='%s'", messageData.Command), fmt.Sprintf("ipAddressPort='%s'", messageData.IPAddressPort))
			return Message2Data(``, ``, SystemMessages.AnswerNACK)
		} else {
			//
			// Case: No error
			//

			// Update the cache as soon as possible:
			InitCacheNow()
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, "An ICCC listener was deleted.", fmt.Sprintf("channel='%s'", messageData.Channel), fmt.Sprintf("command='%s'", messageData.Command), fmt.Sprintf("ipAddressPort='%s'", messageData.IPAddressPort))
			return Message2Data(``, ``, SystemMessages.AnswerACK)
		}
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

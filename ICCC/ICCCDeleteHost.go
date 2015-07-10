package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// The receiver function for the ICCC message, that deletes a host.
func ICCCDeleteHostReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC delete host message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCDeleteHost{}, data)

	// Was it possible to convert the data?
	if obj != nil {

		// Convert the object:
		messageData := obj.(SystemMessages.ICCCDeleteHost)

		// The database selection:
		selectionDelete := bson.D{{`Hostname`, messageData.Hostname}, {`IPAddressPort`, messageData.IPAddressPort}}

		// Delete the entry:
		if errDelete := collectionHosts.Remove(selectionDelete); errDelete != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC delete host message.", errDelete.Error(), fmt.Sprintf("hostname='%s'", messageData.Hostname), fmt.Sprintf("ipAddressPort='%s'", messageData.IPAddressPort))
			return Message2Data(``, ``, SystemMessages.AnswerNACK)
		} else {
			//
			// Case: No error
			//

			// Update the cache as soon as possible:
			InitCacheNow()
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, "An ICCC host was deleted.", fmt.Sprintf("hostname='%s'", messageData.Hostname), fmt.Sprintf("ipAddressPort='%s'", messageData.IPAddressPort))
			return Message2Data(``, ``, SystemMessages.AnswerACK)
		}
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

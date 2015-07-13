package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// The receiver function for the ICCC message, that updates an listener.
func ICCCListenerUpdateReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC update listener message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCListenerUpdate{}, data)

	// Was it possible to convert the data?
	if obj != nil {

		// Convert the object:
		messageData := obj.(SystemMessages.ICCCListenerUpdate)

		// The database selection:
		selectionUpdate := bson.D{{`Channel`, messageData.Channel}, {`Command`, messageData.Command}, {`IPAddressPort`, messageData.IPAddressPort}}

		// The object with holds the new state:
		updatedObject := Scheme.Listener{}
		updatedObject.Channel = messageData.Channel
		updatedObject.Command = messageData.Command
		updatedObject.IPAddressPort = messageData.IPAddressPort
		updatedObject.IsActive = messageData.IsActiveNew
		updatedObject.Kind = messageData.Kind

		// Update the entry:
		if errUpdate := collectionListener.Update(selectionUpdate, updatedObject); errUpdate != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC update listener message.", errUpdate.Error(), fmt.Sprintf("channel='%s'", messageData.Channel), fmt.Sprintf("command='%s'", messageData.Command), fmt.Sprintf("ipAddressPort='%s'", messageData.IPAddressPort))
			return Message2Data(``, ``, SystemMessages.AnswerNACK)
		} else {
			//
			// Case: No error
			//

			// Update the cache as soon as possible:
			InitCacheNow()
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, "An ICCC listener was updated.", fmt.Sprintf("channel='%s'", messageData.Channel), fmt.Sprintf("command='%s'", messageData.Command), fmt.Sprintf("ipAddressPort='%s'", messageData.IPAddressPort), fmt.Sprintf("isActive=%v", messageData.IsActiveNew))
			return Message2Data(``, ``, SystemMessages.AnswerACK)
		}
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

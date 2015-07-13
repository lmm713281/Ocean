package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The receiver function for the ICCC message, that yields the listeners.
func ICCCGetListenersReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC get listeners message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCGetListeners{}, data)

	// Was it possible to convert the data?
	if obj != nil {

		// We have to read from the cache:
		cacheListenerDatabaseLock.RLock()

		// How many listeners we currently known?
		countListeners := cacheListenerDatabase.Len()

		// Prepare the answer object:
		answerMessage := SystemMessages.ICCCGetListenersAnswer{}
		answerMessage.Channels = make([]string, countListeners, countListeners)
		answerMessage.Commands = make([]string, countListeners, countListeners)
		answerMessage.IPAddressesPorts = make([]string, countListeners, countListeners)
		answerMessage.Kinds = make([]byte, countListeners, countListeners)

		// Loop over all hosts which are currently available at the cache:
		n := 0
		for entry := cacheListenerDatabase.Front(); entry != nil; entry = entry.Next() {
			listener := entry.Value.(Scheme.Listener)
			answerMessage.Channels[n] = listener.Channel
			answerMessage.Commands[n] = listener.Command
			answerMessage.IPAddressesPorts[n] = listener.IPAddressPort
			answerMessage.Kinds[n] = listener.Kind
			n++
		}

		// Unlock the cache:
		cacheListenerDatabaseLock.RUnlock()

		// Send the answer:
		return Message2Data(``, ``, answerMessage)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

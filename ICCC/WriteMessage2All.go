package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// Function to broadcast a message to all listeners.
func WriteMessage2All(channel, command string, message interface{}) {
	cacheListenerDatabaseLock.RLock()
	defer cacheListenerDatabaseLock.RUnlock()

	// Convert the message to HTTP data:
	data := message2Data(channel, command, message)
	counter := 0

	// Loop over all listeners which are currently available at the cache:
	for entry := cacheListenerDatabase.Front(); entry != nil; entry = entry.Next() {
		listener := entry.Value.(Scheme.Listener)

		// If the channel and the command matches, deliver the message:
		if listener.Channel == channel && listener.Command == command {
			go sendMessage(listener, data)
			counter++
		}
	}

	// Was not able to deliver to any listener?
	if counter == 0 {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not able to deliver this message, because no listener was found!`, `channel=`+channel, `command=`+command)
	}

	return
}

package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
)

// Function to write a message to any listener.
func WriteMessage2Any(channel, command string, message interface{}) {
	cacheListenerDatabaseLock.RLock()
	defer cacheListenerDatabaseLock.RUnlock()

	// Convert the message to HTTP data:
	data := message2Data(channel, command, message)
	maxCount := cacheListenerDatabase.Len()
	entries := make([]Scheme.Listener, 0, maxCount)
	counter := 0

	// Loop over all listeners which are currently present at the cache:
	for entry := cacheListenerDatabase.Front(); entry != nil; entry = entry.Next() {
		listener := entry.Value.(Scheme.Listener)

		// If the channel and the command matches, store the listener:
		if listener.Channel == channel && listener.Command == command {
			entries = entries[:len(entries)+1]
			entries[counter] = listener
		}
	}

	count := len(entries)
	if count > 0 {
		// Case: Find at least one possible listener. Choose a random one and deliver:
		if len(entries) == 1 {
			listener := entries[0]
			go sendMessage(listener, data)
		} else {
			listener := entries[Tools.RandomInteger(count)]
			go sendMessage(listener, data)
		}
	} else {
		// Case: Find no listener at all.
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not able to deliver this message to any listener, because no listener was found!`, `channel=`+channel, `command=`+command)
	}
}

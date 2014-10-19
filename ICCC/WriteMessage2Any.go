package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
)

func WriteMessage2Any(channel, command string, message interface{}) {
	cacheListenerDatabaseLock.RLock()
	defer cacheListenerDatabaseLock.RUnlock()

	data := message2Data(channel, command, message)
	maxCount := cacheListenerDatabase.Len()
	entries := make([]Scheme.Listener, 0, maxCount)
	counter := 0
	for entry := cacheListenerDatabase.Front(); entry != nil; entry = entry.Next() {
		listener := entry.Value.(Scheme.Listener)
		if listener.Channel == channel && listener.Command == command {
			entries = entries[:len(entries)+1]
			entries[counter] = listener
		}
	}

	count := len(entries)
	if count > 0 {
		listener := entries[Tools.RandomInteger(count)]
		go sendMessage(listener, data)
	} else {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not able to deliver this message to any listener, because no listener was found!`, `channel=`+channel, `command=`+command)
	}
}

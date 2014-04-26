package ICCC

import "github.com/SommerEngineering/Ocean/ICCC/Scheme"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func WriteMessage2All(channel, command string, message interface{}) {
	cacheListenerDatabaseLock.RLock()
	defer cacheListenerDatabaseLock.RUnlock()

	data := message2Data(channel, command, message)
	counter := 0
	for entry := cacheListenerDatabase.Front(); entry != nil; entry = entry.Next() {
		listener := entry.Value.(Scheme.Listener)
		if listener.Channel == channel && listener.Command == command {
			go sendMessage(listener, data)
			counter++
		}
	}

	if counter == 0 {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not able to deliver this message, because no listener was found!`, `channel=`+channel, `command=`+command)
	}

	return
}

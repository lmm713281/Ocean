package System

import "github.com/SommerEngineering/Ocean/ICCC"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func icccSystemStart(data map[string][]string) {
	_, _, obj := ICCC.Data2Message(&ICCCStartUpMessage{}, data)
	messageData := obj.(*ICCCStartUpMessage)
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: The server is now up and ready.`, messageData.IPAddressAndPort)
}

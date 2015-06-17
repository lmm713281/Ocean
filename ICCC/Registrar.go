package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// Register a command to ICCC for a specific channel.
func Registrar(channel, command string, callback func(data map[string][]string)) {
	listenersLock.Lock()
	defer listenersLock.Unlock()

	// Write the command to the database:
	register2Database(channel, command)

	// Register the command at the local cache:
	listeners[fmt.Sprintf(`%s::%s`, channel, command)] = callback

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The registrar has registered a new ICCC command.`, `channel=`+channel, `command=`+command)
}

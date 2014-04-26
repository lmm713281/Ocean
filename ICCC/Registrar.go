package ICCC

import "fmt"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func Registrar(channel, command string, callback func(data map[string][]string)) {
	listenersLock.Lock()
	defer listenersLock.Unlock()

	register2Database(channel, command)
	listeners[fmt.Sprintf(`%s::%s`, channel, command)] = callback

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The registrar has registered a new ICCC command.`, `channel=`+channel, `command=`+command)
}

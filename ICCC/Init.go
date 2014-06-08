package ICCC

import "container/list"
import "github.com/SommerEngineering/Ocean/Tools"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Start init of ICCC.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Done init ICCC.`)

	cacheListenerDatabase = list.New()
	listeners = make(map[string]func(data map[string][]string))

	// Using the local IP address and NOT the configuration "OceanHostnameAndPort":
	correctAddressWithPort = Tools.LocalIPAddressAndPort()

	initDB()
	registerHost2Database()
}

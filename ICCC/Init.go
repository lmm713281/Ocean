package ICCC

import (
	"container/list"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
)

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

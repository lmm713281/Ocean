package ICCC

import "strings"
import "container/list"
import "github.com/SommerEngineering/Ocean/Tools"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Start init of ICCC.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Done init ICCC.`)

	cacheListenerDatabase = list.New()
	listeners = make(map[string]func(data map[string][]string))

	allHostsIPAddresses := Tools.ReadAllIPAddresses4ThisHost()
	oceanHostnameAndPort := ConfigurationDB.Read(`OceanHostnameAndPort`)
	port := oceanHostnameAndPort[strings.Index(oceanHostnameAndPort, `:`):]
	correctAddressWithPort = allHostsIPAddresses[0] + port

	initDB()
	registerHost2Database()
	initCacheTimer()
}

package ICCC

import (
	"container/list"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
)

// Init this package.
func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Start init of ICCC.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Done init ICCC.`)

	// Create the list as cache for all global listener (not only listener from this server):
	cacheListenerDatabase = list.New()

	// Create a mapping as cache for all local listener end-points (functions):
	listeners = make(map[string]func(data map[string][]string) map[string][]string)

	// Using the local IP address:
	correctAddressWithPort = Tools.LocalIPAddressAndPort()

	// Init the database:
	initDB()

	// Register this server to the listener (if not present):
	registerThisHost2Database()
}

package ICCC

import (
	"container/list"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"strings"
)

// Init this package.
func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Start init of ICCC.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Done init ICCC.`)

	// Create the list as cache for all global listener and hosts (not only listener from this server):
	cacheListenerDatabase = list.New()
	cacheHostDatabase = list.New()

	// Create a mapping as cache for all local listener end-points (functions):
	listeners = make(map[string]func(data map[string][]string) map[string][]string)

	// Using the local IP address:
	correctAddressWithPort = Tools.LocalIPAddressAndPort()

	// Determine the correct protocol:
	if publicTLSEnabled := ConfigurationDB.Read(`PublicWebServerUseTLS`); strings.ToLower(publicTLSEnabled) == `true` {
		activeProtocol = "https://"
	} else {
		activeProtocol = "http://"
	}

	// Init the database:
	initDB()

	// Register this server to the listener (if not present):
	registerThisHost2Database()
}

package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"gopkg.in/mgo.v2/bson"
)

// Internal function for the timer-logic thread.
func cacheTimerLogic() {

	// Case: This server goes down now.
	if Shutdown.IsDown() {
		return
	}

	// Get the current counts:
	lastCountListener := cacheListenerDatabase.Len()
	lastCountHosts := cacheHostDatabase.Len()

	// Define the queries:
	selectionListeners := bson.D{{`IsActive`, true}}
	selectionHosts := bson.D{}

	// Get the iterators:
	entriesIteratorListeners := collectionListener.Find(selectionListeners).Iter()
	entriesIteratorHosts := collectionHosts.Find(selectionHosts).Iter()

	//
	// Execute the listeners first:
	//

	entryListener := Scheme.Listener{}
	cacheListenerDatabaseLock.Lock()

	// Re-init the cache:
	cacheListenerDatabase.Init()

	// Loop over all entries
	for entriesIteratorListeners.Next(&entryListener) {
		cacheListenerDatabase.PushBack(entryListener)
	}

	cacheListenerDatabaseLock.Unlock()
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameEXECUTE, `The listener cache was refreshed with the values from the database.`, fmt.Sprintf(`last count=%d`, lastCountListener), fmt.Sprintf(`new count=%d`, cacheListenerDatabase.Len()))

	//
	// Execute now the hosts:
	//

	entryHost := Scheme.Host{}
	cacheHostDatabaseLock.Lock()

	// Re-init the cache:
	cacheHostDatabase.Init()

	// Loop over all entries
	for entriesIteratorHosts.Next(&entryHost) {
		cacheHostDatabase.PushBack(entryHost)
	}

	cacheHostDatabaseLock.Unlock()
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameEXECUTE, `The host cache was refreshed with the values from the database.`, fmt.Sprintf(`last count=%d`, lastCountHosts), fmt.Sprintf(`new count=%d`, cacheHostDatabase.Len()))
}

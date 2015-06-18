package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Internal function for the timer logic thread.
func cacheTimerLogic(waiting bool) {

	// Case: This server goes down now.
	if Shutdown.IsDown() {
		return
	}

	// Define the query and get the iterator:
	lastCount := cacheListenerDatabase.Len()
	selection := bson.D{{`IsActive`, true}}
	entriesIterator := collectionListener.Find(selection).Iter()

	entry := Scheme.Listener{}
	cacheListenerDatabaseLock.Lock()

	// Re-init the cache:
	cacheListenerDatabase.Init()

	// Loop over all entries
	for entriesIterator.Next(&entry) {
		cacheListenerDatabase.PushBack(entry)
	}

	cacheListenerDatabaseLock.Unlock()
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameEXECUTE, `The listener cache was refreshed with the values from the database.`, fmt.Sprintf(`last count=%d`, lastCount), fmt.Sprintf(`new count=%d`, cacheListenerDatabase.Len()))

	// In case, that this function runs at a thread, we want to wait:
	if waiting {
		nextDuration := time.Duration(5) * time.Minute
		if cacheListenerDatabase.Len() == 0 {
			nextDuration = time.Duration(10) * time.Second
		}

		time.Sleep(nextDuration)
	}
}

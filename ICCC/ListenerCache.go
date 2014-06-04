package ICCC

import "fmt"
import "time"
import "labix.org/v2/mgo/bson"
import "github.com/SommerEngineering/Ocean/Shutdown"
import "github.com/SommerEngineering/Ocean/ICCC/Scheme"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func initCacheTimer() {
	go func() {
		for {
			cacheTimerLogic(true)
		}
	}()
}

func cacheTimerLogic(waiting bool) {
	if Shutdown.IsDown() {
		return
	}

	lastCount := cacheListenerDatabase.Len()
	selection := bson.D{{`IsActive`, true}}
	entriesIterator := collectionListener.Find(selection).Iter()
	entry := Scheme.Listener{}

	cacheListenerDatabaseLock.Lock()
	cacheListenerDatabase.Init()
	for entriesIterator.Next(&entry) {
		cacheListenerDatabase.PushBack(entry)
	}

	cacheListenerDatabaseLock.Unlock()
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameEXECUTE, `The listener cache was refreshed with the values from the database.`, fmt.Sprintf(`last count=%d`, lastCount), fmt.Sprintf(`new count=%d`, cacheListenerDatabase.Len()))

	if waiting {
		nextDuration := time.Duration(5) * time.Minute
		if cacheListenerDatabase.Len() == 0 {
			nextDuration = time.Duration(10) * time.Second
		}

		time.Sleep(nextDuration)
	}
}

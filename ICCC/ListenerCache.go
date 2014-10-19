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

func InitCacheNow() {
	startCacheTimerLock.Lock()
	defer startCacheTimerLock.Unlock()

	if cacheTimerRunning {
		return
	}

	cacheTimerLogic(false)
}

func StartCacheTimer() {
	initCacheTimer()
}

func initCacheTimer() {
	startCacheTimerLock.Lock()
	defer startCacheTimerLock.Unlock()

	if cacheTimerRunning {
		return
	} else {
		cacheTimerRunning = true
	}

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

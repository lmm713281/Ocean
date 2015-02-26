package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"time"
)

func ReadSenderNames() (senderNames []string) {
	mutexCacheSenderNames.RLock()
	defer mutexCacheSenderNames.RUnlock()
	senderNames = cacheSenderNames
	return
}

func cacheRefreshSenderNames() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The sender names' refresh thread is now running.`)
	go func() {
		for true {
			mutexCacheSenderNames.Lock()
			cacheSenderNames = readSenderNamesFromDB()
			mutexCacheSenderNames.Unlock()

			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelTALKATIVE, LM.MessageNameEXECUTE, `The sender names' cache was refreshed.`)
			time.Sleep(time.Duration(nameCachesRefreshTimeSeconds) * time.Second)

			if Shutdown.IsDown() {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameSHUTDOWN, `The sender name's refresh thread is now shutting down.`)
				return
			}
		}
	}()
}

func readSenderNamesFromDB() (result []string) {

	var nextSenderNames []string
	if err := logDBCollection.Find(bson.D{}).Distinct(`Sender`, &nextSenderNames); err != nil {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameDATABASE, `Was not able to read the sender names from the database.`, err.Error())
		return
	}

	sort.Strings(nextSenderNames)
	result = nextSenderNames
	return
}

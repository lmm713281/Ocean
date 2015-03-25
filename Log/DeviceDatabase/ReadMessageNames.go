package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"time"
)

func ReadMessageNames() (messageNames []string) {
	mutexCacheMessageNames.RLock()
	defer mutexCacheMessageNames.RUnlock()
	messageNames = cacheMessageNames
	return
}

func cacheRefreshMessageNames() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The message names' refresh thread is now running.`)
	go func() {
		for true {

			data := readMessageNamesFromDB()
			mutexCacheMessageNames.Lock()
			cacheMessageNames = data
			mutexCacheMessageNames.Unlock()

			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelTALKATIVE, LM.MessageNameEXECUTE, `The message names' cache was refreshed.`)
			time.Sleep(time.Duration(nameCachesRefreshTimeSeconds) * time.Second)

			if Shutdown.IsDown() {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameSHUTDOWN, `The message name's refresh thread is now shutting down.`)
				return
			}
		}
	}()
}

func readMessageNamesFromDB() (result []string) {

	var nextMessageNames []string
	if err := logDBCollection.Find(bson.D{}).Distinct(`MessageName`, &nextMessageNames); err != nil {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameDATABASE, `Was not able to read the message names from the database.`, err.Error())
		return
	}

	sort.Strings(nextMessageNames)
	result = nextMessageNames
	return
}

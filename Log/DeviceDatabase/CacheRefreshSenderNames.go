package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"time"
)

// Function for the thread which maintain the sender name cache.
func cacheRefreshSenderNames() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The sender names' refresh thread is now running.`)

	// Use an extra thread:
	go func() {
		// Endless lopp:
		for true {

			// Read the sender names from the DB:
			data := readSenderNamesFromDB()
			mutexCacheSenderNames.Lock()

			// Overwrite the cache:
			cacheSenderNames = data
			mutexCacheSenderNames.Unlock()

			// Sleep some time:
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelTALKATIVE, LM.MessageNameEXECUTE, `The sender names' cache was refreshed.`)
			time.Sleep(time.Duration(nameCachesRefreshTimeSeconds) * time.Second)

			// Case: The server is going down now.
			if Shutdown.IsDown() {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameSHUTDOWN, `The sender name's refresh thread is now shutting down.`)
				return
			}
		}
	}()
}

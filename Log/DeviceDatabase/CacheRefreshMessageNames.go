package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"time"
)

// Function for the thread which maintain the message name cache.
func cacheRefreshMessageNames() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The message names' refresh thread is now running.`)

	// Create an own thread:
	go func() {
		// Endless loop:
		for true {
			// Read the message names rom the DB:
			data := readMessageNamesFromDB()

			// Case: The project name was not set now. This happens by the logging system
			// after adding this logging device.
			if len(data) == 0 {
				// Wait for a moment:
				time.Sleep(time.Second * 3)

				// Try it again:
				continue
			}

			mutexCacheMessageNames.Lock()

			// Overwrite the cache:
			cacheMessageNames = data
			mutexCacheMessageNames.Unlock()

			// Sleep some time:
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelTALKATIVE, LM.MessageNameEXECUTE, `The message names' cache was refreshed.`)
			time.Sleep(time.Duration(nameCachesRefreshTimeSeconds) * time.Second)

			// Case: The server goes down now.
			if Shutdown.IsDown() {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameSHUTDOWN, `The message name's refresh thread is now shutting down.`)
				return
			}
		}
	}()
}

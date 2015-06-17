package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Shutdown"
	"time"
)

// The timeout function writes the logging events to the database afer some time.
func startTimeout() {

	// Starts a new thread:
	go func() {
		for {

			// Case: The system goes down now.
			if Shutdown.IsDown() {
				return
			}

			// Wait for the right time:
			time.Sleep(time.Duration(cacheSizeTime2FlushSeconds) * time.Second)

			// Write the events to the database:
			mutexCacheFull.Lock()
			amount := len(cache)
			for counter := 0; counter < amount; counter++ {
				write2Database(<-cache)
			}
			mutexCacheFull.Unlock()
		}
	}()
}

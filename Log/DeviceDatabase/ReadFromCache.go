package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Shutdown"
	"time"
)

// Case: The cache is full
func cacheFull() {
	mutexCacheFull.Lock()
	defer mutexCacheFull.Unlock()

	if len(cache) < cacheSizeNumberOfEvents {
		return
	}

	for counter := 0; counter < cacheSizeNumberOfEvents; counter++ {
		write2Database(<-cache)
	}
}

// Case: Time out
func initTimeout() {

	go func() {
		for {

			if Shutdown.IsDown() {
				return
			}

			time.Sleep(time.Duration(cacheSizeTime2FlushSeconds) * time.Second)
			mutexCacheFull.Lock()
			amount := len(cache)
			for counter := 0; counter < amount; counter++ {
				write2Database(<-cache)
			}
			mutexCacheFull.Unlock()
		}
	}()
}

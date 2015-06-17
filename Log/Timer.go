package Log

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/Device"
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"time"
)

func initTimer() {

	// Ensure that the timer runs only once:
	if timerIsRunning == true {
		LogFull(senderName, Meta.CategorySYSTEM, Meta.LevelWARN, Meta.SeverityHigh, Meta.ImpactNone, Meta.MessageNameSTARTUP, `The logging timer is already running.`)
		return
	}

	timerIsRunning = true
	LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, Meta.MessageNameSTARTUP, `Create the logging timer now.`, fmt.Sprintf(`Timeout=%d seconds`, logDeviceDelayTimeoutSeconds))

	// Start the timer at a own thread:
	go func() {

		// An endless loop:
		for {

			// Wait for the next run time:
			time.Sleep(time.Duration(logDeviceDelayTimeoutSeconds) * time.Second)

			// Get exklusive access to the buffer:
			mutexDeviceDelays.Lock()

			// Read all the data from the device delay buffer:
			dataArray := logEntryListToArray(deviceDelayBuffer)

			// Re-init the buffer:
			deviceDelayBuffer.Init()

			// Release the lock to the buffer:
			mutexDeviceDelays.Unlock()

			// Read-lock to read the devices list:
			mutexDevices.RLock()

			// For each logging device:
			for entry := devices.Front(); entry != nil; entry = entry.Next() {
				dev := entry.Value.(Device.Device)

				// Deliver the current logging events with an extra thread:
				go dev.Log(dataArray)
			}

			// Release the read-lock:
			mutexDevices.RUnlock()
		}
	}()
}

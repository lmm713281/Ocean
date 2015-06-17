package Log

import (
	"github.com/SommerEngineering/Ocean/Log/Device"
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

// Queue a log event before delivery to the devices.
func deviceDelay(newEntry Meta.Entry) {
	defer checkDeviceDelaySize()

	// Insert the new entry at the correct position (time).
	// To ensure that the causality is guaranteed.
	for logEvent := deviceDelayBuffer.Front(); logEvent != nil; logEvent = logEvent.Next() {
		currentEvent := logEvent.Value.(Meta.Entry)
		if newEntry.Time.Before(currentEvent.Time) {
			mutexDeviceDelays.Lock()
			deviceDelayBuffer.InsertBefore(newEntry, logEvent)
			mutexDeviceDelays.Unlock()
			return
		}
	}

	// Default: Insert at the back!
	mutexDeviceDelays.Lock()
	deviceDelayBuffer.PushBack(newEntry)
	mutexDeviceDelays.Unlock()
}

// Check if the size of the buffer is huge enough to deliver entries.
func checkDeviceDelaySize() {

	// Get exklusive access:
	mutexDeviceDelays.Lock()

	// Is the size huge enough?
	if deviceDelayBuffer.Len() >= logDeviceDelayNumberEvents {

		// Read all entries:
		dataArray := logEntryListToArray(deviceDelayBuffer)

		// Re-init the buffer:
		deviceDelayBuffer.Init()

		// Loop over all devices:
		mutexDevices.RLock()
		for entry := devices.Front(); entry != nil; entry = entry.Next() {
			dev := entry.Value.(Device.Device)

			// Deliver the data with a new thread:
			go dev.Log(dataArray)
		}
		mutexDevices.RUnlock()
	}

	// Release the lock:
	mutexDeviceDelays.Unlock()
}

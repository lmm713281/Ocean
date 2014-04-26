package Log

import "github.com/SommerEngineering/Ocean/Log/Meta"
import "github.com/SommerEngineering/Ocean/Log/Device"

func deviceDelay(newEntry Meta.Entry) {
	defer checkDeviceDelaySize()

	// Insert the new entry at the correct position (time)!
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

func checkDeviceDelaySize() {
	mutexDeviceDelays.Lock()
	if deviceDelayBuffer.Len() >= logDeviceDelayNumberEvents {
		dataArray := logEntryListToArray(deviceDelayBuffer)
		deviceDelayBuffer.Init()

		mutexDevices.RLock()
		for entry := devices.Front(); entry != nil; entry = entry.Next() {
			dev := entry.Value.(Device.Device)
			go dev.Log(dataArray)
		}
		mutexDevices.RUnlock()
	}

	mutexDeviceDelays.Unlock()
}

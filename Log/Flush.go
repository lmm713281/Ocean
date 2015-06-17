package Log

import (
	"github.com/SommerEngineering/Ocean/Log/Device"
)

// Function to force all buffers to flush the events.
func Flush() {

	// Close the entry buffer:
	mutexChannel.Lock()
	channelReady = false
	close(entriesBuffer)
	mutexChannel.Unlock()

	// Wait that the scheduler is done:
	<-schedulerExitSignal

	// Get all log entries from the device delay buffer:
	mutexDeviceDelays.Lock()
	dataArray := logEntryListToArray(deviceDelayBuffer)
	deviceDelayBuffer.Init()
	mutexDeviceDelays.Unlock()

	// Deliver all the events to all devices:
	mutexDevices.RLock()
	for entry := devices.Front(); entry != nil; entry = entry.Next() {
		dev := entry.Value.(Device.Device)
		dev.Log(dataArray) // Want to wait to complete, therefore no new thread here
		dev.Flush()
	}
	mutexDevices.RUnlock()
}

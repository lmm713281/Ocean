package Log

import (
	"github.com/SommerEngineering/Ocean/Log/Device"
)

/*
Please do not call this function your self! This function allows Ocean to flush the logging at the shutting down case.
*/
func Flush() {
	mutexChannel.Lock()
	channelReady = false
	close(entriesBuffer)
	mutexChannel.Unlock()

	<-schedulerExitSignal

	mutexDeviceDelays.Lock()
	dataArray := logEntryListToArray(deviceDelayBuffer)
	deviceDelayBuffer.Init()
	mutexDeviceDelays.Unlock()

	mutexDevices.RLock()
	for entry := devices.Front(); entry != nil; entry = entry.Next() {
		dev := entry.Value.(Device.Device)
		dev.Log(dataArray) // Want to wait to complete, therefore no new thread here
		dev.Flush()
	}
	mutexDevices.RUnlock()
}

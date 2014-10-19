package Log

import (
	"github.com/SommerEngineering/Ocean/Log/Device"
	"time"
)

/*
Please do not call this function your self! This function allows Ocean to flush the logging at the shutting down case.
*/
func Flush() {
	mutexChannel.Lock()
	channelReady = false
	close(entriesBuffer)
	mutexChannel.Unlock()

	// This is a bad design, but the scheduler need some time to write the last messages.
	time.Sleep(15 * time.Second)

	mutexDeviceDelays.Lock()
	dataArray := logEntryListToArray(deviceDelayBuffer)
	deviceDelayBuffer.Init()
	mutexDeviceDelays.Unlock()

	mutexDevices.RLock()
	for entry := devices.Front(); entry != nil; entry = entry.Next() {
		dev := entry.Value.(Device.Device)
		dev.Log(dataArray) // Want to wait to complete, therefore no new thread here
		go dev.Flush()
	}
	mutexDevices.RUnlock()

	// This is a bad design, but the devices need (may) some time to write the last messages:
	time.Sleep(15 * time.Second)
}

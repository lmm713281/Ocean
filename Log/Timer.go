package Log

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/Device"
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"time"
)

func initTimer() {

	if timerIsRunning == true {
		LogFull(senderName, Meta.CategorySYSTEM, Meta.LevelWARN, Meta.SeverityHigh, Meta.ImpactNone, Meta.MessageNameSTARTUP, `The logging timer is already running.`)
		return
	}

	timerIsRunning = true
	LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, Meta.MessageNameSTARTUP, `Create the logging timer now.`, fmt.Sprintf(`Timeout=%d seconds`, logDeviceDelayTimeoutSeconds))

	go func() {

		for {
			time.Sleep(time.Duration(logDeviceDelayTimeoutSeconds) * time.Second)

			mutexDeviceDelays.Lock()
			dataArray := logEntryListToArray(deviceDelayBuffer)
			deviceDelayBuffer.Init()
			mutexDeviceDelays.Unlock()

			mutexDevices.RLock()
			for entry := devices.Front(); entry != nil; entry = entry.Next() {
				dev := entry.Value.(Device.Device)
				go dev.Log(dataArray)
			}
			mutexDevices.RUnlock()

		}
	}()
}

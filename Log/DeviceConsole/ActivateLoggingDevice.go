package DeviceConsole

import "github.com/SommerEngineering/Ocean/Log"

func ActivateLoggingDevice() {
	Log.AddLoggingDevice(Console{})
}

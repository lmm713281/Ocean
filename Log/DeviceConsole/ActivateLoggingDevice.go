package DeviceConsole

import (
	"github.com/SommerEngineering/Ocean/Log"
)

// Function with the setup of the logging device.
func ActivateLoggingDevice() {
	Log.AddLoggingDevice(Console{})
}

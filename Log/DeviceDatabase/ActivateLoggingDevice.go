package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
)

func ActivateLoggingDevice() {
	Log.AddLoggingDevice(Database{})
}

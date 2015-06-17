package DeviceConsole

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

// The logging device.
type Console struct {
}

// This function is the interface between the logging system and the console logger.
func (dev Console) Log(entries []Meta.Entry) {
	for _, entry := range entries {
		fmt.Println(entry.Format())
	}
}

func (dev Console) Flush() {
	// This is not necessary for a console logger
}

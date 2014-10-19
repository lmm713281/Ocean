package DeviceConsole

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

type Console struct {
}

func (dev Console) Log(entries []Meta.Entry) {
	for _, entry := range entries {
		fmt.Println(entry.Format())
	}
}

func (dev Console) Flush() {
	// This is not necessary for a console logger
}

package Device

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

// The interface for every logging device.
type Device interface {
	Log(logEntries []Meta.Entry)
	Flush()
	SetProjectName(projectName string)
}

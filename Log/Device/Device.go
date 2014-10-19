package Device

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

type Device interface {
	Log(logEntries []Meta.Entry)
	Flush()
}

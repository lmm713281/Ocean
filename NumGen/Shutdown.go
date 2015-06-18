package NumGen

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The type for the shutdown function.
type ShutdownFunction struct {
}

// The shutdown handler for this package.
func (a ShutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Shutting down the number generator.`)
}

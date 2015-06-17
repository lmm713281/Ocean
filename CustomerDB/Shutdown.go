package CustomerDB

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The shutdown type.
type ShutdownFunction struct {
}

// The shutdown function for this package.
func (a ShutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Close now the customer database connection.`)
}

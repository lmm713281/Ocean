package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The type for the shutdown function.
type ShutdownFunction struct {
}

// If the Ocean server is shutting down, this function is called to close the database.
func (a ShutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Close now the configuration database connection.`)
	db.Logout()
	session.Close()
}

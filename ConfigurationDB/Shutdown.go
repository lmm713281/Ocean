package ConfigurationDB

import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

/*
Do not use this type by your own! It is a Ocean internal type to provide a shutdown function for the configuration database.
*/
type ShutdownFunction struct {
}

/*
If the Ocean server is shutting down, this function is called to close the database.
*/
func (a ShutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Close now the configuration database connection.`)
	db.Logout()
	session.Close()
}

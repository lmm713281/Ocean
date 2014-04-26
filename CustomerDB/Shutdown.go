package CustomerDB

import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

/*
Please do not use this type. It is an internal type of Ocean to provide a shutdown function!
*/
type ShutdownFunction struct {
}

/*
This function is called if the Ocean server is shutting down.
*/
func (a ShutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Close now the customer database connection.`)
	db.Logout()
	session.Close()
}

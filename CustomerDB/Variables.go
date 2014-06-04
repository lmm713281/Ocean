package CustomerDB

import "labix.org/v2/mgo"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

var (
	mainSession      *mgo.Session = nil
	senderName       LM.Sender    = `System::CustomerDB`
	databaseUsername string       = ``
	databasePassword string       = ``
	databaseDB       string       = ``
)

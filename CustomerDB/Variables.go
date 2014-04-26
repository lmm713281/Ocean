package CustomerDB

import "labix.org/v2/mgo"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

var (
	session    *mgo.Session  = nil
	db         *mgo.Database = nil
	gridFS     *mgo.GridFS   = nil
	senderName LM.Sender     = `System::CustomerDB`
)

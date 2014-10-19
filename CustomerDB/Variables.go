package CustomerDB

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
)

var (
	mainSession      *mgo.Session = nil
	senderName       LM.Sender    = `System::CustomerDB`
	databaseUsername string       = ``
	databasePassword string       = ``
	databaseDB       string       = ``
)

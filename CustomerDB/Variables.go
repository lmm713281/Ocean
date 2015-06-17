package CustomerDB

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
)

var (
	mainSession      *mgo.Session = nil                  // The session for the customer database
	senderName       LM.Sender    = `System::CustomerDB` // This is the name for logging event from this package
	databaseUsername string       = ``                   // The user's name
	databasePassword string       = ``                   // The user's password
	databaseDB       string       = ``                   // The database
)

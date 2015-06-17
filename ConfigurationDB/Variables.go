package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Configuration/Meta"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
)

var (
	session    *mgo.Session       = nil                       // The database session
	db         *mgo.Database      = nil                       // The database
	collection *mgo.Collection    = nil                       // The database collection
	config     Meta.Configuration = Meta.Configuration{}      // The configuration file's data
	senderName LM.Sender          = `System::ConfigurationDB` // This is the name for logging event from this package
)

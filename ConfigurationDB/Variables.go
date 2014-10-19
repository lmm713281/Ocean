package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Configuration/Meta"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
)

var (
	session    *mgo.Session       = nil
	db         *mgo.Database      = nil
	collection *mgo.Collection    = nil
	config     Meta.Configuration = Meta.Configuration{}
	senderName LM.Sender          = `System::ConfigurationDB`
)

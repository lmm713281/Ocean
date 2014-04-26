package ConfigurationDB

import "labix.org/v2/mgo"
import "github.com/SommerEngineering/Ocean/Configuration/Meta"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

var (
	session    *mgo.Session       = nil
	db         *mgo.Database      = nil
	collection *mgo.Collection    = nil
	config     Meta.Configuration = Meta.Configuration{}
	senderName LM.Sender          = `System::ConfigurationDB`
)

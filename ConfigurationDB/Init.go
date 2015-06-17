package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Configuration"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
)

// The init function for this package.
func init() {
	config := Configuration.Read()

	// Connect to MongoDB:
	if newSession, errDial := mgo.Dial(config.ConfigDBHostname); errDial != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to connect to the MongoDB host `+config.ConfigDBHostname, errDial.Error())
		return
	} else {
		session = newSession
	}

	// Use the correct database:
	db = session.DB(config.ConfigDBDatabase)

	// Login:
	if errLogin := db.Login(config.ConfigDBConfigurationCollectionUsername, config.ConfigDBConfigurationCollectionPassword); errLogin != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to login the user `+config.ConfigDBConfigurationCollectionUsername, errLogin.Error())
		return
	}

	// In case of write operations, wait for the majority of servers to be done:
	session.SetSafe(&mgo.Safe{WMode: "majority"})

	// Set the consistency mode to read from any secondary server and write to the primary.
	session.SetMode(mgo.Eventual, true)

	// Get the collection:
	collection = db.C(config.ConfigDBConfigurationCollection)

	// Take care about the index:
	collection.EnsureIndexKey(`Name`)

	// Check the system configuration:
	checkConfiguration()

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `The configuration database is now ready.`)
}

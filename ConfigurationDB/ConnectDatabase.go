package ConfigurationDB

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Configuration/Meta"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"os"
)

// Try to connect to the database.
func connectDatabase(config Meta.Configuration) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[Error] Was not able to connect to the configuration database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", err)
			os.Exit(1)
		}
	}()

	// Connect to MongoDB:
	if newSession, errDial := mgo.Dial(config.ConfigDBHostname); errDial != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to connect to the MongoDB host `+config.ConfigDBHostname, errDial.Error())
		fmt.Printf("[Error] Was not able to connect to the configuration database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", errDial.Error())
		os.Exit(1)
	} else {
		session = newSession
	}

	// Use the correct database:
	db = session.DB(config.ConfigDBDatabase)

	// Login:
	if errLogin := db.Login(config.ConfigDBConfigurationCollectionUsername, config.ConfigDBConfigurationCollectionPassword); errLogin != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to login the user `+config.ConfigDBConfigurationCollectionUsername, errLogin.Error())
		fmt.Printf("[Error] Was not able to connect to the configuration database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", errLogin.Error())
		os.Exit(1)
	}

	// In case of write operations, wait for the majority of servers to be done:
	session.SetSafe(&mgo.Safe{WMode: "majority"})

	// Set the consistency mode to read from any secondary server and write to the primary.
	session.SetMode(mgo.Eventual, true)

	// Get the collection:
	collection = db.C(config.ConfigDBConfigurationCollection)

	// Take care about the index:
	collection.EnsureIndexKey(`Name`)
}

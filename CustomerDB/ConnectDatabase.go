package CustomerDB

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"os"
)

// Try to connect to the database.
func connectDatabase(host, username, password, database string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[Error] Was not able to connect to the customer database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", err)
			os.Exit(0)
		}
	}()

	// Connect to MongoDB:
	if newSession, errDial := mgo.Dial(host); errDial != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to connect to the MongoDB host `+host, errDial.Error())
		fmt.Printf("[Error] Was not able to connect to the customer database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", errDial.Error())
		os.Exit(0)
	} else {
		mainSession = newSession
	}

	// Use the correct database:
	db := mainSession.DB(database)
	if db == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		fmt.Printf("[Error] Was not able to connect to the customer database. Please read https://github.com/SommerEngineering/Ocean.\n")
		os.Exit(0)
	}

	// Login:
	if errLogin := db.Login(username, password); errLogin != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to login the user `+databaseUsername, errLogin.Error())
		fmt.Printf("[Error] Was not able to connect to the customer database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", errLogin.Error())
		os.Exit(0)
	}

	// In case of write operations, wait for the majority of servers to be done:
	mainSession.SetSafe(&mgo.Safe{WMode: "majority"})

	// Set the consistency mode to read from any secondary server and write to the primary.
	// Copied sessions can overwrite this setting of necessary.
	mainSession.SetMode(mgo.Eventual, true)

	// Get the GridFS:
	gridFS := db.GridFS(`fs`)
	if gridFS == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the GridFS from the database.`)
		fmt.Printf("[Error] Was not able to connect to the customer database. Please read https://github.com/SommerEngineering/Ocean.\n")
		os.Exit(0)
	}

	// Ensure the indexes for the GridFS:
	filesCollection := gridFS.Files
	filesCollection.EnsureIndexKey(`uploadDate`)
	filesCollection.EnsureIndexKey(`filename`)
	filesCollection.EnsureIndexKey(`filename`, `uploadDate`)
}

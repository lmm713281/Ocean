package CustomerDB

import "labix.org/v2/mgo"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Init the customer database.`)

	databaseHost := ConfigurationDB.Read(`CustomerDBHost`)
	databaseDB := ConfigurationDB.Read(`CustomerDBDatabase`)
	databaseUsername := ConfigurationDB.Read(`CustomerDBUsername`)
	databasePassword := ConfigurationDB.Read(`CustomerDBPassword`)

	// Connect to MongoDB:
	if newSession, errDial := mgo.Dial(databaseHost); errDial != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to connect to the MongoDB host `+databaseHost, errDial.Error())
		return
	} else {
		session = newSession
	}

	// Use the correct database:
	db = session.DB(databaseDB)
	if db == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	// Login:
	if errLogin := db.Login(databaseUsername, databasePassword); errLogin != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to login the user `+databaseUsername, errLogin.Error())
		return
	}

	// Get the GridFS:
	gridFS = db.GridFS(`fs`)
	if gridFS == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the GridFS from the database.`)
		return
	}

	// Ensure the indexes for the GridFS:
	filesCollection := gridFS.Files
	filesCollection.EnsureIndexKey(`uploadDate`)
	filesCollection.EnsureIndexKey(`filename`)
	filesCollection.EnsureIndexKey(`filename`, `uploadDate`)

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Customer database is now ready.`)
}

package CustomerDB

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The init function for this package.
func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Init the customer database.`)

	// Read the configuration values:
	databaseHost := ConfigurationDB.Read(`CustomerDBHost`)
	databaseDB = ConfigurationDB.Read(`CustomerDBDatabase`)
	databaseUsername = ConfigurationDB.Read(`CustomerDBUsername`)
	databasePassword = ConfigurationDB.Read(`CustomerDBPassword`)

	// Try to connect to the database:
	connectDatabase(databaseHost, databaseUsername, databasePassword, databaseDB)
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Customer database is now ready.`)
}

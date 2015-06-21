package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Configuration"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The init function for this package.
func init() {
	config := Configuration.Read()

	// Connect to the database:
	connectDatabase(config)

	// Check the system configuration:
	checkConfiguration()

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `The configuration database is now ready.`)
}

package DeviceDatabase

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"strconv"
)

// The init function for this package.
func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting now the database logging.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the database logging done.`)

	// Init the database first:
	initDatabase()

	//
	// Read all configuration values:
	//

	if value, err := strconv.Atoi(ConfigurationDB.Read(`LogDBCacheSizeNumberOfEvents`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityHigh, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not possible to read the LogDBCacheSizeNumberOfEvents configuration.`, `The default value will be used.`, fmt.Sprintf(`Default value is %d.`, cacheSizeNumberOfEvents))
	} else {
		cacheSizeNumberOfEvents = value
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration LogDBCacheSizeNumberOfEvents was loaded.`, fmt.Sprintf(`The value is %d.`, cacheSizeNumberOfEvents))
	}

	if value, err := strconv.Atoi(ConfigurationDB.Read(`LogDBCacheSizeTime2FlushSeconds`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityHigh, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not possible to read the LogDBCacheSizeTime2FlushSeconds configuration.`, `The default value will be used.`, fmt.Sprintf(`Default value is %d.`, cacheSizeTime2FlushSeconds))
	} else {
		cacheSizeTime2FlushSeconds = value
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration LogDBCacheSizeTime2FlushSeconds was loaded.`, fmt.Sprintf(`The value is %d.`, cacheSizeTime2FlushSeconds))
	}

	if value, err := strconv.Atoi(ConfigurationDB.Read(`LogDBWebInterfaceNameCacheRefreshTimeSeconds`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityHigh, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not possible to read the LogDBWebInterfaceNameCacheRefreshTimeSeconds configuration.`, `The default value will be used.`, fmt.Sprintf(`Default value is %d.`, nameCachesRefreshTimeSeconds))
	} else {
		nameCachesRefreshTimeSeconds = value
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration LogDBWebInterfaceNameCacheRefreshTimeSeconds was loaded.`, fmt.Sprintf(`The value is %d.`, nameCachesRefreshTimeSeconds))
	}

	// Create the cache:
	cache = make(chan LogDBEntry, cacheSizeNumberOfEvents)

	// Starts a thread to write events based on time-outs:
	startTimeout()

	// Starts a thread to refresh the sender name cache:
	cacheRefreshSenderNames()

	// Starts a thread to refresh the message name cache:
	cacheRefreshMessageNames()
}

package DeviceDatabase

import "strconv"
import "fmt"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting now the database logging.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the database logging done.`)

	initDatabase()
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

	cache = make(chan LogDBEntry, cacheSizeNumberOfEvents)
	initTimeout()
}

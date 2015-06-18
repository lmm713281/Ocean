package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// Internal function to check the system configuration.
func checkConfiguration() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Check now the configuration database.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Done checking the configuration database.`)

	CheckSingleConfigurationPresentsAndAddIfMissing(`DefaultLanguageCode`, `en-GB`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`AdminWebServerBinding`, `127.0.0.1:60000`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`AdminWebServerEnabled`, `True`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`AdminWebServerReadTimeoutSeconds`, `10`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`AdminWebServerWriteTimeoutSeconds`, `10`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`AdminWebServerMaxHeaderLenBytes`, `1048576`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`PublicWebServerPort`, `60000`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`PublicWebServerReadTimeoutSeconds`, `10`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`PublicWebServerWriteTimeoutSeconds`, `10`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`PublicWebServerMaxHeaderLenBytes`, `1048576`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`InternalCommPassword`, `please replace this with e.g. a random GUID, etc.`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`CustomerDBHost`, `localhost:27017`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`CustomerDBDatabase`, `Ocean`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`CustomerDBUsername`, `root`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`CustomerDBPassword`, `please replace this with a good password`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBHost`, `localhost:27017`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBDatabase`, `Ocean`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBUsername`, `root`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBPassword`, `please replace this with a good password`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBCacheSizeNumberOfEvents`, `50`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBCacheSizeTime2FlushSeconds`, `6`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBEventsExpire`, `True`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBEventsExpireAfterDays`, `365`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDBWebInterfaceNameCacheRefreshTimeSeconds`, `500`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogBufferSize`, `500`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDeviceDelayNumberEvents`, `600`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDeviceDelayTime2FlushSeconds`, `5`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogTimeoutSeconds`, `4`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogStaticFileRequests`, `false`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogUseDatabaseLogging`, `false`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogUseConsoleLogging`, `true`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`OceanUtilizeCPUs`, `2`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`FilenameWebResources`, `web.zip`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`MapStaticFiles2Root`, `false`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`MapStaticFiles2RootRootFile`, `index.html`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`EnableStaticFiles`, `true`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`robots.txt`, `User-agent: *
Disallow:`)
}

// Use this function to ensure that the database contains at least a default value for the configuration.
func CheckSingleConfigurationPresentsAndAddIfMissing(name, value string) {
	if !checkSingleConfigurationPresents(name) {
		addSingleConfiguration(name, value)
	}
}

// Check if a configuration value is present.
func checkSingleConfigurationPresents(name string) (result bool) {
	selection := bson.D{{"Name", name}}
	count, _ := collection.Find(selection).Count()

	return count > 0
}

// Adds a configuration value.
func addSingleConfiguration(name, value string) {
	entry := ConfigurationDBEntry{}
	entry.Name = name
	entry.Value = value
	collection.Insert(entry)

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Add a missing configuration to the configuration database.`, `Name=`+name, `Value=`+value)
}

package ConfigurationDB

import "labix.org/v2/mgo/bson"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func checkConfiguration() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Check now the configuration database.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Done checking the configuration database.`)

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
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogBufferSize`, `500`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDeviceDelayNumberEvents`, `600`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogDeviceDelayTime2FlushSeconds`, `5`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogTimeoutSeconds`, `4`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogStaticFileRequests`, `false`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogUseDatabaseLogging`, `false`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`LogUseConsoleLogging`, `true`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`NumGenActiveHosts`, `please replace this with the correct hostname of the host which is the master number generator`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`NumGenGetHandler`, `http://localhost:80/next/number`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`NumGenBufferSize`, `12`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`OceanHostnameAndPort`, `:60000`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`OceanServerPort`, `60000`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`OceanUtilizeCPUs`, `2`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`FilenameWebResources`, `web.zip`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`MapStaticFiles2Root`, `false`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`MapStaticFiles2RootRootFile`, `index.html`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`EnableStaticFiles`, `true`)
	CheckSingleConfigurationPresentsAndAddIfMissing(`robots.txt`, `User-agent: *
Disallow:`)
}

/*
Use this function to ensure that the database contains at least a default value for the configuration.
*/
func CheckSingleConfigurationPresentsAndAddIfMissing(name, value string) {
	if !checkSingleConfigurationPresents(name) {
		addSingleConfiguration(name, value)
	}
}

func checkSingleConfigurationPresents(name string) (result bool) {
	selection := bson.D{{"Name", name}}
	count, _ := collection.Find(selection).Count()

	return count > 0
}

func addSingleConfiguration(name, value string) {
	entry := ConfigurationDBEntry{}
	entry.Name = name
	entry.Value = value
	collection.Insert(entry)

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameDATABASE, `Add a missing configuration to the configuration database.`, `Name=`+name, `Value=`+value)
}

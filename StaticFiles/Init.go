package StaticFiles

import "io/ioutil"
import "github.com/SommerEngineering/Ocean/CustomerDB"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting now the static files component.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the static files component done.`)

	if ConfigurationDB.Read(`EnableStaticFiles`) != `true` {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Static files are disabled.`)
		return
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Static files are enabled.`)

	// Read the configuration:
	if ConfigurationDB.Read(`MapStaticFiles2Root`) == `true` {
		startFile4Map2Root = ConfigurationDB.Read(`MapStaticFiles2RootRootFile`)
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The desired root document was set.`, `rootDocument=`+startFile4Map2Root)
	}

	logStaticFileRequests = ConfigurationDB.Read(`LogStaticFileRequests`) == `true`

	// Read the static files' data from GridFS:
	gridFS := CustomerDB.GridFS()
	if gridFile, errGridFile := gridFS.Open(`staticFiles.zip`); errGridFile != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to open the static files out of the GridFS!`, errGridFile.Error())
		return
	} else {
		defer gridFile.Close()
		if data, ioError := ioutil.ReadAll(gridFile); ioError != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to read the static files.`, ioError.Error())
			return
		} else {
			zipData = data
		}
	}

	return
}

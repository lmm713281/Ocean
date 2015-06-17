package StaticFiles

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"io/ioutil"
)

// The init of the static file package.
func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting now the static files component.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the static files component done.`)

	// If the static files are disabled, stop here:
	if ConfigurationDB.Read(`EnableStaticFiles`) != `true` {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Static files are disabled.`)
		return
	}

	// Case: Static files are enabled.
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Static files are enabled.`)

	// Read the configuration:
	if ConfigurationDB.Read(`MapStaticFiles2Root`) == `true` {
		startFile4Map2Root = ConfigurationDB.Read(`MapStaticFiles2RootRootFile`)
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The desired root document was set.`, `rootDocument=`+startFile4Map2Root)
	}

	logStaticFileRequests = ConfigurationDB.Read(`LogStaticFileRequests`) == `true`

	// Read the static files' data from GridFS and keep it in-memory:
	dbSession, gridFS := CustomerDB.GridFS()
	defer dbSession.Close()

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

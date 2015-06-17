package WebContent

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"io/ioutil"
)

/*
Init the web content package. It is used to provided some static data for web
frameworks like e.g. jQuery, d3.js, Bootstrap, etc.
*/
func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Run init for the web content.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init for web content done.`)

	// Ensure that we init this package only once:
	if isInit {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityHigh, LM.ImpactNone, LM.MessageNameINIT, `The web content is already fine.`)
		return
	}

	// Get the filename out of the database:
	filename = ConfigurationDB.Read(`FilenameWebResources`)
	dbSession, gridFS := CustomerDB.GridFS()
	defer dbSession.Close()

	// Open the file out of the grid file system:
	if gridFile, errGridFile := gridFS.Open(filename); errGridFile != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to open the web content out of the GridFS!`, filename, errGridFile.Error())
		return
	} else {
		// Read all the data in the memeory cache:
		defer gridFile.Close()
		if data, ioError := ioutil.ReadAll(gridFile); ioError != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to read the web content file.`, filename, ioError.Error())
			return
		} else {
			zipData = data
		}
	}

	return
}

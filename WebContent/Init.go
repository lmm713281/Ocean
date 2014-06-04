package WebContent

import "io/ioutil"
import "github.com/SommerEngineering/Ocean/CustomerDB"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Run init for the web content.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init for web content done.`)

	if isInit {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityHigh, LM.ImpactNone, LM.MessageNameINIT, `The web content is already fine.`)
		return
	}

	filename = ConfigurationDB.Read(`FilenameWebResources`)
	dbSession, gridFS := CustomerDB.GridFS()
	defer dbSession.Close()

	if gridFile, errGridFile := gridFS.Open(filename); errGridFile != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to open the web content out of the GridFS!`, filename, errGridFile.Error())
		return
	} else {
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

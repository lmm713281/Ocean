package System

import "github.com/SommerEngineering/Ocean/Log/DeviceConsole"
import "github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func initLoggingDevices() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the logging devices.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the logging devices done.`)

	if ConfigurationDB.Read(`LogUseDatabaseLogging`) == `true` {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The database logger is active.`)
		activateDatabaseLogger()
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The database logger is NOT active.`)
	}

	if ConfigurationDB.Read(`LogUseConsoleLogging`) == `true` {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The console logger is active.`)
		activateConsoleLogger()
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The console logger is NOT active.`)
	}
}

func activateDatabaseLogger() {
	DeviceDatabase.ActivateLoggingDevice()
}

func activateConsoleLogger() {
	DeviceConsole.ActivateLoggingDevice()
}

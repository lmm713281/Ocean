package System

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	"github.com/SommerEngineering/Ocean/Log/DeviceConsole"
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// Init the logging devices.
func initLoggingDevices() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the logging devices.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the logging devices done.`)

	// Is the database logger enabled?
	if ConfigurationDB.Read(`LogUseDatabaseLogging`) == `true` {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The database logger is active.`)
		activateDatabaseLogger()
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The database logger is NOT active.`)
	}

	// Is the console logger enabled?
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

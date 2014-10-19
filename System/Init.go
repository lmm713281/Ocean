package System

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/NumGen"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"runtime"
	"strconv"
)

func initSystem() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The system is now starting.`)

	// Set the desired amount of CPUs:
	utilizeCPUs := 2
	if value, err := strconv.Atoi(ConfigurationDB.Read(`OceanUtilizeCPUs`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to read the OceanUtilizeCPUs configuration.`, `Use the default value instead.`)
	} else {
		utilizeCPUs = value
	}

	runtime.GOMAXPROCS(utilizeCPUs)
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration OceanUtilizeCPUs is set.`, fmt.Sprintf(`value=%d`, utilizeCPUs))

	// Apply all desired logging devices:
	initLoggingDevices()

	// Set the logging buffer size:
	logBufferSize := 500
	if value, err := strconv.Atoi(ConfigurationDB.Read(`LogBufferSize`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to read the LogBufferSize configuration.`, `Use the default value instead.`)
	} else {
		logBufferSize = value
	}

	Log.SetBufferSize(logBufferSize)
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration LogBufferSize is set.`, fmt.Sprintf(`value=%d`, logBufferSize))

	// Set the logging device delay (number of events):
	logDeviceDelayNumberEvents := 600
	if value, err := strconv.Atoi(ConfigurationDB.Read(`LogDeviceDelayNumberEvents`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to read the LogDeviceDelayNumberEvents configuration.`, `Use the default value instead.`)
	} else {
		logDeviceDelayNumberEvents = value
	}

	Log.SetDeviceDelayNumberEvents(logDeviceDelayNumberEvents)
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration LogDeviceDelayNumberEvents is set.`, fmt.Sprintf(`value=%d`, logDeviceDelayNumberEvents))

	// Set the logging device delay time to flush (seconds):
	logDeviceDelayTime2FlushSeconds := 5
	if value, err := strconv.Atoi(ConfigurationDB.Read(`LogDeviceDelayTime2FlushSeconds`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to read the LogDeviceDelayTime2FlushSeconds configuration.`, `Use the default value instead.`)
	} else {
		logDeviceDelayTime2FlushSeconds = value
	}

	Log.SetDeviceDelayTimeoutSeconds(logDeviceDelayTime2FlushSeconds)
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration LogDeviceDelayTime2FlushSeconds is set.`, fmt.Sprintf(`value=%d`, logDeviceDelayTime2FlushSeconds))

	// Set the logging timeout (seconds):
	logTimeoutSeconds := 3
	if value, err := strconv.Atoi(ConfigurationDB.Read(`LogTimeoutSeconds`)); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `Was not able to read the LogTimeoutSeconds configuration.`, `Use the default value instead.`)
	} else {
		logTimeoutSeconds = value
	}

	Log.SetTimeoutSeconds(logTimeoutSeconds)
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Configuration LogTimeoutSeconds is set.`, fmt.Sprintf(`value=%d`, logTimeoutSeconds))

	// Apply these changes:
	Log.ApplyConfigurationChanges()
	Log.LoggingIsReady()

	// Register all system shutdown handlers:
	//
	// Please notice: If the shutdown event occurs ...
	//		* all application handlers are called (order: last comed, first served)
	//		* then, these system handlers are called (order: last comed, first served)
	//		* and finally, the logging device / system gets closed
	Shutdown.InitShutdown()
	Shutdown.AddShutdownHandler(ConfigurationDB.ShutdownFunction{})
	Shutdown.AddShutdownHandler(CustomerDB.ShutdownFunction{})
	Shutdown.AddShutdownHandler(ICCC.ShutdownFunction{})
	Shutdown.AddShutdownHandler(NumGen.ShutdownFunction{})

	// The logging subsystem is not registered here, because it will be automated called at the end

	// Register all system ICCC commands:
	ICCC.Registrar(ICCC.ChannelSYSTEM, `System::Start`, icccSystemStart)

	// Start the ICCC Listener Cache:
	ICCC.InitCacheNow() // Blocking, until the job is done
	ICCC.StartCacheTimer()
}

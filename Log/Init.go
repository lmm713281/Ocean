package Log

import (
	"container/list"
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"strconv"
	"sync"
)

// Init the logging package.
func init() {

	// Read the project name:
	readProjectName()

	// Create the mutexe:
	mutexDeviceDelays = sync.Mutex{}
	mutexPreChannelBuffer = sync.Mutex{}
	mutexChannel = sync.RWMutex{}

	// Create buffers:
	preChannelBuffer = list.New()
	deviceDelayBuffer = list.New()

	// Create the device list:
	devices = list.New()

	// Channel to exit the scheduler:
	schedulerExitSignal = make(chan bool)

	initTimer()
	initCode()
}

func initCode() {
	// Creates the buffer for logging entries:
	entriesBuffer = make(chan Meta.Entry, logBufferSize)
	LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, `Starting`, `The logger is now starting.`, `logBufferSize=`+strconv.Itoa(int(logBufferSize)), `logBufferTimeoutSeconds=`+strconv.Itoa(int(logBufferTimeoutSeconds)))

	// Start the scheduler as new thread:
	go scheduler(entriesBuffer)
}

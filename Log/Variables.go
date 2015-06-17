package Log

import (
	"container/list"
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"sync"
)

var (
	entriesBuffer                chan Meta.Entry = nil            // The channel / buffer for new log entries
	schedulerExitSignal          chan bool       = nil            // Exit signal for the scheduler
	logBufferSize                int             = 500            // Buffer size for the logging
	logBufferTimeoutSeconds      int             = 4              // Timeout for the logging
	logDeviceDelayNumberEvents   int             = 600            // Delay of # of events for the devices
	logDeviceDelayTimeoutSeconds int             = 5              // Timeout for the logging devices
	channelReady                 bool            = false          // State of the channel
	preChannelBufferUsed         bool            = false          // State of the logging (pre or ready?)
	preChannelBuffer             *list.List      = nil            // Extra buffer for the pre logging phase
	deviceDelayBuffer            *list.List      = nil            // Buffer for the batch write to the devices
	devices                      *list.List      = nil            // List of all devices
	mutexDeviceDelays            sync.Mutex      = sync.Mutex{}   // Mutex for buffer
	mutexPreChannelBuffer        sync.Mutex      = sync.Mutex{}   // Mutex for buffer
	mutexChannel                 sync.RWMutex    = sync.RWMutex{} // Mutex for the main channel
	mutexDevices                 sync.RWMutex    = sync.RWMutex{} // Mutex for the devices
	timerIsRunning               bool            = false          // Status of timer
	projectName                  string          = `not set`      // The project name for the logging
	senderName                   Meta.Sender     = `System::Log`  // This is the name for logging event from this package
)

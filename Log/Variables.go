package Log

import (
	"container/list"
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"sync"
)

var (
	entriesBuffer                chan Meta.Entry = nil
	logBufferSize                int             = 500
	logBufferTimeoutSeconds      int             = 4
	logDeviceDelayNumberEvents   int             = 600
	logDeviceDelayTimeoutSeconds int             = 5
	channelReady                 bool            = false
	preChannelBufferUsed         bool            = false
	preChannelBuffer             *list.List      = nil
	deviceDelayBuffer            *list.List      = nil
	devices                      *list.List      = nil
	mutexDeviceDelays            sync.Mutex      = sync.Mutex{}
	mutexPreChannelBuffer        sync.Mutex      = sync.Mutex{}
	mutexChannel                 sync.RWMutex    = sync.RWMutex{}
	mutexDevices                 sync.RWMutex    = sync.RWMutex{}
	timerIsRunning               bool            = false
	projectName                  string          = `not set`
	senderName                   Meta.Sender     = `System::Log`
)

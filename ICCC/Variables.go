package ICCC

import (
	"container/list"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"sync"
)

// Some pre-defined channels:
const (
	ChannelSYSTEM   string = `System`           // The common system channel.
	ChannelNUMGEN   string = `System::NumGen`   // A channel for the number generator.
	ChannelSHUTDOWN string = `System::Shutdown` // A channel for system shutdown messages.
	ChannelSTARTUP  string = `System::Startup`  // A channel for system startup messages.
	ChannelICCC     string = `System::ICCC`     // A common ICCC channel.
	ChannelPING     string = `System::Ping`     // A channel for pings.
	ChannelLOGGING  string = `System::Logging`  // A channel for send log events to the logging system
)

var (
	senderName                LM.Sender                                                     = `ICCC`                                                                               // This is the name for logging event from this package
	db                        *mgo.Database                                                 = nil                                                                                  // The database
	dbSession                 *mgo.Session                                                  = nil                                                                                  // The database session
	collectionListener        *mgo.Collection                                               = nil                                                                                  // The database collection for listeners
	collectionHosts           *mgo.Collection                                               = nil                                                                                  // The database collection for hosts
	reservedSystemChannels    []string                                                      = []string{ChannelSYSTEM, ChannelNUMGEN, ChannelSHUTDOWN, ChannelSTARTUP, ChannelICCC} // The reserved and pre-defined system channels
	listeners                 map[string]func(data map[string][]string) map[string][]string = nil                                                                                  // The listener cache for all local available listeners with local functions
	listenersLock             sync.RWMutex                                                  = sync.RWMutex{}                                                                       // The mutex for the listener cache
	cacheListenerDatabase     *list.List                                                    = nil                                                                                  // The globally cache for all listeners from all servers
	cacheListenerDatabaseLock sync.RWMutex                                                  = sync.RWMutex{}                                                                       // The mutex for the globally cache
	cacheHostDatabase         *list.List                                                    = nil                                                                                  // The cache for all hosts entries
	cacheHostDatabaseLock     sync.RWMutex                                                  = sync.RWMutex{}                                                                       // The read-write mutex for the host cache
	startCacheTimerLock       sync.Mutex                                                    = sync.Mutex{}                                                                         // Mutex for the start timer
	cacheTimerRunning         bool                                                          = false                                                                                // Is the timer running?
	correctAddressWithPort    string                                                        = ``                                                                                   // The IP address and port of the this local server
)

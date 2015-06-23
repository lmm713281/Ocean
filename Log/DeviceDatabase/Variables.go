package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"sync"
)

var (
	senderName                   LM.Sender             = `System::Logger::Database` // This is the name for logging event from this package
	mutexCacheFull               sync.Mutex            = sync.Mutex{}               // Mutex for the cache full event
	mutexCacheSenderNames        sync.RWMutex          = sync.RWMutex{}             // Read/write mutex for the sender names
	mutexCacheMessageNames       sync.RWMutex          = sync.RWMutex{}             // Read/write mutex for the messages names
	cache                        chan LogDBEntry       = nil                        // The cache
	cacheSizeNumberOfEvents      int                   = 50                         // How many events are cached?
	cacheSizeTime2FlushSeconds   int                   = 6                          // Wait how many seconds before forcing to write events?
	nameCachesRefreshTimeSeconds int                   = 300                        // Wait how many seconds until we reload the sender and message names?
	cacheSenderNames             []Scheme.Sender       = nil                        // Cache for the sender names
	cacheMessageNames            []Scheme.MessageNames = nil                        // Cache for the message names
	logDB                        *mgo.Database         = nil                        // The logging database
	logDBSession                 *mgo.Session          = nil                        // The logging database session
	logDBCollection              *mgo.Collection       = nil                        // The logging collection
	loggingViewerPageSize        int                   = 26                         // How many records per page for the logging web viewer?
	projectName                  string                = `not set`                  // The project name for the logging
	isProjectNameSet             bool                  = false                      // Status about the project name
)

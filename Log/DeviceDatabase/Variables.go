package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"sync"
)

var (
	// This is the name for logging event from this package:
	senderName                   LM.Sender             = `System::Logger::Database`
	mutexCacheFull               sync.Mutex            = sync.Mutex{}
	mutexCacheSenderNames        sync.RWMutex          = sync.RWMutex{}
	mutexCacheMessageNames       sync.RWMutex          = sync.RWMutex{}
	cache                        chan LogDBEntry       = nil
	cacheSizeNumberOfEvents      int                   = 50
	cacheSizeTime2FlushSeconds   int                   = 6
	nameCachesRefreshTimeSeconds int                   = 300
	cacheSenderNames             []Scheme.Sender       = nil
	cacheMessageNames            []Scheme.MessageNames = nil
	logDB                        *mgo.Database         = nil
	logDBSession                 *mgo.Session          = nil
	logDBCollection              *mgo.Collection       = nil
)

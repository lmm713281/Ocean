package DeviceDatabase

import "sync"
import "labix.org/v2/mgo"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

var (
	senderName                 LM.Sender       = `System::Logger::Database`
	mutexCacheFull             sync.Mutex      = sync.Mutex{}
	cache                      chan LogDBEntry = nil
	cacheSizeNumberOfEvents    int             = 50
	cacheSizeTime2FlushSeconds int             = 6
	logDB                      *mgo.Database   = nil
	logDBSession               *mgo.Session    = nil
	logDBCollection            *mgo.Collection = nil
)

package DeviceDatabase

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"sync"
)

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
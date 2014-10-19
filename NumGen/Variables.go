package NumGen

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"sync"
)

var (
	correctPassword   string                = ``
	senderName        LM.Sender             = `System::NumGen::Producer`
	isActive          bool                  = false
	getHandler        string                = ``
	db                *mgo.Database         = nil
	dbSession         *mgo.Session          = nil
	collectionNumGen  *mgo.Collection       = nil
	channelBufferSize int                   = 10
	channelList       map[string]chan int64 = nil
	channelListLock   sync.RWMutex          = sync.RWMutex{}
)

const (
	badNumber64  int64 = 9222222222222222222
	startValue64 int64 = -9223372036854775808
)

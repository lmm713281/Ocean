package ICCC

import (
	"container/list"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"sync"
)

const (
	ChannelSYSTEM   string = `System`
	ChannelNUMGEN   string = `System::NumGen`
	ChannelSHUTDOWN string = `System::Shutdown`
	ChannelSTARTUP  string = `System::Startup`
	ChannelICCC     string = `System::ICCC`
)

var (
	senderName                LM.Sender                                 = `ICCC`
	db                        *mgo.Database                             = nil
	dbSession                 *mgo.Session                              = nil
	collectionListener        *mgo.Collection                           = nil
	collectionHosts           *mgo.Collection                           = nil
	reservedSystemChannels    []string                                  = []string{ChannelSYSTEM, ChannelNUMGEN, ChannelSHUTDOWN, ChannelSTARTUP, ChannelICCC}
	listeners                 map[string]func(data map[string][]string) = nil
	listenersLock             sync.RWMutex                              = sync.RWMutex{}
	cacheListenerDatabase     *list.List                                = nil
	cacheListenerDatabaseLock sync.RWMutex                              = sync.RWMutex{}
	startCacheTimerLock       sync.Mutex                                = sync.Mutex{}
	cacheTimerRunning         bool                                      = false
	correctAddressWithPort    string                                    = ``
)

package NumGen

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"sync"
	"time"
)

var (
	senderName                   LM.Sender  = `System::NumGen` // This is the name for logging event from this package
	genLock                      sync.Mutex = sync.Mutex{}     // The mutex for the generator
	genCurrentTime               time.Time  = time.Now().UTC() // The time for the last generated number
	genCurrentMillisecond        int        = 0                // The millisecond for the last generated number
	genCurrentMillisecondCounter int        = 0                // The counter of how many numbers are generated at the same time
)

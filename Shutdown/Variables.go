package Shutdown

import (
	"container/list"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"os"
)

var (
	shutdownSignal   chan os.Signal = make(chan os.Signal) // A channel for the shutdown event
	shutdownHandlers *list.List     = nil                  // All shutdown handlers
	senderName       LM.Sender      = `System::Shutdown`   // This is the name for logging event from this package
	stopAllRequests  bool           = false                // Does the system goes down now?
)

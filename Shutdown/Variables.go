package Shutdown

import (
	"container/list"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"os"
)

var (
	shutdownSignal   chan os.Signal = make(chan os.Signal)
	shutdownHandlers *list.List     = nil
	senderName       LM.Sender      = `System::Shutdown`
	stopAllRequests  bool           = false
)

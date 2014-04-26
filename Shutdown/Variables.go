package Shutdown

import "os"
import "container/list"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

var (
	shutdownSignal   chan os.Signal = make(chan os.Signal)
	shutdownHandlers *list.List     = nil
	senderName       LM.Sender      = `System::Shutdown`
	stopAllRequests  bool           = false
)

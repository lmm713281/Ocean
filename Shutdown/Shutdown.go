package Shutdown

import "os"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

type ShutdownHandler interface {
	Shutdown()
}

func AddShutdownHandler(handler ShutdownHandler) {
	shutdownHandlers.PushFront(handler)
}

func executeShutdown() {
	sig := <-shutdownSignal
	stopAllRequests = true

	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSHUTDOWN, `The system was called to shutting down.`, sig.String(), `Call now all shutdown handlers.`)
	for handler := shutdownHandlers.Front(); handler != nil; handler = handler.Next() {
		h := handler.Value.(ShutdownHandler)
		h.Shutdown()
	}

	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSHUTDOWN, `The system is shutting down now.`)
	Log.Flush()

	os.Exit(6)
}

package Shutdown

import (
	"container/list"
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"os"
)

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
		safeCall(handler)
	}

	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSHUTDOWN, `The system is shutting down now.`)
	Log.Flush()
	os.Exit(0)
}

func safeCall(handler *list.Element) {
	defer func() {
		if err := recover(); err != nil {
			errObj := fmt.Errorf("%v", err)
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.SeverityNone, LM.ImpactNone, LM.MessageNameSHUTDOWN, `An error occurs for a shutdown handler.`, errObj.Error())
		}
	}()

	h := handler.Value.(ShutdownHandler)
	h.Shutdown()
}

package Shutdown

import (
	"container/list"
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"os"
)

// A type for shutdown handlers.
type ShutdownHandler interface {
	Shutdown()
}

// Function to add new shutdown handlers.
func AddShutdownHandler(handler ShutdownHandler) {
	shutdownHandlers.PushFront(handler)
}

// The thread which waits for the shutdown event.
func executeShutdown() {

	// Wait for the shutdown event:
	sig := <-shutdownSignal
	stopAllRequests = true
	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSHUTDOWN, `The system was called to shutting down.`, sig.String(), `Call now all shutdown handlers.`)

	// Execute all shutdown handlers:
	for handler := shutdownHandlers.Front(); handler != nil; handler = handler.Next() {
		safeCall(handler)
	}

	// Shutdown the logging system:
	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSHUTDOWN, `The system is shutting down now.`)
	Log.Flush()

	// Stop the whole server:
	os.Exit(0)
}

// This function is a wrapper to call safely a shutdown handler.
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

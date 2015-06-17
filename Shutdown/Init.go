package Shutdown

import (
	"container/list"
	"os"
	"os/signal"
)

// Init the package.
func init() {
	shutdownHandlers = list.New()
}

// The manual init for the shutdown notify handler.
func InitShutdown() {

	// Apply the shutdown handler:
	signal.Notify(shutdownSignal, os.Interrupt, os.Kill)

	// Start a new thread which waits for the shutdown event:
	go executeShutdown()
}

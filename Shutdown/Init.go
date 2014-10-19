package Shutdown

import (
	"container/list"
	"os"
	"os/signal"
)

func init() {
	shutdownHandlers = list.New()
}

func InitShutdown() {

	// Apply the shutdown handler:
	signal.Notify(shutdownSignal, os.Interrupt, os.Kill)
	go executeShutdown()
}

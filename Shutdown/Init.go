package Shutdown

import "container/list"
import "os/signal"
import "os"

func init() {
	shutdownHandlers = list.New()
}

func InitShutdown() {

	// Apply the shutdown handler:
	signal.Notify(shutdownSignal, os.Interrupt, os.Kill)
	go executeShutdown()
}

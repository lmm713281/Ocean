package Shutdown

import "container/list"
import "os/signal"
import "os"

func InitShutdown() {
	shutdownHandlers = list.New()

	// Apply the shutdown handler:
	signal.Notify(shutdownSignal, os.Interrupt, os.Kill)
	go executeShutdown()
}

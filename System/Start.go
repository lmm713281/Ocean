package System

import (
	"github.com/SommerEngineering/Ocean/WebServer"
	"time"
)

// The main function for the application.
func StartAndBlockForever() {

	// Starts the public and private web server with own threads:
	WebServer.Start()

	// Wait forever:
	for {
		time.Sleep(1 * time.Second)
	}
}

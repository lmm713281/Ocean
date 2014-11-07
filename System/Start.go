package System

import (
	"github.com/SommerEngineering/Ocean/WebServer"
	"time"
)

func StartAndBlockForever() {
	WebServer.Start()
	for {
		time.Sleep(1 * time.Second)
	}
}

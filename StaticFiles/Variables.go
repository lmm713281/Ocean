package StaticFiles

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

var (
	senderName            LM.Sender = `System::StaticFiles` // This is the name for logging event from this package
	startFile4Map2Root    string    = `index.html`          // The default filename in case of mapping static files to the root
	logStaticFileRequests bool      = false                 // Logging each access?
	zipData               []byte    = nil                   // The in-memory cache of the data
)

package StaticFiles

import LM "github.com/SommerEngineering/Ocean/Log/Meta"

var (
	senderName            LM.Sender = `System::StaticFiles`
	startFile4Map2Root    string    = `index.html`
	logStaticFileRequests bool      = false
	zipData               []byte    = nil
)

package WebContent

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

var (
	isInit               = false                // Ensure that the init happens only once
	filename             = ""                   // The filename of web content file at the grid file system
	zipData    []byte    = nil                  // The memory cache of the data
	senderName LM.Sender = `System::WebContent` // This is the name for logging event from this package
)

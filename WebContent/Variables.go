package WebContent

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

var (
	isInit               = false
	filename             = ""
	zipData    []byte    = nil
	senderName LM.Sender = `System::WebContent`
)

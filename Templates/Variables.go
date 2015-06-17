package Templates

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"html/template"
)

var (
	templates  *template.Template = nil                 // The in-memory cache for the templates
	senderName LM.Sender          = `System::Templates` // This is the name for logging event from this package
	zipData    []byte             = nil                 // The in-memory cache of the ZIP file
	isInit     bool               = false               // State, if the init event is done
)

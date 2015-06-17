package Web

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"html/template"
)

var (
	templates  *template.Template = nil              // The web logging templates
	senderName LM.Sender          = `System::WebLog` // This is the name for logging event from this package
)

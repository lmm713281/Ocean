package Web

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"html/template"
)

var (
	templates  *template.Template = nil
	senderName LM.Sender          = `System::WebLog`
)

package Admin

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"html/template"
)

var (
	AdminTemplates *template.Template = nil             // The admin templates
	senderName     LM.Sender          = `System::Admin` // This is the name for logging event from this package
)

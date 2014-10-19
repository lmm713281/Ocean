package Templates

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"html/template"
)

var (
	templates  *template.Template = nil
	senderName LM.Sender          = `System::Templates`
	zipData    []byte             = nil
	isInit     bool               = false
)

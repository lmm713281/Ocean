package Templates

import "html/template"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

var (
	templates  *template.Template = nil
	senderName LM.Sender          = `System::Templates`
	zipData    []byte             = nil
	isInit     bool               = false
)

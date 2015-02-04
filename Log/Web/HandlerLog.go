package Web

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Log/Web/Scheme"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
)

func HandlerWebLog(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	data := Scheme.Viewer{}
	data.Events = make([]Scheme.LogEvent, 3)
	data.Events[0].AB = Scheme.A
	data.Events[0].LogLevel = Scheme.LogINFO
	data.Events[0].LogLine = `hello world`
	data.Title = `Web Log Viewer`

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
	if executeError := templates.ExecuteTemplate(response, `WebLog`, data); executeError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the web log viewer template.`, executeError.Error())
	}
}

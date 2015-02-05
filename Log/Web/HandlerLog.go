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

	request.ParseForm()
	countParameters := len(request.Form)

	data := Scheme.Viewer{}
	data.Title = `Web Log Viewer`

	if countParameters < 9 {

		// Initial view => refresh & first page (latest logs)
		data.Events = readLatest()
	} else {

		// Custom view
		data.Events = readCustom(request.FormValue(`timeRange`), request.FormValue(`Level`), request.FormValue(`Category`), request.FormValue(`Impact`), request.FormValue(`Severity`), request.FormValue(`MSGName`), request.FormValue(`Sender`), request.FormValue(`CurrentPage`))
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
	if executeError := templates.ExecuteTemplate(response, `WebLog`, data); executeError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the web log viewer template.`, executeError.Error())
	}
}

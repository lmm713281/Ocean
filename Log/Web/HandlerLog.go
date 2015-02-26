package Web

import (
	"github.com/SommerEngineering/Ocean/Log"
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Log/Web/Scheme"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
	"strings"
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
	data.Sender = DeviceDatabase.ReadSenderNames()
	data.MessageNames = DeviceDatabase.ReadMessageNames()

	if countParameters < 9 {

		// Initial view => refresh & first page (latest logs)
		data.Events = readLatest()
		data.SetLiveView = true

	} else {

		// Custom view
		currentLevel := request.FormValue(`Level`)
		currentTimeRange := request.FormValue(`timeRange`)
		currentCategory := request.FormValue(`Category`)
		currentImpact := request.FormValue(`Impact`)
		currentSeverity := request.FormValue(`Severity`)
		currentMessageName := request.FormValue(`MSGName`)
		currentSender := request.FormValue(`Sender`)
		currentPage := request.FormValue(`CurrentPage`)
		currentLiveView := request.FormValue(`LiveView`)

		data.Events = readCustom(currentTimeRange, currentLevel, currentCategory, currentImpact, currentSeverity, currentMessageName, currentSender, currentPage)

		if strings.ToLower(currentLiveView) == `true` {
			data.SetLiveView = true
		}

		if currentLevel != `` {
			data.CurrentLevel = currentLevel
		} else {
			data.CurrentLevel = `*`
		}

		if currentTimeRange != `` {
			data.CurrentTimeRange = currentTimeRange
		} else {
			data.CurrentTimeRange = `*`
		}

		if currentCategory != `` {
			data.CurrentCategory = currentCategory
		} else {
			data.CurrentCategory = `*`
		}

		if currentImpact != `` {
			data.CurrentImpact = currentImpact
		} else {
			data.CurrentImpact = `*`
		}

		//
		// TODO
		//
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
	if executeError := templates.ExecuteTemplate(response, `WebLog`, data); executeError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the web log viewer template.`, executeError.Error())
	}
}

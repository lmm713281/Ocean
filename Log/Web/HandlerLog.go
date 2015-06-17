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

// Handler for accessing the web logging.
func HandlerWebLog(response http.ResponseWriter, request *http.Request) {

	// Case: The system goes down now.
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	// Execute the HTTP form:
	request.ParseForm()
	countParameters := len(request.Form)

	// Setup the data for the HTML template:
	data := Scheme.Viewer{}
	data.Title = `Web Log Viewer`
	data.Sender = DeviceDatabase.ReadSenderNames()
	data.MessageNames = DeviceDatabase.ReadMessageNames()

	// To less parameters?
	if countParameters < 9 {
		// Initial view => refresh & first page (latest logs)
		data.Events = readLatest()
		data.SetLiveView = true
	} else {
		// Case: Custom view
		currentLevel := request.FormValue(`Level`)
		currentTimeRange := request.FormValue(`timeRange`)
		currentCategory := request.FormValue(`Category`)
		currentImpact := request.FormValue(`Impact`)
		currentSeverity := request.FormValue(`Severity`)
		currentMessageName := request.FormValue(`MSGName`)
		currentSender := request.FormValue(`Sender`)
		currentPage := request.FormValue(`CurrentPage`)
		currentLiveView := request.FormValue(`LiveView`)

		// Store the events for the template:
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

		if currentSeverity != `` {
			data.CurrentSeverity = currentSeverity
		} else {
			data.CurrentSeverity = `*`
		}

		if currentMessageName != `` {
			data.CurrentMessageName = currentMessageName
		} else {
			data.CurrentMessageName = `*`
		}

		if currentSender != `` {
			data.CurrentSender = currentSender
		} else {
			data.CurrentSender = `*`
		}
	}

	// Write the MIME type and execute the template:
	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
	if executeError := templates.ExecuteTemplate(response, `WebLog`, data); executeError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the web log viewer template.`, executeError.Error())
	}
}

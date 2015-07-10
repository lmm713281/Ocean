package Web

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Admin"
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
	"strconv"
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

	// The current page as number:
	currentPageNumber := 1

	// The max. page as number:
	lastPageNumber := 1

	// Setup the data for the HTML template:
	data := Scheme.LoggingViewer{}
	data.Title = `Logging Viewer`
	data.Sender = DeviceDatabase.ReadSenderNames()
	data.MessageNames = DeviceDatabase.ReadMessageNames()

	// Get the current page as number:
	currentPage := request.FormValue(`CurrentPage`)
	if currentPage != `` {
		if number, err := strconv.Atoi(currentPage); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityNone, LM.ImpactNone, LM.MessageNameEXECUTE, `Was not able to parse the page number.`, err.Error())
		} else {
			currentPageNumber = number
		}
	}

	// To less parameters?
	if countParameters < 9 {
		// Initial view => first page (latest logs)
		data.Events, lastPageNumber = readLatest()
		data.SetLiveView = true
		if currentPageNumber+1 > lastPageNumber {
			data.NextPage = fmt.Sprintf("%d", lastPageNumber)
		} else {
			data.NextPage = `2`
		}

		data.PreviousPage = `1`
		data.CurrentLevel = `*`
		data.CurrentTimeRange = `*`
		data.CurrentCategory = `*`
		data.CurrentImpact = `*`
		data.CurrentSeverity = `*`
		data.CurrentMessageName = `*`
		data.CurrentSender = `*`
	} else {
		// Case: Custom view
		currentLevel := request.FormValue(`Level`)
		currentTimeRange := request.FormValue(`TimeRange`)
		currentCategory := request.FormValue(`Category`)
		currentImpact := request.FormValue(`Impact`)
		currentSeverity := request.FormValue(`Severity`)
		currentMessageName := request.FormValue(`MSGName`)
		currentSender := request.FormValue(`Sender`)
		currentLiveView := request.FormValue(`LiveView`)

		// Store the events for the template:
		data.Events, lastPageNumber = readCustom(currentTimeRange, currentLevel, currentCategory, currentImpact, currentSeverity, currentMessageName, currentSender, currentPageNumber)

		if strings.ToLower(currentLiveView) == `true` {
			data.SetLiveView = true
		}

		//
		// Correct the form's values to '*' for the any-case:
		//

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

	// Calculate the current, last, previous and next page:
	if currentPage != `` {
		data.CurrentPage = fmt.Sprintf("%d", currentPageNumber)
		data.LastPage = fmt.Sprintf("%d", lastPageNumber)
		if currentPageNumber+1 > lastPageNumber {
			data.NextPage = fmt.Sprintf("%d", lastPageNumber)
		} else {
			data.NextPage = fmt.Sprintf("%d", currentPageNumber+1)
		}

		if currentPageNumber > 1 {
			data.PreviousPage = fmt.Sprintf("%d", currentPageNumber-1)
		} else {
			data.PreviousPage = `1`
		}
		data.CurrentPage = currentPage
	} else {
		data.CurrentPage = `1`
		data.LastPage = fmt.Sprintf("%d", lastPageNumber)
		if currentPageNumber+1 > lastPageNumber {
			data.NextPage = fmt.Sprintf("%d", lastPageNumber)
		} else {
			data.NextPage = `2`
		}
		data.PreviousPage = `1`
	}

	// Write the MIME type and execute the template:
	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
	if executeError := Admin.AdminTemplates.ExecuteTemplate(response, `WebLog`, data); executeError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the web log viewer template.`, executeError.Error())
	}
}

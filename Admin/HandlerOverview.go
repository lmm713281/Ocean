package Admin

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"github.com/SommerEngineering/Ocean/System/Version"
	"net/http"
)

// Handler for accessing the admin's overview.
func HandlerOverview(response http.ResponseWriter, request *http.Request) {

	// Case: The system goes down now.
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	// Get the data ready:
	data := AdminWebOverview{}
	data.Version = Version.GetVersion()

	// Write the MIME type and execute the template:
	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
	if executeError := AdminTemplates.ExecuteTemplate(response, `Overview`, data); executeError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the admin's overview template.`, executeError.Error())
	}
}

package Admin

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
	"strings"
)

// Handler for accessing the file upload function.
func HandlerConfiguration(response http.ResponseWriter, request *http.Request) {

	// Case: The system goes down now.
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	if strings.ToLower(request.Method) == `get` {
		//
		// Case: Send the website to the client
		//

		// Read all configuration values:
		values := ConfigurationDB.ReadAll()

		// Build the data type for the template:
		data := AdminWebConfiguration{}
		data.Configuration = values

		// Write the MIME type and execute the template:
		MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
		if executeError := AdminTemplates.ExecuteTemplate(response, `Configuration`, data); executeError != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the configuration template.`, executeError.Error())
		}
	} else {
		//
		// Case: Receive the changed configuration
		//

		// Read all configuration values:
		values := ConfigurationDB.ReadAll()

		// Loop over all current known values:
		for _, value := range values {

			// Read the new value from the client side:
			newValue := request.FormValue(value.Name)

			// Store the new value:
			value.Value = newValue

			// Update the database:
			ConfigurationDB.UpdateValue(value.Name, value)
		}

		// Redirect the client to the admin's overview:
		defer http.Redirect(response, request, "/configuration", 302)
	}

}

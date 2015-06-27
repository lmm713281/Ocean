package Admin

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"io"
	"net/http"
	"strings"
)

// Handler for accessing the file upload function.
func HandlerFileUpload(response http.ResponseWriter, request *http.Request) {

	// Case: The system goes down now.
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	if strings.ToLower(request.Method) == `get` {
		//
		// Case: Send the website to the client
		//

		// Write the MIME type and execute the template:
		MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
		if executeError := AdminTemplates.ExecuteTemplate(response, `FileUpload`, nil); executeError != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the file upload template.`, executeError.Error())
		}
	} else {
		//
		// Case: Receive the file to upload
		//

		if file, fileHeader, fileError := request.FormFile(`file`); fileError != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameBROWSER, `Was not able to access the file uploaded.`, fileError.Error())
		} else {
			//
			// Case: Access was possible.
			//

			// Get the GridFS from the database:
			dbSession, gridFS := CustomerDB.GridFS()
			defer dbSession.Close()

			// Try to create the desired file at the grid file system:
			if newFile, errNewFile := gridFS.Create(fileHeader.Filename); errNewFile != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to create the desired file at the grid file system.`, errNewFile.Error(), fmt.Sprintf("filename='%s'", fileHeader.Filename))
			} else {

				// Close the files afterwards:
				defer file.Close()
				defer newFile.Close()

				// Try to copy the file's content to the database:
				if _, errCopy := io.Copy(newFile, file); errCopy != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNETWORK, `Was not able to copy the desired file's content to the grid file system.`, errNewFile.Error(), fmt.Sprintf("filename='%s'", fileHeader.Filename))
				} else {
					// Try to determine the MIME type:
					if mimeType, errMime := MimeTypes.DetectType(fileHeader.Filename); errMime != nil {
						Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityLow, LM.ImpactLow, LM.MessageNamePARSE, `Was not able to parse the desired file's MIME type.`, errMime.Error(), fmt.Sprintf("filename='%s'", fileHeader.Filename))
					} else {
						// Set also the MIME type in the database:
						newFile.SetContentType(mimeType.MimeType)
					}
				}
			}
		}

		// Redirect the client to the admin's overview:
		defer http.Redirect(response, request, "/", 302)
	}

}

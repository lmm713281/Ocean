package StaticFiles

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
	"strings"
)

// Handler to deliver static files.
func HandlerStaticFiles(response http.ResponseWriter, request *http.Request) {

	// Case: The system goes down.
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	// Prepare the path by removing the prefix:
	path := strings.Replace(request.RequestURI, `/staticFiles/`, ``, 1)
	path = strings.Replace(path, `%20`, ` `, -1)
	fileType := ``

	// Determine the MIME type:
	if mimeType, errMime := MimeTypes.DetectType(path); errMime != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityMiddle, LM.ImpactMiddle, LM.MessageNameNOTFOUND, `Was not able to detect the MIME type of the font.`, errMime.Error(), path)
		http.NotFound(response, request)
		return
	} else {
		fileType = mimeType.MimeType
	}

	// Case: Was not able to determine the file's type.
	if fileType == `` {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNOTFOUND, `The mime type is unknown.`, path)
		http.NotFound(response, request)
		return
	}

	// Read the file's content:
	contentData := FindAndReadFile(path)

	// Case: Was not able to read the file.
	if contentData == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `The desired file was not found.`, path)
		http.NotFound(response, request)
		return
	}

	// Send the meta data and the content to the client:
	fileLenText := fmt.Sprintf(`%d`, len(contentData))
	response.Header().Add(`Content-Length`, fileLenText)
	response.Header().Add(`Content-Type`, fileType)
	response.Write(contentData)

	// Log all static files accesses?
	if logStaticFileRequests {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameBROWSER, `A static file was requested.`, path)
	}
	return
}

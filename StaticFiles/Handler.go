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

func HandlerStaticFiles(response http.ResponseWriter, request *http.Request) {
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	// Prepare the path:
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

	if fileType == `` {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNOTFOUND, `The mime type is unknown.`, path)
		http.NotFound(response, request)
		return
	}

	contentData := FindAndReadFile(path)
	if contentData == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `The desired file was not found.`, path)
		http.NotFound(response, request)
		return
	}

	fileLenText := fmt.Sprintf(`%d`, len(contentData))
	response.Header().Add(`Content-Length`, fileLenText)
	response.Header().Add(`Content-Type`, fileType)
	response.Write(contentData)

	if logStaticFileRequests {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameBROWSER, `A static file was requested.`, path)
	}
	return

	http.NotFound(response, request)
	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNOTFOUND, `The desired file is not part of the ZIP file.`, `Do you use an old version?`, path)
	return
}

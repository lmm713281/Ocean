package WebContent

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
	"strings"
)

func HandlerDeliverFramework(response http.ResponseWriter, request *http.Request) {
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	path := strings.Replace(request.URL.Path, "/framework/", "", 1)
	sendError := SendContent(response, path)

	if sendError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameWRITE, `Was not able to send the desired web content file.`, request.URL.Path, sendError.Error())
		http.NotFound(response, request)
	}
}

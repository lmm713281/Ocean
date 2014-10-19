package StaticFiles

import (
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
)

func HandlerMapStaticFiles2Root(response http.ResponseWriter, request *http.Request) {
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	if request.RequestURI == `/` {
		request.RequestURI = `/staticFiles/` + startFile4Map2Root
	} else {
		request.RequestURI = `/staticFiles` + request.RequestURI
	}

	HandlerStaticFiles(response, request)
}

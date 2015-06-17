package StaticFiles

import (
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
)

// Handler to map the static files to the root. Use it for static web sites.
func HandlerMapStaticFiles2Root(response http.ResponseWriter, request *http.Request) {

	// Case: The system goes down.
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

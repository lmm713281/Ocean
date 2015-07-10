package Version

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
)

// Handler for the access to Ocean's version
func HandlerVersion(response http.ResponseWriter, request *http.Request) {

	// Case: The system goes down now?
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	fmt.Fprintf(response, "%s", oceansVersion)
}

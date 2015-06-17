package Web

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/Web/Assets"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
)

// Handler for some CSS data for the web logger.
func HandlerCSSNormalize(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebCSS)
	fmt.Fprint(response, Assets.CSSNormalize)
}

// Handler for some CSS data for the web logger.
func HandlerCSSWebflow(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebCSS)
	fmt.Fprint(response, Assets.CSSWebflow)
}

// Handler for some CSS data for the web logger.
func HandlerCSSLog(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebCSS)
	fmt.Fprint(response, Assets.CSSLog)
}

// Handler for some JS for the web logger.
func HandlerJSModernizr(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSModernizr)
}

// Handler for some JS for the web logger.
func HandlerJSWebflow(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSWebflow)
}

// Handler for some JS for the web logger.
func HandlerJSjQuery(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSjQuery)
}

// Handler for some JS for the web logger.
func HandlerJSjQueryMap(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSjQueryMap)
}

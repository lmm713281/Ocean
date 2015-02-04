package Web

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/Web/Assets"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
)

func HandlerCSSNormalize(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebCSS)
	fmt.Fprint(response, Assets.CSSNormalize)
}

func HandlerCSSWebflow(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebCSS)
	fmt.Fprint(response, Assets.CSSWebflow)
}

func HandlerCSSLog(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebCSS)
	fmt.Fprint(response, Assets.CSSLog)
}

func HandlerJSModernizr(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSModernizr)
}

func HandlerJSWebflow(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSWebflow)
}

func HandlerJSjQuery(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSjQuery)
}

func HandlerJSjQueryMap(response http.ResponseWriter, request *http.Request) {

	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebJavaScript)
	fmt.Fprint(response, Assets.JSjQueryMap)
}

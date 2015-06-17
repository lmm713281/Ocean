package Tools

import (
	"net/http"
)

// Send the chosen language for the content.
func SendChosenLanguage(response http.ResponseWriter, lang Language) {
	response.Header().Add(`Content-Language`, lang.Language)
}

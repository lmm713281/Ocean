package Tools

import (
	"net/http"
)

func SendChosenLanguage(response http.ResponseWriter, lang Language) {
	response.Header().Add(`Content-Language`, lang.Language)
}

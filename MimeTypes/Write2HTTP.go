package MimeTypes

import "net/http"

func Write2HTTP(response http.ResponseWriter, mime MimeType) {
	response.Header().Add(`Content-Type`, mime.MimeType)
}

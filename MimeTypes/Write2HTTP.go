package MimeTypes

import (
	"net/http"
)

// Function to write a MIME type to a client.
func Write2HTTP(response http.ResponseWriter, mime MimeType) {
	response.Header().Add(`Content-Type`, mime.MimeType)
}

package Handlers

import (
	"net/http"
)

// Returns the muxer for the public web server.
func GetPublicMux() (mux *http.ServeMux) {
	mux = muxPublic
	return
}

// Returns the muxer for the private web server.
func GetAdminMux() (mux *http.ServeMux) {
	mux = muxAdmin
	return
}

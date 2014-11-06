package Handlers

import (
	"net/http"
)

func GetPublicMux() (mux *http.ServeMux) {
	mux = muxPublic
	return
}

func GetAdminMux() (mux *http.ServeMux) {
	mux = muxAdmin
	return
}

package Handlers

import (
	"net/http"
)

func AddPublicHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	muxPublic.HandleFunc(pattern, handler)
}

func AddAdminHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	muxAdmin.HandleFunc(pattern, handler)
}

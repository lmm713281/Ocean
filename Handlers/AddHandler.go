package Handlers

import (
	"net/http"
)

// Function to add a new public handler.
func AddPublicHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	muxPublic.HandleFunc(pattern, handler)
}

// Function to add a new private handler.
func AddAdminHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	muxAdmin.HandleFunc(pattern, handler)
}

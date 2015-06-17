package Handlers

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"net/http"
)

var (
	senderName LM.Sender      = `System::Handlers` // This is the name for logging event from this package
	muxPublic  *http.ServeMux = http.NewServeMux() // The muxer for the public web server
	muxAdmin   *http.ServeMux = http.NewServeMux() // The muxer for the private web server
)

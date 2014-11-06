package Handlers

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"net/http"
)

var (
	senderName LM.Sender      = `System::Handlers`
	muxPublic  *http.ServeMux = http.NewServeMux()
	muxAdmin   *http.ServeMux = http.NewServeMux()
)

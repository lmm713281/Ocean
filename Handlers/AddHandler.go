package Handlers

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"net/http"
)

// Function to add a new public handler.
func AddPublicHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// In case of an error, catch the error:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, fmt.Sprintf("Was not able to add a public handler, because the path '%s' is already in use. %s", pattern, err))
			return
		}
	}()

	// Add the handler:
	muxPublic.HandleFunc(pattern, handler)
}

// Function to add a new private handler.
func AddAdminHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// In case of an error, catch the error:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, fmt.Sprintf("Was not able to add a private admin handler, because the path '%s' is already in use. %s", pattern, err))
			return
		}
	}()

	// Add the handler:
	muxAdmin.HandleFunc(pattern, handler)
}

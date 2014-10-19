package Robots

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"net/http"
)

func HandlerRobots(response http.ResponseWriter, request *http.Request) {
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameNETWORK, `The robots.txt was requested.`, request.RemoteAddr)
	fmt.Fprintf(response, `%s`, robotsContent)
}

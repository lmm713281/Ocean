package Robots

import "fmt"
import "net/http"
import "github.com/SommerEngineering/Ocean/Shutdown"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func HandlerRobots(response http.ResponseWriter, request *http.Request) {
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameNETWORK, `The robots.txt was requested.`, request.RemoteAddr)
	fmt.Fprintf(response, `%s`, robotsContent)
}

package System

import "net/http"
import "github.com/SommerEngineering/Ocean/ICCC"
import "github.com/SommerEngineering/Ocean/WebContent"
import "github.com/SommerEngineering/Ocean/StaticFiles"
import "github.com/SommerEngineering/Ocean/NumGen"
import "github.com/SommerEngineering/Ocean/Robots"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func InitHandlers() {

	initSystem()
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all system handlers.`)

	http.HandleFunc(`/framework/`, WebContent.HandlerDeliverFramework)
	http.HandleFunc(`/staticFiles/`, StaticFiles.HandlerStaticFiles)
	http.HandleFunc(`/next/number`, NumGen.HandlerGetNext)
	http.HandleFunc(`/robots.txt`, Robots.HandlerRobots)
	http.HandleFunc(`/ICCC`, ICCC.ICCCHandler)

	if ConfigurationDB.Read(`MapStaticFiles2Root`) == "true" {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The static files are mapped to the root.`)
		http.HandleFunc(`/`, StaticFiles.HandlerMapStaticFiles2Root)
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Done with registering all system handler.`)
}

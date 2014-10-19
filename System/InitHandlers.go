package System

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/NumGen"
	"github.com/SommerEngineering/Ocean/Robots"
	"github.com/SommerEngineering/Ocean/StaticFiles"
	"github.com/SommerEngineering/Ocean/WebContent"
	"net/http"
)

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

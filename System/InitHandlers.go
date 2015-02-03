package System

import (
	"github.com/SommerEngineering/Ocean/BinaryAssets"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/NumGen"
	"github.com/SommerEngineering/Ocean/Robots"
	"github.com/SommerEngineering/Ocean/StaticFiles"
	"github.com/SommerEngineering/Ocean/WebContent"
)

func InitHandlers() {

	initSystem()
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all system handlers.`)

	// Public Handlers:
	Handlers.AddPublicHandler(`/framework/`, WebContent.HandlerDeliverFramework)
	Handlers.AddPublicHandler(`/staticFiles/`, StaticFiles.HandlerStaticFiles)
	Handlers.AddPublicHandler(`/next/number`, NumGen.HandlerGetNext)
	Handlers.AddPublicHandler(`/robots.txt`, Robots.HandlerRobots)
	Handlers.AddPublicHandler(`/ICCC`, ICCC.ICCCHandler)

	// Private Handlers:
	Handlers.AddAdminHandler(`/framework/`, WebContent.HandlerDeliverFramework)
	Handlers.AddAdminHandler(`/staticFiles/`, StaticFiles.HandlerStaticFiles)
	Handlers.AddAdminHandler(`/next/number`, NumGen.HandlerGetNext)
	Handlers.AddAdminHandler(`/ICCC`, ICCC.ICCCHandler)
	Handlers.AddAdminHandler(`/binaryAssets/`, BinaryAssets.HandlerBinaryAssets)

	if ConfigurationDB.Read(`MapStaticFiles2Root`) == "true" {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The static files are mapped to the root.`)
		Handlers.AddPublicHandler(`/`, StaticFiles.HandlerMapStaticFiles2Root)
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Done with registering all system handler.`)
}
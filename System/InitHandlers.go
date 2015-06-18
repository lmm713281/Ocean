package System

import (
	"github.com/SommerEngineering/Ocean/BinaryAssets"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Log/Web"
	"github.com/SommerEngineering/Ocean/Robots"
	"github.com/SommerEngineering/Ocean/StaticFiles"
	"github.com/SommerEngineering/Ocean/WebContent"
)

// Init the system and then the system's handlers.
func InitHandlers() {

	initSystem()
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all system handlers.`)

	//
	// Public Handlers:
	//

	// Handler for the web frameworks like e.g. jQuery, Bootstrap, etc.
	Handlers.AddPublicHandler(`/framework/`, WebContent.HandlerDeliverFramework)

	// Handler for other static files:
	Handlers.AddPublicHandler(`/staticFiles/`, StaticFiles.HandlerStaticFiles)

	// Handler for the robots.txt:
	Handlers.AddPublicHandler(`/robots.txt`, Robots.HandlerRobots)

	// Handler for the ICCC to the public:
	Handlers.AddPublicHandler(`/ICCC`, ICCC.ICCCHandler)

	//
	// Private Handlers:
	//

	// Handler for the web frameworks like e.g. jQuery, Bootstrap, etc.
	Handlers.AddAdminHandler(`/framework/`, WebContent.HandlerDeliverFramework)

	// Handler for other static files:
	Handlers.AddAdminHandler(`/staticFiles/`, StaticFiles.HandlerStaticFiles)

	// Handler for the ICCC to the private side:
	Handlers.AddAdminHandler(`/ICCC`, ICCC.ICCCHandler)

	// Handler for binary assets, used for the admin pages:
	Handlers.AddAdminHandler(`/binaryAssets/`, BinaryAssets.HandlerBinaryAssets)

	// Handler for the web logging:
	Handlers.AddAdminHandler(`/log`, Web.HandlerWebLog)

	// Handler for the web logging's CSS and JS:
	Handlers.AddAdminHandler(`/log/css/normalize.css`, Web.HandlerCSSNormalize)
	Handlers.AddAdminHandler(`/log/css/webflow.css`, Web.HandlerCSSWebflow)
	Handlers.AddAdminHandler(`/log/css/log.css`, Web.HandlerCSSLog)
	Handlers.AddAdminHandler(`/log/js/modernizr.js`, Web.HandlerJSModernizr)
	Handlers.AddAdminHandler(`/log/js/jquery.min.js`, Web.HandlerJSjQuery)
	Handlers.AddAdminHandler(`/log/js/jquery.min.map`, Web.HandlerJSjQueryMap)
	Handlers.AddAdminHandler(`/log/js/webflow.js`, Web.HandlerJSWebflow)

	// Are the static files mapped to the public?
	if ConfigurationDB.Read(`MapStaticFiles2Root`) == "true" {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The static files are mapped to the root.`)
		Handlers.AddPublicHandler(`/`, StaticFiles.HandlerMapStaticFiles2Root)
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Done with registering all system handler.`)
}

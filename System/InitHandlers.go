package System

import (
	"github.com/SommerEngineering/Ocean/Admin"
	"github.com/SommerEngineering/Ocean/BinaryAssets"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Log/Web"
	"github.com/SommerEngineering/Ocean/Robots"
	"github.com/SommerEngineering/Ocean/StaticFiles"
	"github.com/SommerEngineering/Ocean/System/Version"
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
	// Private/Admin Handlers:
	//

	// Handler for the web frameworks like e.g. jQuery, Bootstrap, etc.
	Handlers.AddAdminHandler(`/framework/`, WebContent.HandlerDeliverFramework)

	// Handler for other static files:
	Handlers.AddAdminHandler(`/staticFiles/`, StaticFiles.HandlerStaticFiles)

	// Handler for the ICCC to the private side:
	Handlers.AddAdminHandler(`/ICCC`, ICCC.ICCCHandler)

	// Handler for binary assets, used for the admin pages:
	Handlers.AddAdminHandler(`/binaryAssets/`, BinaryAssets.HandlerBinaryAssets)

	// Handler for the admin's overview:
	Handlers.AddAdminHandler(`/`, Admin.HandlerOverview)

	// Handler for the web logging:
	Handlers.AddAdminHandler(`/log`, Web.HandlerWebLog)

	// Handler for the access to Ocean's version:
	Handlers.AddAdminHandler(`/version`, Version.HandlerVersion)

	// Handler for the file upload:
	Handlers.AddAdminHandler(`/upload`, Admin.HandlerFileUpload)

	// Handler for the configuration view:
	Handlers.AddAdminHandler(`/configuration`, Admin.HandlerConfiguration)

	// Handler for the admin area:
	Handlers.AddAdminHandler(`/admin/css/normalize.css`, Admin.HandlerCSSNormalize)
	Handlers.AddAdminHandler(`/admin/css/webflow.css`, Admin.HandlerCSSWebflow)
	Handlers.AddAdminHandler(`/admin/css/admin.css`, Admin.HandlerCSSAdmin)
	Handlers.AddAdminHandler(`/admin/js/modernizr.js`, Admin.HandlerJSModernizr)
	Handlers.AddAdminHandler(`/admin/js/jquery.min.js`, Admin.HandlerJSjQuery)
	Handlers.AddAdminHandler(`/admin/js/jquery.min.map`, Admin.HandlerJSjQueryMap)
	Handlers.AddAdminHandler(`/admin/js/webflow.js`, Admin.HandlerJSWebflow)

	// Are the static files mapped to the public?
	if ConfigurationDB.Read(`MapStaticFiles2Root`) == "true" {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The static files are mapped to the root.`)
		Handlers.AddPublicHandler(`/`, StaticFiles.HandlerMapStaticFiles2Root)
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Done with registering all system handler.`)
}

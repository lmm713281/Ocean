package WebServer

import (
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"time"
)

func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Init the web server now.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Done init the web server.`)

	serverPublicAddressPort := Tools.LocalIPAddressAndPort()

	serverPublic = &http.Server{}
	serverPublic.Addr = serverPublicAddressPort
	serverPublic.Handler = Handlers.GetPublicMux()
	serverPublic.ReadTimeout = 10 * time.Second
	serverPublic.WriteTimeout = 10 * time.Second
	serverPublic.MaxHeaderBytes = 1024
	serverPublic.SetKeepAlivesEnabled(true)

	serverAdmin = &http.Server{}
	serverAdmin.Addr = serverAdminAddressPort
	serverAdmin.Handler = Handlers.GetAdminMux()
	serverAdmin.ReadTimeout = 10 * time.Second
	serverAdmin.WriteTimeout = 10 * time.Second
	serverAdmin.MaxHeaderBytes = 1024
	serverAdmin.SetKeepAlivesEnabled(true)
}

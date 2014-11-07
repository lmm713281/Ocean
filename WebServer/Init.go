package WebServer

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Init the web server now.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Done init the web server.`)

	serverPublicAddressPort = Tools.LocalIPAddressAndPort()
	serverAdminAddressPort = ConfigurationDB.Read(`AdminWebServerBinding`)

	serverPublic = &http.Server{}
	serverPublic.Addr = serverPublicAddressPort
	serverPublic.Handler = Handlers.GetPublicMux()
	serverPublic.SetKeepAlivesEnabled(true)

	serverPublic.ReadTimeout = 10 * time.Second
	serverPublic.WriteTimeout = 10 * time.Second
	serverPublic.MaxHeaderBytes = 1024

	if strings.ToLower(ConfigurationDB.Read(`AdminWebServerEnabled`)) == `true` {
		serverAdmin = &http.Server{}
		serverAdmin.Addr = serverAdminAddressPort
		serverAdmin.Handler = Handlers.GetAdminMux()
		serverAdmin.SetKeepAlivesEnabled(true)

		if readTimeoutSeconds, errTimeout := strconv.Atoi(ConfigurationDB.Read(`AdminWebServerReadTimeoutSeconds`)); errTimeout != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the admin server's read timeout value. Use the default of 10 seconds instead.`, errTimeout.Error())
			serverAdmin.ReadTimeout = 10 * time.Second
		} else {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The admin web server's read timeout was set to %d seconds.", readTimeoutSeconds))
			serverAdmin.ReadTimeout = time.Duration(readTimeoutSeconds) * time.Second
		}

		if writeTimeoutSeconds, errTimeout := strconv.Atoi(ConfigurationDB.Read(`AdminWebServerWriteTimeoutSeconds`)); errTimeout != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the admin server's write timeout value. Use the default of 10 seconds instead.`, errTimeout.Error())
			serverAdmin.WriteTimeout = 10 * time.Second
		} else {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The admin web server's write timeout was set to %d seconds.", writeTimeoutSeconds))
			serverAdmin.WriteTimeout = time.Duration(writeTimeoutSeconds) * time.Second
		}

		if maxHeaderBytes, errHeaderBytes := strconv.Atoi(ConfigurationDB.Read(`AdminWebServerMaxHeaderLenBytes`)); errHeaderBytes != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the admin server's maximum count of bytes for the headers. Use the default of 1048576 bytes instead.`, errHeaderBytes.Error())
			serverAdmin.MaxHeaderBytes = 1048576
		} else {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The admin web server's count of maximal bytes for the headers was set to %d bytes.", maxHeaderBytes))
			serverAdmin.MaxHeaderBytes = maxHeaderBytes
		}
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The admin web server is disabled.`)
	}
}

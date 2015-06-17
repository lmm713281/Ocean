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

// Setup the web server.
func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Init the web server now.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Done init the web server.`)

	// Public web server: use the servers public facing IP address and the configured port:
	serverPublicAddressPort = Tools.LocalIPAddressAndPort()

	// Private web server: use the configured end point:
	serverAdminAddressPort = ConfigurationDB.Read(`AdminWebServerBinding`)

	// Setup the public web server:
	serverPublic = &http.Server{}
	serverPublic.Addr = serverPublicAddressPort
	serverPublic.Handler = Handlers.GetPublicMux()
	serverPublic.SetKeepAlivesEnabled(true)

	// Public Web Server: Read Timeout
	if readTimeoutSeconds, errTimeout := strconv.Atoi(ConfigurationDB.Read(`PublicWebServerReadTimeoutSeconds`)); errTimeout != nil {

		// Case: Error! Use a default timeout:
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the public server's read timeout value. Use the default of 10 seconds instead.`, errTimeout.Error())
		serverPublic.ReadTimeout = 10 * time.Second
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The public web server's read timeout was set to %d seconds.", readTimeoutSeconds))
		serverPublic.ReadTimeout = time.Duration(readTimeoutSeconds) * time.Second
	}

	// Public Web Server: Write Timeout
	if writeTimeoutSeconds, errTimeout := strconv.Atoi(ConfigurationDB.Read(`PublicWebServerWriteTimeoutSeconds`)); errTimeout != nil {

		// Case: Error! Use a default timeout:
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the public server's write timeout value. Use the default of 10 seconds instead.`, errTimeout.Error())
		serverPublic.WriteTimeout = 10 * time.Second
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The public web server's write timeout was set to %d seconds.", writeTimeoutSeconds))
		serverPublic.WriteTimeout = time.Duration(writeTimeoutSeconds) * time.Second
	}

	// Public Web Server: Max. Header Size
	if maxHeaderBytes, errHeaderBytes := strconv.Atoi(ConfigurationDB.Read(`PublicWebServerMaxHeaderLenBytes`)); errHeaderBytes != nil {

		// Case: Error! Use a default threshold for the header size:
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the public server's max. header size. Use the default of 1048576 bytes instead.`, errHeaderBytes.Error())
		serverPublic.MaxHeaderBytes = 1048576
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The public web server's max. header size was set to %d bytes.", maxHeaderBytes))
		serverPublic.MaxHeaderBytes = maxHeaderBytes
	}

	// Is the private web server (i.e. administration server) enabled?
	if strings.ToLower(ConfigurationDB.Read(`AdminWebServerEnabled`)) == `true` {

		// Setup the private web server:
		serverAdmin = &http.Server{}
		serverAdmin.Addr = serverAdminAddressPort
		serverAdmin.Handler = Handlers.GetAdminMux()
		serverAdmin.SetKeepAlivesEnabled(true)

		// Admin Web Server: Read Timeout
		if readTimeoutSeconds, errTimeout := strconv.Atoi(ConfigurationDB.Read(`AdminWebServerReadTimeoutSeconds`)); errTimeout != nil {

			// Case: Error! Use a default timeout:
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the admin server's read timeout value. Use the default of 10 seconds instead.`, errTimeout.Error())
			serverAdmin.ReadTimeout = 10 * time.Second
		} else {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The admin web server's read timeout was set to %d seconds.", readTimeoutSeconds))
			serverAdmin.ReadTimeout = time.Duration(readTimeoutSeconds) * time.Second
		}

		// Admin Web Server: Write Timeout
		if writeTimeoutSeconds, errTimeout := strconv.Atoi(ConfigurationDB.Read(`AdminWebServerWriteTimeoutSeconds`)); errTimeout != nil {

			// Case: Error! Use a default timeout:
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the admin server's write timeout value. Use the default of 10 seconds instead.`, errTimeout.Error())
			serverAdmin.WriteTimeout = 10 * time.Second
		} else {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The admin web server's write timeout was set to %d seconds.", writeTimeoutSeconds))
			serverAdmin.WriteTimeout = time.Duration(writeTimeoutSeconds) * time.Second
		}

		// Admin Web Server: Max. Header Size
		if maxHeaderBytes, errHeaderBytes := strconv.Atoi(ConfigurationDB.Read(`AdminWebServerMaxHeaderLenBytes`)); errHeaderBytes != nil {

			// Case: Error! Use a default threshold for the header size:
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameCONFIGURATION, `Was not able to read the admin server's max. header size. Use the default of 1048576 bytes instead.`, errHeaderBytes.Error())
			serverAdmin.MaxHeaderBytes = 1048576
		} else {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("The admin web server's max. header size was set to %d bytes.", maxHeaderBytes))
			serverAdmin.MaxHeaderBytes = maxHeaderBytes
		}
	} else {
		// Private web server is disabled:
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The admin web server is disabled.`)
	}
}

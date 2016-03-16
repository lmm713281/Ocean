package WebServer

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/StaticFiles"
	"github.com/SommerEngineering/Ocean/Tools"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

	// Is TLS configured?
	if publicTLSEnabled := ConfigurationDB.Read(`PublicWebServerUseTLS`); strings.ToLower(publicTLSEnabled) == `true` {

		// TLS is enabled. Copy the certificate and private key to the source directory.
		publicTLSCertificate := StaticFiles.FindAndReadFileINTERNAL(ConfigurationDB.Read(`PublicWebServerTLSCertificateName`))
		publicTLSPrivateKey := StaticFiles.FindAndReadFileINTERNAL(ConfigurationDB.Read(`PublicWebServerTLSPrivateKey`))

		// Access to the working directory?
		currentDir, dirError := os.Getwd()
		if dirError != nil {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameCONFIGURATION, `Was not able to read the working directory. Thus, cannot store the TLS certificates!`, dirError.Error())
		} else {
			// Build the filenames:
			pathCertificate := filepath.Join(currentDir, ConfigurationDB.Read(`PublicWebServerTLSCertificateName`))
			pathPrivateKey := filepath.Join(currentDir, ConfigurationDB.Read(`PublicWebServerTLSPrivateKey`))

			// Write the files:
			if writeError := ioutil.WriteFile(pathCertificate, publicTLSCertificate, 0660); writeError != nil {
				Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameCONFIGURATION, `Was not able to write the TLS certificate to the working directory.`, writeError.Error())
			}
			if writeError := ioutil.WriteFile(pathPrivateKey, publicTLSPrivateKey, 0660); writeError != nil {
				Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameCONFIGURATION, `Was not able to write the TLS private key to the working directory.`, writeError.Error())
			}
		}
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

		// Is TLS configured?
		if adminTLSEnabled := ConfigurationDB.Read(`AdminWebServerUseTLS`); strings.ToLower(adminTLSEnabled) == `true` {

			// TLS is enabled. Copy the certificate and private key to the source directory.
			adminTLSCertificate := StaticFiles.FindAndReadFileINTERNAL(ConfigurationDB.Read(`AdminWebServerTLSCertificateName`))
			adminTLSPrivateKey := StaticFiles.FindAndReadFileINTERNAL(ConfigurationDB.Read(`AdminWebServerTLSPrivateKey`))

			// Access to the working directory?
			currentDir, dirError := os.Getwd()
			if dirError != nil {
				Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameCONFIGURATION, `Was not able to read the working directory. Thus, cannot store the TLS certificates!`, dirError.Error())
			} else {
				// Build the filenames:
				pathCertificate := filepath.Join(currentDir, ConfigurationDB.Read(`AdminWebServerTLSCertificateName`))
				pathPrivateKey := filepath.Join(currentDir, ConfigurationDB.Read(`AdminWebServerTLSPrivateKey`))

				// Write the files:
				if writeError := ioutil.WriteFile(pathCertificate, adminTLSCertificate, 0660); writeError != nil {
					Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameCONFIGURATION, `Was not able to write the TLS certificate to the working directory.`, writeError.Error())
				}
				if writeError := ioutil.WriteFile(pathPrivateKey, adminTLSPrivateKey, 0660); writeError != nil {
					Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameCONFIGURATION, `Was not able to write the TLS private key to the working directory.`, writeError.Error())
				}
			}
		}
	} else {
		// Private web server is disabled:
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The admin web server is disabled.`)
	}
}

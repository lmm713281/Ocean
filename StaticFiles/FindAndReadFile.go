package StaticFiles

import (
	"archive/zip"
	"bytes"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"io/ioutil"
	"strings"
)

// Try to read a static file.
func FindAndReadFile(filename string) (result []byte) {

	// Case: The system goes down.
	if Shutdown.IsDown() {
		return
	}

	//
	// Ensure that the TLS keys are secure and save:
	//
	if strings.ToLower(filename) == strings.ToLower(ConfigurationDB.Read(`AdminWebServerTLSCertificateName`)) {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityNone, LM.ImpactNone, LM.MessageNameREQUEST, `Someone tried to read the TLS certificate of the admin server. The attempt was inhibited.`)
		return
	}

	if strings.ToLower(filename) == strings.ToLower(ConfigurationDB.Read(`AdminWebServerTLSPrivateKey`)) {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityNone, LM.ImpactNone, LM.MessageNameREQUEST, `Someone tried to read the TLS certificate's private key of the admin server. The attempt was inhibited.`)
		return
	}

	if strings.ToLower(filename) == strings.ToLower(ConfigurationDB.Read(`PublicWebServerTLSCertificateName`)) {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityNone, LM.ImpactNone, LM.MessageNameREQUEST, `Someone tried to read the TLS certificate of the public server. The attempt was inhibited.`)
		return
	}

	if strings.ToLower(filename) == strings.ToLower(ConfigurationDB.Read(`PublicWebServerTLSPrivateKey`)) {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityNone, LM.ImpactNone, LM.MessageNameREQUEST, `Someone tried to read the TLS certificate's private key of the public server. The attempt was inhibited.`)
		return
	}

	result = FindAndReadFileINTERNAL(filename)
	return
}

func FindAndReadFileINTERNAL(filename string) (result []byte) {

	// Case: The system goes down.
	if Shutdown.IsDown() {
		return
	}

	// Prepare the path:
	path := filename

	// Read the content from the ZIP file:
	reader, readerError := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if readerError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the ZIP file.`, readerError.Error())
		return
	}

	// Loop over all files inside the ZIP file:
	for _, file := range reader.File {

		// Is this the desired file?
		if file.Name == path {

			// Open the file:
			fileReader, openError := file.Open()
			defer fileReader.Close()
			if openError == nil {
				// Read all the content:
				contentData, readError := ioutil.ReadAll(fileReader)

				if readError != nil {
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the content of the desired file.`, readError.Error(), path)
					return
				}

				result = contentData
				return
			}
		}
	}

	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNOTFOUND, `The desired file is not part of the ZIP file.`, `Do you use an old version?`, path)
	return
}

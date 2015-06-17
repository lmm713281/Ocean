package StaticFiles

import (
	"archive/zip"
	"bytes"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
	"io/ioutil"
)

// Try to read a static file.
func FindAndReadFile(filename string) (result []byte) {

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

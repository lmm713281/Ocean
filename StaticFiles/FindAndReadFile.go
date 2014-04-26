package StaticFiles

import "io/ioutil"
import "bytes"
import "archive/zip"
import "github.com/SommerEngineering/Ocean/Shutdown"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func FindAndReadFile(filename string) (result []byte) {
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

	for _, file := range reader.File {
		if file.Name == path {

			fileReader, openError := file.Open()
			defer fileReader.Close()

			if openError == nil {
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

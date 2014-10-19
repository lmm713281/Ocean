package WebContent

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"io/ioutil"
)

func GetContent(path string) (content []byte, err error) {

	reader, readerError := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if readerError != nil {
		err = errors.New("Was not able to read the ZIP file: " + readerError.Error())
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
					err = errors.New("Was not able to read the content of the desired file: " + readError.Error())
					Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the content of the desired file.`, readError.Error(), path)
					return
				}

				content = contentData
				return
			}
		}
	}

	err = errors.New("File not found!")
	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameNOTFOUND, `The desired file is not part of the ZIP file.`, `Do you use an old version?`, path)
	return
}

package Templates

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"html/template"
	"io/ioutil"
)

func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the template engine.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the template engine done.`)

	dbSession, gridFS := CustomerDB.GridFS()
	defer dbSession.Close()

	if gridFile, errGridFile := gridFS.Open(`templates.zip`); errGridFile != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to open the templates out of the GridFS!`, errGridFile.Error())
		return
	} else {
		defer gridFile.Close()
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Read the templates.zip file from the grid file system.`, `Upload time UTC: `+Tools.FormatTime(gridFile.UploadDate().UTC()))
		if data, ioError := ioutil.ReadAll(gridFile); ioError != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to read the templates.`, ioError.Error())
			return
		} else {
			zipData = data
		}
	}

	// Read the content from the ZIP file:
	reader, readerError := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if readerError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the ZIP file.`, readerError.Error())
		return
	}

	templates = template.New(`root`)
	for _, file := range reader.File {

		fileReader, openError := file.Open()
		if openError == nil {
			contentData, readError := ioutil.ReadAll(fileReader)
			fileReader.Close()

			if readError != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the content of the desired template.`, readError.Error(), file.FileInfo().Name())
				continue
			}

			templateData := string(contentData)
			templates.Parse(templateData)
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameEXECUTE, fmt.Sprintf(`The template '%s' was parsed.`, file.FileInfo().Name()))
		} else {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to open a template.`, file.FileInfo().Name())
		}
	}

	isInit = true
}

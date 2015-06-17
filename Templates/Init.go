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

// Init the template package.
func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the template engine.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Starting the template engine done.`)

	// We read the templates out of the grid file system:
	dbSession, gridFS := CustomerDB.GridFS()
	defer dbSession.Close()

	// Read the current version of the ZIP file:
	if gridFile, errGridFile := gridFS.Open(`templates.zip`); errGridFile != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to open the templates out of the GridFS!`, errGridFile.Error())
		return
	} else {
		// Read all data in the memory cache:
		defer gridFile.Close()
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Read the templates.zip file from the grid file system.`, `Upload time UTC: `+Tools.FormatTime(gridFile.UploadDate().UTC()))
		if data, ioError := ioutil.ReadAll(gridFile); ioError != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to read the templates.`, ioError.Error())
			return
		} else {
			zipData = data
		}
	}

	// Opens the ZIP file from memory:
	reader, readerError := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if readerError != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the ZIP file.`, readerError.Error())
		return
	}

	// The in-memory container for all templates:
	templates = template.New(`root`)

	// Loop over all files inside the ZIP file:
	for _, file := range reader.File {

		// Opens a file:
		fileReader, openError := file.Open()
		if openError == nil {

			// Read all bytes:
			contentData, readError := ioutil.ReadAll(fileReader)
			fileReader.Close()

			if readError != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the content of the desired template.`, readError.Error(), file.FileInfo().Name())
				continue
			}

			// Converts the bytes to string:
			templateData := string(contentData)

			// Try to parse the string as template:
			if _, err := templates.Parse(templateData); err != nil {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactMiddle, LM.MessageNamePARSE, fmt.Sprintf(`The template '%s' cannot be parsed.`, file.FileInfo().Name()), err.Error())
			} else {
				Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameEXECUTE, fmt.Sprintf(`The template '%s' was parsed.`, file.FileInfo().Name()))
			}
		} else {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to open a template.`, file.FileInfo().Name())
		}
	}

	isInit = true
}

package WebContent

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"net/http"
)

// Send any web content file to the client.
func SendContent(response http.ResponseWriter, path string) (err error) {

	// Read the bytes:
	content, contError := GetContent(path)
	if contError != nil {
		err = errors.New("Was not able to read the needed content: " + contError.Error())
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameREAD, `Was not able to read the needed content.`, contError.Error(), path)
		return
	}

	// Determine some meta data:
	contentLength := len(content)
	contentType, typeError := MimeTypes.DetectType(path)

	if typeError != nil {
		err = errors.New("Was not able to detect the MIME type: " + typeError.Error())
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameANALYSIS, `Was not able to detect the MIME type for the file.`, path, typeError.Error())
		return
	}

	// Send the meta data to the client:
	response.Header().Add("Content-Length", fmt.Sprintf("%d", contentLength))
	response.Header().Add("Content-Type", contentType.MimeType)
	response.WriteHeader(http.StatusOK)

	// Write the binary data to the client:
	buffer := bytes.NewBuffer(content)
	_, writeError := buffer.WriteTo(response)

	if writeError != nil {
		err = errors.New("Was not able to write the data to the net: " + writeError.Error())
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameWRITE, `Was not able to write the file to the browser.`, path, writeError.Error())
		return
	}

	return
}

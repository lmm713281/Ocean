package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Send a message to a listener.
func sendMessage(listener Scheme.Listener, data map[string][]string) (result map[string][]string) {

	// Lets sign the data:
	valuesHTTP := signMessage(data)

	// Try to deliver the message:
	if response, err := http.PostForm(activeProtocol+listener.Hostname+`/ICCC`, valuesHTTP); err != nil {
		// Case: Was not possible to deliver.
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNETWORK, `Was not able to send the ICCC message.`, err.Error())
	} else {
		// Case: Delivery was fine.
		defer response.Body.Close()
		if responseData, err := ioutil.ReadAll(response.Body); err != nil {
			// Case: Was not possible to read the answer.
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNETWORK, `Was not able to read the ICCC answer.`, err.Error())
		} else {
			// Case: Was able to read the answer.
			if dataObj, errObj := url.ParseQuery(string(responseData)); errObj != nil {
				// Case: Was not able to parse the answer to values.
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNETWORK, `Was not able to parse the answer to values.`, errObj.Error())
			} else {
				// Case: Everything was fine.
				result = map[string][]string(dataObj)
			}
		}
	}

	return
}

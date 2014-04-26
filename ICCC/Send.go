package ICCC

import "net/http"
import "net/url"
import "github.com/SommerEngineering/Ocean/Tools"
import "github.com/SommerEngineering/Ocean/ICCC/Scheme"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func sendMessage(listener Scheme.Listener, data map[string][]string) {

	valuesHTTP := url.Values(data)
	valuesHTTP.Add(`InternalCommPassword`, Tools.InternalCommPassword())
	if _, err := http.PostForm(`http://`+listener.IPAddressPort+`/ICCC`, valuesHTTP); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNETWORK, `Was not able to send the ICCC message.`, err.Error())
	}

	return
}

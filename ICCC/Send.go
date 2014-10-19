package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"net/url"
)

func sendMessage(listener Scheme.Listener, data map[string][]string) {

	valuesHTTP := url.Values(data)
	valuesHTTP.Add(`InternalCommPassword`, Tools.InternalCommPassword())
	if _, err := http.PostForm(`http://`+listener.IPAddressPort+`/ICCC`, valuesHTTP); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameNETWORK, `Was not able to send the ICCC message.`, err.Error())
	}

	return
}

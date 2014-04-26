package NumGen

import "net/http"
import "net/url"
import "strconv"
import "github.com/SommerEngineering/Ocean/Shutdown"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func GetNextInt64(name string) (result int64) {
	result = badNumber64

	if Shutdown.IsDown() {
		return
	}

	if responseData, errRequest := http.PostForm(getHandler, url.Values{"name": {name}, "password": {correctPassword}}); errRequest != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameGENERATOR, `Requesting the next number was not possible.`, errRequest.Error())
		return
	} else {
		nextNumberText := responseData.Header.Get(`nextNumber`)
		if number, errAtio := strconv.Atoi(nextNumberText); errAtio != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameGENERATOR, `It was not possible to convert the answer into an int64.`, errAtio.Error())
			return
		} else {
			result = int64(number)
			return
		}
	}
}

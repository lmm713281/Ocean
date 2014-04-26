package NumGen

import "fmt"
import "net/http"
import "github.com/SommerEngineering/Ocean/Shutdown"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func HandlerGetNext(response http.ResponseWriter, request *http.Request) {
	if Shutdown.IsDown() {
		http.NotFound(response, request)
		return
	}

	if !isActive {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactNone, LM.MessageNameCONFIGURATION, `Called the get handler on an inactive host.`, `Wrong configuration?`)
		http.NotFound(response, request)
		return
	}

	if correctPassword == `` {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSECURITY, `No communication password was set.`)
		http.NotFound(response, request)
		return
	}

	name := request.FormValue(`name`)
	pwd := request.FormValue(`password`)

	if pwd != correctPassword {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityCritical, LM.ImpactNone, LM.MessageNamePASSWORD, `A wrong password was used to access this system handler.`, `This should never happens: Is this a hacking attempt?`, `IP address of requester=`+request.RemoteAddr)
		http.NotFound(response, request)
		return
	}

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelDEBUG, LM.MessageNameANALYSIS, `Next number requested.`, name, pwd)
	channel := requestChannel4Name(name)
	nextNumber := <-channel

	response.Header().Add(`nextNumber`, fmt.Sprintf(`%d`, nextNumber))
}

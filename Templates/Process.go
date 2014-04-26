package Templates

import "net/http"
import "github.com/SommerEngineering/Ocean/MimeTypes"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func ProcessHTML(templateName string, response http.ResponseWriter, data interface{}) {

	if !isInit {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameINIT, `The template engine is not (yet) init.`)
		return
	}

	MimeTypes.Write2HTTP(response, MimeTypes.TypeWebHTML)
	if executeError := templates.ExecuteTemplate(response, templateName, data); executeError != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameEXECUTE, `Was not able to execute the template.`, templateName, executeError.Error())
	}
}

package Web

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	WebTemp "github.com/SommerEngineering/Ocean/Log/Web/Templates"
	"html/template"
)

// The init function for this package.
func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the web log.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the web log done.`)

	// Create the cache of all web logging templates:
	templates = template.New(`root`)
	if _, err := templates.Parse(WebTemp.Viewer); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to parse the template for the web log viewer.`, err.Error())
	}
}

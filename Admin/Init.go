package Admin

import (
	"github.com/SommerEngineering/Ocean/Admin/Templates"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"html/template"
)

// The init function for this package.
func init() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the admin area.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Init the admin area done.`)

	// Create the cache of all admin templates:
	AdminTemplates = template.New(`root`)
	if _, err := AdminTemplates.Parse(Templates.LoggingViewer); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to parse the template for the web log viewer.`, err.Error())
	}

	if _, err := AdminTemplates.Parse(Templates.Overview); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to parse the template for the admin overview.`, err.Error())
	}

	if _, err := AdminTemplates.Parse(Templates.FileUpload); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to parse the template for the file upload.`, err.Error())
	}

	if _, err := AdminTemplates.Parse(Templates.Configuration); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNamePARSE, `Was not able to parse the template for the configuration.`, err.Error())
	}
}

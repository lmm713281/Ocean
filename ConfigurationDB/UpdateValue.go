package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// This function updates a configuration value e.g. from the admin's configuration web interface.
func UpdateValue(name string, configuration ConfigurationDBEntry) {

	// Check the configuration's name:
	if name == `` {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, `Was not able to update a configuration value.`, `The given name was nil!`)
		return
	}

	// Ensure, that the configuration is already present:
	if count, errFind := collection.Find(bson.D{{"Name", name}}).Count(); errFind != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to update a configuration value.`, `Error while find the old value.`, errFind.Error())
		return
	} else {
		// Is the configuration already present?
		if count == 0 {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, `Was not able to update a configuration value.`, `The desired configuration is not present.`)
			return
		}
	}

	// Ensure, that the configuration value also uses the same name:
	if name != configuration.Name {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameSTATE, `Was not able to update a configuration value.`, `The given name was different with the name of the desired configuration value.`)
		return
	}

	//
	// Case: Any precondition is fulfilled
	//

	// Selection of the correct configuration (the name is a unique value):
	selector := bson.D{{"Name", name}}
	if errUpdate := collection.Update(selector, configuration); errUpdate != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to update a configuration value.`, `Error while updating the database.`, errUpdate.Error())
	}

	return
}

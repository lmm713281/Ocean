package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// This function reads the current configuration value.
func Read(name string) (value string) {
	if name == `` {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to read a configuration out of the database.`, `The given name was nil!`)
		return
	}

	result := ConfigurationDBEntry{}
	if errFind := collection.Find(bson.D{{"Name", name}}).One(&result); errFind != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to read a configuration out of the database.`, `Error while find.`, errFind.Error())
		return
	}

	value = result.Value
	return
}

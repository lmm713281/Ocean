package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

/*
This function writes the configuration value.
*/
func Write(name, value string) {
	if name == `` {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to write a configuration to the database.`, `The given name was nil!`)
		return
	}

	result := ConfigurationDBEntry{}
	if errFind := collection.Find(bson.D{{"Name", name}}).One(&result); errFind != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to write a configuration to the database.`, `Error while find.`, errFind.Error())
		return
	}

	result.Value = value
	collection.Update(bson.D{{"Name", name}}, result)
	return
}

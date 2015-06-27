package ConfigurationDB

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// This function reads all configuration values e.g. for the admin's configuration web interface.
func ReadAll() (values []ConfigurationDBEntry) {
	if errFind := collection.Find(bson.D{}).All(&values); errFind != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to read all configuration values out of the database.`, `Error while find.`, errFind.Error())
	}

	return
}

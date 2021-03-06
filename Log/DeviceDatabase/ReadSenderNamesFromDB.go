package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
	"sort"
)

// Reads the sender names from the database without any caching.
func readSenderNamesFromDB() (result []Scheme.Sender) {
	var nextSenderNames []string
	if err := logDBCollection.Find(bson.D{{"Project", projectName}}).Distinct(`Sender`, &nextSenderNames); err != nil {
		// Case: Was not possible to write to the database.
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameDATABASE, `Was not able to read the sender names from the database.`, err.Error())
		return
	}

	// Sort the sender names:
	sort.Strings(nextSenderNames)

	// Transform the values to the right format:
	for _, entry := range nextSenderNames {
		result = append(result, Scheme.Sender(entry))
	}
	return
}

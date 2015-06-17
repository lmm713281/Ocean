package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
	"sort"
)

// Reads the sender names from the database without any caching.
func readSenderNamesFromDB() (result []string) {
	var nextSenderNames []string
	if err := logDBCollection.Find(bson.D{}).Distinct(`Sender`, &nextSenderNames); err != nil {
		// Case: Was not possible to write to the database.
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameDATABASE, `Was not able to read the sender names from the database.`, err.Error())
		return
	}

	// Sort the sender names:
	sort.Strings(nextSenderNames)
	result = nextSenderNames
	return
}

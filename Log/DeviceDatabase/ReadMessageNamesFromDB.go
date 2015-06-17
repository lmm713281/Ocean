package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
	"sort"
)

// Read the message names from the database without any cache.
func readMessageNamesFromDB() (result []string) {
	var nextMessageNames []string
	if err := logDBCollection.Find(bson.D{}).Distinct(`MessageName`, &nextMessageNames); err != nil {
		// Case: Error, was not able to write the event to the database:
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.MessageNameDATABASE, `Was not able to read the message names from the database.`, err.Error())
		return
	}

	// Sort the sender names:
	sort.Strings(nextMessageNames)
	result = nextMessageNames
	return
}

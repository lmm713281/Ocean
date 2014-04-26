package NumGen

import "labix.org/v2/mgo"
import "github.com/SommerEngineering/Ocean/CustomerDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func initDB() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Start init of number generator collection.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Done init of number generator collection.`)

	// Get the database:
	db = CustomerDB.DB()

	if db == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	// Get my collection:
	collectionNumGen = db.C(`NumGen`)

	// Take care about the indexes:
	indexName := mgo.Index{}
	indexName.Key = []string{`Name`}
	indexName.Unique = true
	collectionNumGen.EnsureIndex(indexName)
}

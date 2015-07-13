package ICCC

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
)

// Init the database.
func initDB() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Start init of the ICCC collections.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Done init the ICCC collection.`)

	// Get the database:
	dbSession, db = CustomerDB.DB()

	// Case: Error?
	if db == nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	// Get my collections:
	collectionListener = db.C(`ICCCListener`)
	collectionHosts = db.C(`ICCCHosts`)

	//
	// Take care about the indexes for ICCCListener:
	//
	collectionListener.EnsureIndexKey(`Kind`)
	collectionListener.EnsureIndexKey(`Command`)
	collectionListener.EnsureIndexKey(`Command`, `IsActive`)

	collectionListener.EnsureIndexKey(`Command`, `Channel`)
	collectionListener.EnsureIndexKey(`Command`, `Channel`, `IsActive`)

	collectionListener.EnsureIndexKey(`Channel`)
	collectionListener.EnsureIndexKey(`Channel`, `IsActive`)
	collectionListener.EnsureIndexKey(`Channel`, `Command`, `IPAddressPort`, `IsActive`)
	collectionListener.EnsureIndexKey(`Channel`, `Command`, `IsActive`)

	collectionListener.EnsureIndexKey(`IsActive`)
	collectionListener.EnsureIndexKey(`IsActive`, `IPAddressPort`)

	collectionListener.EnsureIndexKey(`IPAddressPort`)

	indexName1 := mgo.Index{}
	indexName1.Key = []string{`Channel`, `Command`, `IPAddressPort`}
	indexName1.Unique = true
	collectionListener.EnsureIndex(indexName1)

	//
	// Index for hosts:
	//
	collectionHosts.EnsureIndexKey(`Kind`)
	collectionHosts.EnsureIndexKey(`Hostname`)
	collectionHosts.EnsureIndexKey(`IPAddressPort`)

	indexName2 := mgo.Index{}
	indexName2.Key = []string{`Hostname`, `IPAddressPort`}
	indexName2.Unique = true
	collectionHosts.EnsureIndex(indexName2)
}

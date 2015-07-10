package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// Type to provide a shutdown function.
type ShutdownFunction struct {
}

// The shutdown function for ICCC.
func (a ShutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Shutting down now all ICCC listener for this host.`)

	// Define the database query:
	selection := bson.D{{`IPAddressPort`, correctAddressWithPort}}

	// Reserve the memory for an answer:
	entry := Scheme.Listener{}

	// Execute the query and iterate over the results:
	iterator := collectionListener.Find(selection).Iter()
	for iterator.Next(&entry) {
		// Update the entry and set it to active=false:
		selectionUpdate := bson.D{{`Channel`, entry.Channel}, {`Command`, entry.Command}, {`IPAddressPort`, correctAddressWithPort}}
		entry.IsActive = false

		// Update the entry:
		collectionListener.Update(selectionUpdate, entry)
	}

	// Disconnect the database:
	db.Logout()
	dbSession.Close()
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Done shutting down all ICCC listener for this host.`)
}

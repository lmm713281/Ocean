package ICCC

import "labix.org/v2/mgo/bson"
import "github.com/SommerEngineering/Ocean/ICCC/Scheme"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

/*
Please do not use this type. It is an internal type of Ocean to provide a shutdown function!
*/
type ShutdownFunction struct {
}

/*
This function is called if the Ocean server is shutting down.
*/
func (a ShutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Shutting down now all ICCC listener for this host.`)

	selection := bson.D{{`IPAddressPort`, correctAddressWithPort}}
	entry := Scheme.Listener{}
	iterator := collectionListener.Find(selection).Iter()
	for iterator.Next(&entry) {
		selectionUpdate := bson.D{{`Channel`, entry.Channel}, {`Command`, entry.Command}, {`IPAddressPort`, correctAddressWithPort}}
		entry.IsActive = false
		collectionListener.Update(selectionUpdate, entry)
	}

	db.Logout()
	dbSession.Close()
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `Done shutting down now all ICCC listener for this host.`)
}

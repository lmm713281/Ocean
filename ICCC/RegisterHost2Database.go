package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"gopkg.in/mgo.v2/bson"
)

// Function to register this server to the ICCC.
func registerHost2Database() {

	// Create the host entry:
	host := Scheme.Host{}
	host.Hostname = Tools.ThisHostname()
	host.IPAddressPort = correctAddressWithPort

	// The query to find already existing entries:
	selection := bson.D{{`Hostname`, host.Hostname}, {`IPAddressPort`, host.IPAddressPort}}

	// Count the already existing entries:
	count, _ := collectionHosts.Find(selection).Count()

	// Already exist?
	if count == 1 {
		// Case: Exists!
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `This host is already registered!`, `host=`+host.Hostname, `address=`+host.IPAddressPort)
	} else {
		// Case: Not exist.
		if errInsert := collectionHosts.Insert(host); errInsert != nil {
			// Case: Was not able insert in the database
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to register this host.`, errInsert.Error(), `host=`+host.Hostname, `address=`+host.IPAddressPort)
		} else {
			// Case: Everything was fine.
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `This host is now registered.`, `host=`+host.Hostname, `address=`+host.IPAddressPort)
		}
	}
}

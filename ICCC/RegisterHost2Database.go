package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// Function to register a server to the ICCC.
func registerHost2Database(hostname, ipAddressPort string, kind byte) {

	// Create the host entry:
	host := Scheme.Host{}
	host.Hostname = hostname
	host.IPAddressPort = ipAddressPort
	host.Kind = kind

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

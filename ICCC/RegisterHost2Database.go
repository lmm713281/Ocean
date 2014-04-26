package ICCC

import "labix.org/v2/mgo/bson"
import "github.com/SommerEngineering/Ocean/ICCC/Scheme"
import "github.com/SommerEngineering/Ocean/Tools"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func registerHost2Database() {
	host := Scheme.Host{}
	host.Hostname = Tools.ThisHostname()
	host.IPAddressPort = correctAddressWithPort

	selection := bson.D{{`Hostname`, host.Hostname}, {`IPAddressPort`, host.IPAddressPort}}
	count, _ := collectionHosts.Find(selection).Count()

	if count == 1 {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `This host is already registered!`, `host=`+host.Hostname, `address=`+host.IPAddressPort)
	} else {
		if errInsert := collectionHosts.Insert(host); errInsert != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to register this host.`, errInsert.Error(), `host=`+host.Hostname, `address=`+host.IPAddressPort)
		} else {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `This host is now registered.`, `host=`+host.Hostname, `address=`+host.IPAddressPort)
		}
	}
}

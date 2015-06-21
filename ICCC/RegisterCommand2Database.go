package ICCC

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// The internal function to register a command to ICCC.
func registerCommand2Database(channel, command, ipAddressPort string, isActive bool) {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Register this ICCC command in to the database.`, `channel=`+channel, `command=`+command, `IPAddressPort=`+ipAddressPort, fmt.Sprintf("isActive=%v", isActive))

	entry := Scheme.Listener{}
	entry.Channel = channel
	entry.Command = command
	entry.IsActive = isActive
	entry.IPAddressPort = ipAddressPort

	//
	// Case: Exists?
	//
	selection := bson.D{{`Channel`, channel}, {`Command`, command}, {`IPAddressPort`, ipAddressPort}}
	count1, _ := collectionListener.Find(selection).Count()
	if count1 == 1 {
		//
		// Case: Exist but maybe not active
		//
		collectionListener.Update(selection, entry)
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Updating the existing ICCC command.`, `channel=`+channel, `command=`+command, `IPAddressPort=`+ipAddressPort)
		return
	}

	//
	// Case: Not exist
	//
	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactNone, LM.MessageNameCONFIGURATION, `This ICCC command is not known.`, `Create now a new entry!`, `channel=`+channel, `command=`+command, `IPAddressPort=`+ipAddressPort)
	if err := collectionListener.Insert(entry); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `It was not possible to add this ICCC command!`, err.Error(), `channel=`+channel, `command=`+command, `IPAddressPort=`+ipAddressPort)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `This ICCC command is now known and active.`, `channel=`+channel, `command=`+command, `IPAddressPort=`+ipAddressPort)
	}
}

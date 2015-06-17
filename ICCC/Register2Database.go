package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
)

// The internal function to register a command to ICCC.
func register2Database(channel, command string) {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Register this ICCC command in to the database.`, `channel=`+channel, `command=`+command)

	//
	// Case: Exist and active :)
	//
	emptyEntry := Scheme.Listener{}
	selection := bson.D{{`Channel`, channel}, {`Command`, command}, {`IPAddressPort`, correctAddressWithPort}, {`IsActive`, true}}
	count1, _ := collectionListener.Find(selection).Count()

	if count1 == 1 {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.SeverityHigh, LM.ImpactHigh, LM.MessageNameSTARTUP, `This ICCC command is already known and active.`, `Please shutdown the system next time!`)
		return
	}

	//
	// Case: Exist but not active
	//
	selection = bson.D{{`Channel`, channel}, {`Command`, command}, {`IPAddressPort`, correctAddressWithPort}, {`IsActive`, false}}
	notActiveEntry := Scheme.Listener{}
	collectionListener.Find(selection).One(&notActiveEntry)

	if notActiveEntry != emptyEntry {
		notActiveEntry.IsActive = true
		collectionListener.Update(selection, notActiveEntry)
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.SeverityNone, LM.ImpactNone, LM.MessageNameSTARTUP, `This ICCC command is already known but it was not active.`, `The command is active now!`)
		return
	}

	//
	// Case: Not exist
	//
	Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactNone, LM.MessageNameCONFIGURATION, `This ICCC command is not known.`, `Create now a new entry!`)

	entry := Scheme.Listener{}
	entry.Channel = channel
	entry.Command = command
	entry.IsActive = true
	entry.IPAddressPort = correctAddressWithPort

	if err := collectionListener.Insert(entry); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `It was not possible to add this ICCC command!`, err.Error(), `channel=`+channel, `command=`+command)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `This ICCC command is now known and active.`)
	}
}

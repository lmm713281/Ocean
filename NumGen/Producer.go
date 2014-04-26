package NumGen

import "time"
import "labix.org/v2/mgo/bson"
import "github.com/SommerEngineering/Ocean/Shutdown"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func producer(name string) {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The NumGen producer is now starting.`, `name=`+name)

	// Get my channel:
	myChannel := requestChannel4Name(name)

	// Read my next free number:
	currentNextFreeNumber := nextFreeNumberFromDatabase(name)

	// Where is the next "reload"?
	nextReload := currentNextFreeNumber + int64(channelBufferSize)

	// Set the next free number to the database:
	updateNextFreeNumber2Database(name, nextReload)

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `The NumGen producer is now running.`, `name=`+name)
	for nextNumber := currentNextFreeNumber; true; {
		if Shutdown.IsDown() {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSHUTDOWN, `The NumGen producer is now down.`, `name=`+name)
			return
		}

		if nextNumber > nextReload {
			nextReload = nextReload + int64(channelBufferSize)
			updateNextFreeNumber2Database(name, nextReload)

			// Enables the administrator to monitor the frequence of chunks and is able to reconfigure the settings:
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelDEBUG, LM.MessageNamePRODUCER, `The NumGen producer creates the next chunk.`, `name=`+name)
		}

		// Enqueue the next number:
		select {
		case myChannel <- nextNumber:
			nextNumber++
		case <-time.After(time.Millisecond * 500):
		}
	}
}

func nextFreeNumberFromDatabase(name string) (result int64) {
	selection := bson.D{{`Name`, name}}
	searchResult := NumberGenScheme{}

	count, _ := collectionNumGen.Find(selection).Count()
	if count == 1 {
		collectionNumGen.Find(selection).One(&searchResult)
		result = searchResult.NextFreeNumber
	} else {
		searchResult.Name = name
		searchResult.NextFreeNumber = startValue64
		collectionNumGen.Insert(searchResult)
		result = searchResult.NextFreeNumber
	}
	return
}

func updateNextFreeNumber2Database(name string, nextFreeNumber int64) {
	selection := bson.D{{`Name`, name}}
	collectionNumGen.Update(selection, NumberGenScheme{Name: name, NextFreeNumber: nextFreeNumber})
}

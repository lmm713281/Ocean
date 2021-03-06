package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
	"math"
)

// Read the latest logging events from the database.
func ReadLatest() (events []LogDBEntry, numPages int) {
	// Define the query:
	query := logDBCollection.Find(bson.D{{"Project", projectName}}).Sort(`-TimeUTC`)

	// How many record we have all over?
	numRecords := loggingViewerPageSize
	numPages = 1
	if number, err := query.Count(); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityNone, LM.ImpactNone, LM.MessageNameDATABASE, `It was not possible to find the total number of records for the latest logging reader.`, err.Error())
	} else {
		numRecords = number
		numPages = int(math.Ceil(float64(numRecords) / float64(loggingViewerPageSize)))
	}

	// Set now the page's record limit:
	query = query.Limit(loggingViewerPageSize)
	count := loggingViewerPageSize

	// Execute the query and count the results:
	if n, err := query.Count(); err == nil {
		count = n
	}

	// The iterator for the results:
	iter := query.Iter()
	entry := LogDBEntry{}
	pos := 0

	// Reserve the memory for the results:
	events = make([]LogDBEntry, count)

	// Loop over all entries and store it:
	for iter.Next(&entry) {
		// Convert the time instance to UTC:
		entry.TimeUTC = entry.TimeUTC.UTC()

		// Store it:
		events[pos] = entry
		pos++
	}

	return
}

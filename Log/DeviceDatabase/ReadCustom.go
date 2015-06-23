package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2/bson"
	"math"
	"time"
)

func ReadCustom(timeRange, logLevel, logCategory, logImpact, logSeverity, logMessageName, logSender string, logPage int) (events []LogDBEntry, numPages int) {

	// The base query:
	selection := bson.D{{"Project", projectName}}

	//
	// Build the selection statement regarding the admin's choice:
	//
	// IMPORTANT: The order of the arguments e.g. Project->TimeUTC->Sender...
	//            is very important to enable the database to use the indexes!
	//

	if timeRange != `*` {
		nowUTC := time.Now().UTC()
		switch timeRange {
		case `last5min`:
			selection = append(selection, bson.DocElem{"TimeUTC", bson.D{{"$gte", nowUTC.Add(time.Minute * -5)}}})
		case `last30min`:
			selection = append(selection, bson.DocElem{"TimeUTC", bson.D{{"$gte", nowUTC.Add(time.Minute * -30)}}})
		case `last60min`:
			selection = append(selection, bson.DocElem{"TimeUTC", bson.D{{"$gte", nowUTC.Add(time.Minute * -60)}}})
		case `last24h`:
			selection = append(selection, bson.DocElem{"TimeUTC", bson.D{{"$gte", nowUTC.Add(time.Hour * -24)}}})
		case `last7d`:
			selection = append(selection, bson.DocElem{"TimeUTC", bson.D{{"$gte", nowUTC.Add(time.Hour * -24 * 7)}}})
		case `lastMonth`:
			selection = append(selection, bson.DocElem{"TimeUTC", bson.D{{"$gte", nowUTC.Add(time.Hour * -24 * 31)}}})
		}
	}

	if logSender != `*` {
		selection = append(selection, bson.DocElem{"Sender", logSender})
	}

	if logMessageName != `*` {
		selection = append(selection, bson.DocElem{"MessageName", logMessageName})
	}

	if logLevel != `*` {
		value := `L:` + logLevel
		selection = append(selection, bson.DocElem{"Level", value})
	}

	if logCategory != `*` {
		value := `C:` + logCategory
		selection = append(selection, bson.DocElem{"Category", value})
	}

	if logImpact != `*` {
		value := `I:` + logImpact
		selection = append(selection, bson.DocElem{"Impact", value})
	}

	if logSeverity != `*` {
		value := `S:` + logSeverity
		selection = append(selection, bson.DocElem{"Severity", value})
	}

	// Build the query:
	query := logDBCollection.Find(selection)

	// How many record we have all over for this project?
	numRecords := loggingViewerPageSize
	numPages = 1
	if number, err := query.Count(); err != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityNone, LM.ImpactNone, LM.MessageNameDATABASE, `It was not possible to find the total number of records for the custom logging reader.`, err.Error())
	} else {
		numRecords = number
		numPages = int(math.Ceil(float64(numRecords) / float64(loggingViewerPageSize)))
	}

	// Sort all results:
	query = query.Sort(`-TimeUTC`)

	// Set now the page's record limit:
	query = query.Skip((logPage - 1) * loggingViewerPageSize).Limit(loggingViewerPageSize)
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

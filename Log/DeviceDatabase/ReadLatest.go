package DeviceDatabase

import (
	"gopkg.in/mgo.v2/bson"
)

// Read the latest logging events from the database.
func ReadLatest() (events []LogDBEntry) {
	// Define the query:
	query := logDBCollection.Find(bson.D{}).Sort(`-TimeUTC`).Limit(26)
	count := 26

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
		events[pos] = entry
		pos++
	}

	return
}

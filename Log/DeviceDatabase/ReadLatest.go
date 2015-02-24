package DeviceDatabase

import (
	"gopkg.in/mgo.v2/bson"
)

func ReadLatest() (events []LogDBEntry) {

	query := logDBCollection.Find(bson.D{}).Sort(`TimeUTC`).Limit(26)
	count := 26

	if n, err := query.Count(); err == nil {
		count = n
	}

	iter := query.Iter()
	entry := LogDBEntry{}
	pos := 0
	events = make([]LogDBEntry, count)

	for iter.Next(&entry) {
		events[pos] = entry
		pos++
	}

	return
}

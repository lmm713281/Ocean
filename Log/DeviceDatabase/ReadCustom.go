package DeviceDatabase

import (
	"gopkg.in/mgo.v2/bson"
)

func ReadCustom(timeRange, logLevel, logCategory, logImpact, logSeverity, logMessageName, logSender, logPage string) (events []LogDBEntry) {

	//
	// TODO => Is currently stub
	//

	query := logDBCollection.Find(bson.D{}).Sort(`-TimeUTC`).Limit(26)
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

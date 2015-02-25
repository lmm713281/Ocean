package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

func write2Cache(entries []Meta.Entry) {
	for _, entry := range entries {
		if len(cache) == cacheSizeNumberOfEvents {
			go cacheFull()
		}

		logDBentry := LogDBEntry{}
		logDBentry.Category = entry.Category.Format()
		logDBentry.Impact = entry.Impact.Format()
		logDBentry.Level = entry.Level.Format()
		logDBentry.MessageDescription = entry.MessageDescription
		logDBentry.MessageName = string(entry.MessageName)
		logDBentry.Parameters = entry.Parameters
		logDBentry.Project = entry.Project
		logDBentry.Sender = string(entry.Sender)
		logDBentry.Severity = entry.Severity.Format()
		logDBentry.TimeUTC = entry.Time
		cache <- logDBentry
	}
}

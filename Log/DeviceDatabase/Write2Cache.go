package DeviceDatabase

import "github.com/SommerEngineering/Ocean/Log/Meta"

func write2Cache(entries []Meta.Entry) {
	for _, entry := range entries {
		if len(cache) == cacheSizeNumberOfEvents {
			go cacheFull()
		}

		logDBentry := LogDBEntry{}
		logDBentry.Category = Meta.FormatCategory(entry.Category)
		logDBentry.Impact = Meta.FormatImpact(entry.Impact)
		logDBentry.Level = Meta.FormatLevel(entry.Level)
		logDBentry.MessageDescription = entry.MessageDescription
		logDBentry.MessageName = string(entry.MessageName)
		logDBentry.Parameters = entry.Parameters
		logDBentry.Project = entry.Project
		logDBentry.Sender = string(entry.Sender)
		logDBentry.Severity = Meta.FormatSeverity(entry.Severity)
		logDBentry.TimeUTC = entry.Time
		cache <- logDBentry
	}
}

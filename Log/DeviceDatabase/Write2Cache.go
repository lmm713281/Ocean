package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

// This function writes a batch of log entries to the cache.
func write2Cache(entries []Meta.Entry) {
	// Loop over each entry:
	for _, entry := range entries {
		// If the cache is full, execute it:
		if len(cache) == cacheSizeNumberOfEvents {
			// Execute the cache with a new thread:
			go cacheFull()
		}

		// Convert the log entry to the database format:
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

		// Write it to the cache:
		cache <- logDBentry
	}
}

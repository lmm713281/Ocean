package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

// Function to format a logging database entry as string.
func (entry LogDBEntry) Format() (result string) {
	// First, we convert the logging db entry to the common logging type:
	converted := Meta.Entry{}
	converted.Time = entry.TimeUTC
	converted.Project = entry.Project
	converted.Sender = Meta.Sender(entry.Sender)
	converted.Category = Meta.ParseCategory(entry.Category)
	converted.Level = Meta.ParseLevel(entry.Level)
	converted.Severity = Meta.ParseSeverity(entry.Severity)
	converted.Impact = Meta.ParseImpact(entry.Impact)
	converted.MessageName = Meta.MessageName(entry.MessageName)
	converted.MessageDescription = entry.MessageDescription
	converted.Parameters = entry.Parameters

	// Second, we can use then the format operation of these type:
	result = converted.Format()
	return
}

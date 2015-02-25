package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

func (entry LogDBEntry) Format() (result string) {

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

	result = converted.Format()
	return
}

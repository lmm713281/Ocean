package Web

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	"strings"
)

// Read a custom event range from the database.
func readCustom(timeRange, logLevel, logCategory, logImpact, logSeverity, logMessageName, logSender, logPage string) (events []Scheme.LogEvent) {

	// Get the custom events:
	eventsFromDB := DeviceDatabase.ReadCustom(timeRange, logLevel, logCategory, logImpact, logSeverity, logMessageName, logSender, logPage)
	count := len(eventsFromDB)

	// Array with all events, prepared for the website:
	events = make([]Scheme.LogEvent, count)

	// Copy each event to the right format:
	for n := 0; n < count; n++ {
		eventFromDB := eventsFromDB[n]
		events[n] = Scheme.LogEvent{}
		events[n].LogLine = eventFromDB.Format()
		events[n].LogLevel = fmt.Sprintf("log%s", strings.ToLower(eventFromDB.Level[2:]))

		// Vary the color of each line:
		if n%2 == 0 {
			events[n].AB = Scheme.B
		} else {
			events[n].AB = Scheme.A
		}
	}

	return
}

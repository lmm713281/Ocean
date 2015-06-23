package Web

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	"strings"
)

// Read a custom event range from the database.
func readCustom(timeRange, logLevel, logCategory, logImpact, logSeverity, logMessageName, logSender string, logPage int) (events []Scheme.LogEvent, numPages int) {

	// Get the custom events:
	eventsFromDB, totalNumberPages := DeviceDatabase.ReadCustom(timeRange, logLevel, logCategory, logImpact, logSeverity, logMessageName, logSender, logPage)
	count := len(eventsFromDB)

	// Array with all events, prepared for the website:
	events = make([]Scheme.LogEvent, count)

	// Copy each event to the right format:
	for n := 0; n < count; n++ {
		eventFromDB := eventsFromDB[n]
		events[n] = Scheme.LogEvent{}
		events[n].LogLine = eventFromDB.Format()

		// Transfer the log level:
		if len(eventFromDB.Level) > 2 {
			events[n].LogLevel = fmt.Sprintf("log%s", strings.ToLower(eventFromDB.Level[2:]))
		} else {
			events[n].LogLevel = fmt.Sprintf("log%s", strings.ToLower(eventFromDB.Level))
		}

		// Vary the color of each line:
		if n%2 == 0 {
			events[n].AB = Scheme.B
		} else {
			events[n].AB = Scheme.A
		}
	}

	numPages = totalNumberPages
	return
}

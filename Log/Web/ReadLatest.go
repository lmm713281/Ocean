package Web

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	"strings"
)

// Read the latest log events from the database
func readLatest() (events []Scheme.LogEvent, numPages int) {
	// Get the latest events from the database
	eventsFromDB, totalNumberPages := DeviceDatabase.ReadLatest()
	count := len(eventsFromDB)

	// Array for the log events, prepared for the website:
	events = make([]Scheme.LogEvent, count)

	// Copy each event to the right format for the website:
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

	numPages = totalNumberPages
	return
}

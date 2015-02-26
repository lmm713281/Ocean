package Web

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	"github.com/SommerEngineering/Ocean/Log/Web/Scheme"
	"strings"
)

func readLatest() (events []Scheme.LogEvent) {

	eventsFromDB := DeviceDatabase.ReadLatest()
	count := len(eventsFromDB)
	events = make([]Scheme.LogEvent, count)

	for n := 0; n < count; n++ {
		eventFromDB := eventsFromDB[n]
		events[n] = Scheme.LogEvent{}
		events[n].LogLine = eventFromDB.Format()
		events[n].LogLevel = fmt.Sprintf("log%s", strings.ToLower(eventFromDB.Level[2:]))

		if n%2 == 0 {
			events[n].AB = Scheme.B
		} else {
			events[n].AB = Scheme.A
		}
	}

	return
}

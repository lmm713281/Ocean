package Web

import (
	"github.com/SommerEngineering/Ocean/Log/DeviceDatabase"
	"github.com/SommerEngineering/Ocean/Log/Web/Scheme"
)

func readLatest() (events []Scheme.LogEvent) {

	eventsFromDB := DeviceDatabase.ReadLatest()
	count := len(eventsFromDB)
	events = make([]Scheme.LogEvent, count)

	for n := 0; n < count; n++ {
		eventFromDB := eventsFromDB[n]
		events[n] = Scheme.LogEvent{}
		//events[n].LogLine = eventFromDB.

		if n%2 == 0 {
			events[n].AB = Scheme.B
		} else {
			events[n].AB = Scheme.A
		}
	}

	// data.Events = make([]Scheme.LogEvent, 3)
	// data.Events[0].AB = Scheme.A
	// data.Events[0].LogLevel = Scheme.LogINFO
	// data.Events[0].LogLine = `hello world`
	return
}

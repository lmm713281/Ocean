package Log

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"time"
)

/*
The scheduler function which runs at a own thread.
Pleae note: The scheduler is the consumer for the logging channel.
*/
func scheduler(logBuffer chan Meta.Entry) {

	LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, Meta.MessageNameSTARTUP, `The scheduler runs now.`)
	var stopNextTime = false

	// Endless loop:
	for {

		// Enable the loop to stop:
		if stopNextTime {
			break
		}

		// Read one entry from the buffer (channel):
		nextEntry, ok := <-logBuffer

		// Case: The channel was closed.
		if !ok {

			// Create a log message for this event.
			stopNextTime = true
			nextEntry = Meta.Entry{}
			nextEntry.Project = projectName
			nextEntry.Time = time.Now().UTC()
			nextEntry.Sender = senderName
			nextEntry.Category = Meta.CategorySYSTEM
			nextEntry.Level = Meta.LevelWARN
			nextEntry.Severity = Meta.SeverityCritical
			nextEntry.Impact = Meta.ImpactNone
			nextEntry.MessageName = Meta.MessageNameCOMMUNICATION
			nextEntry.MessageDescription = `The logging channel was closed!`
		}

		// Queue the log event for the delivery to the devices:
		deviceDelay(nextEntry)
	}

	// Exit the scheduler. Send the signal.
	LogFull(senderName, Meta.CategorySYSTEM, Meta.LevelWARN, Meta.SeverityCritical, Meta.ImpactNone, Meta.MessageNameSHUTDOWN, `The scheduler is down now.`)
	schedulerExitSignal <- true
}

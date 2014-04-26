package Log

import "github.com/SommerEngineering/Ocean/Log/Meta"
import "time"

// Note: The scheduler is the consumer for the logging channel!
func scheduler(logBuffer chan Meta.Entry) {

	LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, Meta.MessageNameSTARTUP, `The scheduler runs now.`)
	var stopNextTime = false

	for {
		if stopNextTime {
			break
		}

		nextEntry, ok := <-logBuffer

		if !ok {
			// The channel was closed!
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

		deviceDelay(nextEntry)
	}

	LogFull(senderName, Meta.CategorySYSTEM, Meta.LevelWARN, Meta.SeverityCritical, Meta.ImpactNone, Meta.MessageNameSHUTDOWN, `The scheduler is down now.`)
}

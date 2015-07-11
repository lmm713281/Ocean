package ICCCLog

import (
	"github.com/SommerEngineering/Ocean/ICCC"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"time"
)

// The receiver function for the ICCC new log event message.
func ICCCNewLogEventReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC new log event message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := ICCC.Data2Message(SystemMessages.ICCCNewLogEvent{}, data)

	// Was it possible to convert the data?
	if obj != nil {
		// Cast the object to the right type:
		messageData := obj.(SystemMessages.ICCCNewLogEvent)

		// Provide a log entry:
		logEntry := LM.Entry{}
		logEntry.Sender = LM.Sender(messageData.Sender)
		logEntry.Category = LM.ParseCategory(messageData.Category)
		logEntry.Level = LM.ParseLevel(messageData.Level)
		logEntry.Severity = LM.ParseSeverity(messageData.Severity)
		logEntry.Impact = LM.ParseImpact(messageData.Impact)
		logEntry.MessageName = LM.MessageName(messageData.MessageName)
		logEntry.MessageDescription = messageData.MessageDescription
		logEntry.Parameters = messageData.Parameters
		logEntry.Time = time.Unix(messageData.UnixTimestampUTC, 0).UTC()

		// Deliver the log event:
		Log.TakeEntry(logEntry)

		// An answer is necessary:
		return ICCC.Message2Data("", "", SystemMessages.AnswerACK)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

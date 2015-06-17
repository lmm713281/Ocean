package Log

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"strings"
	"time"
)

// Writes a log message to the channel.
func writeToChannel(logEntry Meta.Entry) {
	select {
	case entriesBuffer <- logEntry:
	case <-time.After(time.Duration(int64(logBufferTimeoutSeconds)) * time.Second):
		// Warn: Cannot log here to prevent endless loop and memory leak!
		fmt.Println(`Warning: Was not able to write to the logging buffer! Message=` + logEntry.Format())
	}
}

/*
If, for any reason, you want to deliver a whole log entry to the logging system, then use this function to do so :) Normally,
just use the LogFull() and LogShort() functions instead ;) There is one reason to use this: If you have to deliver old log
events you have may imported or you have received a batch of log events of an remote system, etc. Because the LogShort() and
LogFull() methods are setting the time to the current UTC time. Therefore, TakeEntry() is the only way to provide the time!
*/
func TakeEntry(logEntry Meta.Entry) {

	logEntry = clearEntry(logEntry)
	mutexChannel.RLock()
	defer mutexChannel.RUnlock()

	if !channelReady {

		mutexPreChannelBuffer.Lock()
		defer mutexPreChannelBuffer.Unlock()

		preChannelBuffer.PushBack(logEntry)
		return
	}

	if !preChannelBufferUsed {
		preChannelBufferUsed = true
		mutexPreChannelBuffer.Lock()
		for entry := preChannelBuffer.Front(); entry != nil; entry = entry.Next() {
			writeToChannel(entry.Value.(Meta.Entry))
		}
		preChannelBuffer.Init()
		mutexPreChannelBuffer.Unlock()
	}

	writeToChannel(logEntry)
}

/*
Create and deliver a full log message with all fields and with the current UTC time as logging time.

	Sender		The sender of this message (name of your component, like e.g. "APP::LOGIC::PAYPAL" etc.)
	Category	Use "CategoryBUSINESS" for business events (new payments, etc.), use "CategoryUSER" for log events
			regarding users (e.g. new user registered, etc.) or "CategoryAPP" (for anything else).
	Level 		Choose a logging level.
	Severity	Choose a degree of severity.
	Impact		Choose a degree of impact.
	MessageName	Choose a message name. This is like a common category for this event!
	Description	The logging message you want to deliver.
	Parameters	Provide as many additional parameters (type string) as you need.
*/
func LogFull(sender Meta.Sender, category Meta.Category, level Meta.Level, severity Meta.Severity, impact Meta.Impact, messageName Meta.MessageName, messageDescription string, parameters ...string) {

	entry := Meta.Entry{}
	entry.Project = projectName
	entry.Time = time.Now().UTC()
	entry.Category = category
	entry.Level = level
	entry.MessageDescription = messageDescription
	entry.MessageName = messageName
	entry.Parameters = parameters
	entry.Severity = severity
	entry.Impact = impact
	entry.Sender = sender

	TakeEntry(entry)
}

/*
Create and deliver a short log message with the current UTC time as logging time. The fields severity and impact
are both set to "none" value. Therefore, this short logging function does not fit for logging problems, errors etc.

	Sender		The sender of this message (name of your component, like e.g. "APP::LOGIC::PAYPAL" etc.)
	Category	Use "CategoryBUSINESS" for business events (new payments, etc.), use "CategoryUSER" for log events
			regarding users (e.g. new user registered, etc.) or "CategoryAPP" (for anything else).
	Level 		Choose a logging level.
	MessageName	Choose a message name. This is like a common category for this event!
	Description	The logging message you want to deliver.
	Parameters	Provide as many additional parameters (type string) as you need.
*/
func LogShort(sender Meta.Sender, category Meta.Category, level Meta.Level, messageName Meta.MessageName, messageDescription string, parameters ...string) {

	entry := Meta.Entry{}
	entry.Project = projectName
	entry.Time = time.Now().UTC()
	entry.Category = category
	entry.Level = level
	entry.MessageDescription = messageDescription
	entry.MessageName = messageName
	entry.Parameters = parameters
	entry.Severity = Meta.SeverityNone
	entry.Impact = Meta.ImpactNone
	entry.Sender = sender

	TakeEntry(entry)
}

// Removes white spaces from the message.
func clearEntry(entry Meta.Entry) (result Meta.Entry) {
	entry.MessageDescription = removeWhitespaces(entry.MessageDescription)
	entry.Parameters = clearParameters(entry.Parameters)
	result = entry
	return
}

// Remove white spaces from the parameters.
func clearParameters(oldParameters []string) (result []string) {
	for n := 0; n < len(oldParameters); n++ {
		oldParameters[n] = removeWhitespaces(oldParameters[n])
	}

	result = oldParameters
	return
}

// Removes white spaces from a string.
func removeWhitespaces(text string) (result string) {
	text = strings.Replace(text, "\n", ` `, -1)
	text = strings.Replace(text, "\t", ` `, -1)
	text = strings.Replace(text, "\r", ``, -1)

	result = text
	return
}

package Meta

import (
	"time"
)

// Type for a log entry.
type Entry struct {
	Project            string
	Time               time.Time
	Sender             Sender
	Category           Category
	Level              Level
	Severity           Severity
	Impact             Impact
	MessageName        MessageName
	MessageDescription string
	Parameters         []string
}

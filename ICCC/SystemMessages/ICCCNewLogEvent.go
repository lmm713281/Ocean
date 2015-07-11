package SystemMessages

// The message to send logging events over ICCC.
type ICCCNewLogEvent struct {
	Project            string
	UnixTimestampUTC   int64
	Sender             string
	Category           string
	Level              string
	Severity           string
	Impact             string
	MessageName        string
	MessageDescription string
	Parameters         []string
}

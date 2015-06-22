package Scheme

// The type for the web logger viewer template
type Viewer struct {
	Title              string
	SetLiveView        bool
	CurrentLevel       string
	CurrentTimeRange   string
	CurrentCategory    string
	CurrentImpact      string
	CurrentSeverity    string
	CurrentMessageName string
	CurrentSender      string
	CurrentPage        string
	MessageNames       []string
	Sender             []string
	Events             []LogEvent
}

// Type for a log event
type LogEvent struct {
	LogLine  string
	LogLevel string // logwarn || logdebug || logerror || loginfo || logtalkative || logsecurity
	AB       string // loga || logb
}

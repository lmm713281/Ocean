package Scheme

// Type for a log event
type LogEvent struct {
	LogLine  string
	LogLevel string // logwarn || logdebug || logerror || loginfo || logtalkative || logsecurity
	AB       string // loga || logb
}

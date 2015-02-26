package Scheme

type Viewer struct {
	Title        string
	SetLiveView  bool
	MessageNames []string
	Sender       []string
	Events       []LogEvent
}

// <li class="logline loga logwarn">
//	 <div>....</div>
// </li>
// <li class="logline logb logwarn">
// 	<div>....</div>
// </li>
type LogEvent struct {
	LogLine  string
	LogLevel string // logwarn || logdebug || logerror || loginfo || logtalkative || logsecurity
	AB       string // loga || logb
}

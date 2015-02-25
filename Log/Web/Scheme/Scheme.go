package Scheme

type Viewer struct {
	Title        string
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
	LogLevel string // L:DEBUG || L:ERROR || L:INFO || L:SECURITY || L:TALKATIVE || L:WARN
	AB       string // loga || logb
}

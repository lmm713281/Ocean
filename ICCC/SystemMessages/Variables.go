package SystemMessages

var (
	AnswerACK     ICCCDefaultAnswer = ICCCDefaultAnswer{true, 0}   // The command was successful
	AnswerNACK    ICCCDefaultAnswer = ICCCDefaultAnswer{false, 0}  // The command was not successful
	AnswerUNKNOWN ICCCDefaultAnswer = ICCCDefaultAnswer{false, -1} // The answer is unknown e.g. an error while reading the answer (HTTP errors, etc.)
)

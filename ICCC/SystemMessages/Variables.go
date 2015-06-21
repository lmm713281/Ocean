package SystemMessages

var (
	AnswerACK     DefaultAnswer = DefaultAnswer{true, 0}   // The command was successful
	AnswerNACK    DefaultAnswer = DefaultAnswer{false, 0}  // The command was not successful
	AnswerUNKNOWN DefaultAnswer = DefaultAnswer{false, -1} // The answer is unknown e.g. an error while reading the answer (HTTP errors, etc.)
)

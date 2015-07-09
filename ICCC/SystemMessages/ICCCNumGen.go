package SystemMessages

// The message to request a new number from NumGen package.
type ICCCNumGenNext struct {
}

// The response to the NumGen request.
type ICCCAnswerNumGen struct {
	Number int64
}

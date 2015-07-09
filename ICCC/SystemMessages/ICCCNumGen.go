package SystemMessages

// The message to request a new number from NumGen package.
type ICCCNumGenNext struct {
}

// The response to the NumGen request.
type ICCCNumGenNextAnswer struct {
	Number int64
}

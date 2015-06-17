package Log

/*
A function to change the state of the logging after the database is
accessible.
*/
func LoggingIsReady() {
	channelReady = true
	preChannelBufferUsed = false
}

package Log

/*
This function is used just internal by Ocean. Please do not call this function by your self!
*/
func LoggingIsReady() {
	channelReady = true
	preChannelBufferUsed = false
}

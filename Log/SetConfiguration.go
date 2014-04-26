package Log

/*
This function is used just internal by Ocean to change some configuration afterwards, after the first runtime stage
(transition from not configured state in to the desired configurated stage, after the configuration database is ready).
Please do not call this function by your self!
*/
func SetBufferSize(bufferSize int) {
	logBufferSize = bufferSize
}

/*
This function is used just internal by Ocean to change some configuration afterwards, after the first runtime stage
(transition from not configured state in to the desired configurated stage, after the configuration database is ready).
Please do not call this function by your self!
*/
func SetTimeoutSeconds(timeoutSeconds int) {
	logBufferTimeoutSeconds = timeoutSeconds
}

/*
This function is used just internal by Ocean to change some configuration afterwards, after the first runtime stage
(transition from not configured state in to the desired configurated stage, after the configuration database is ready).
Please do not call this function by your self!
*/
func SetDeviceDelayNumberEvents(numberEvents int) {
	logDeviceDelayNumberEvents = numberEvents
}

/*
This function is used just internal by Ocean to change some configuration afterwards, after the first runtime stage
(transition from not configured state in to the desired configurated stage, after the configuration database is ready).
Please do not call this function by your self!
*/
func SetDeviceDelayTimeoutSeconds(seconds int) {
	logDeviceDelayTimeoutSeconds = seconds
}

/*
This function is used just internal by Ocean to change some configuration afterwards, after the first runtime stage
(transition from not configured state in to the desired configurated stage, after the configuration database is ready).
Please do not call this function by your self!
*/
func SetProjectName(project string) {
	projectName = project
}

/*
This function is used just internal by Ocean to change some configuration afterwards, after the first runtime stage
(transition from not configured state in to the desired configurated stage, after the configuration database is ready).
Please do not call this function by your self!
*/
func ApplyConfigurationChanges() {

	mutexChannel.Lock()
	channelReady = false
	close(entriesBuffer)
	mutexChannel.Unlock()

	initCode()
}

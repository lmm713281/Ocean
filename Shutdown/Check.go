package Shutdown

// A function to check if the system goes down right now.
func IsDown() (result bool) {
	result = stopAllRequests
	return
}

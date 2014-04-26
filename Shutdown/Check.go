package Shutdown

func IsDown() (result bool) {
	result = stopAllRequests
	return
}

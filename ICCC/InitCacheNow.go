package ICCC

// Starts the timer cache once and exit it after (no thread, no endless loop).
func InitCacheNow() {
	startCacheTimerLock.Lock()
	defer startCacheTimerLock.Unlock()

	if cacheTimerRunning {
		return
	}

	cacheTimerLogic(false)
}

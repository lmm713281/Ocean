package ICCC

// Setup and starts the cache timer.
func initCacheTimer() {
	startCacheTimerLock.Lock()
	defer startCacheTimerLock.Unlock()

	if cacheTimerRunning {
		return
	} else {
		cacheTimerRunning = true
	}

	// Start another thread with the timer logic:
	go func() {
		// Endless loop:
		for {
			cacheTimerLogic(true)
		}
	}()
}

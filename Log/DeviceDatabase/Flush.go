package DeviceDatabase

// Flush the cache and write all messages to the database.
func (dev Database) Flush() {
	mutexCacheFull.Lock()
	defer mutexCacheFull.Unlock()

	amount := len(cache)
	for counter := 0; counter < amount; counter++ {
		write2Database(<-cache)
	}

	// Shutdown the database connection:
	logDB.Logout()
	logDBSession.Close()
}

package DeviceDatabase

func (dev Database) Flush() {
	mutexCacheFull.Lock()
	defer mutexCacheFull.Unlock()

	amount := len(cache)
	for counter := 0; counter < amount; counter++ {
		write2Database(<-cache)
	}

	logDB.Logout()
	logDBSession.Close()
}

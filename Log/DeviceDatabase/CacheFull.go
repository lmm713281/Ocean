package DeviceDatabase

// Function to check if the cache is full. If so, write all events to the database.
func cacheFull() {
	mutexCacheFull.Lock()
	defer mutexCacheFull.Unlock()

	// Is the cache full?
	if len(cache) < cacheSizeNumberOfEvents {
		// Case: Cache is not full.
		return
	}

	// Case: The cache is full. Write all events to the database.
	for counter := 0; counter < cacheSizeNumberOfEvents; counter++ {
		write2Database(<-cache)
	}
}

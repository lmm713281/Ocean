package DeviceDatabase

// Read the message names out of the cache.
func ReadMessageNames() (messageNames []string) {
	mutexCacheMessageNames.RLock()
	defer mutexCacheMessageNames.RUnlock()
	messageNames = cacheMessageNames
	return
}

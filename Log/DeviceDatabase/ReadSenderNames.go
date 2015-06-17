package DeviceDatabase

// Read the sender names out of the cache.
func ReadSenderNames() (senderNames []string) {
	mutexCacheSenderNames.RLock()
	defer mutexCacheSenderNames.RUnlock()
	senderNames = cacheSenderNames
	return
}

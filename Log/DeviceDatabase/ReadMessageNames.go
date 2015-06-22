package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
)

// Read the message names out of the cache.
func ReadMessageNames() (messageNames []Scheme.MessageNames) {
	mutexCacheMessageNames.RLock()
	defer mutexCacheMessageNames.RUnlock()

	// Transform the values to the right format:
	for _, entry := range cacheMessageNames {
		messageNames = append(messageNames, Scheme.MessageNames(entry))
	}
	return
}

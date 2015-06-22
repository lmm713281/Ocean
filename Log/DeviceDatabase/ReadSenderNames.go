package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Admin/Scheme"
)

// Read the sender names out of the cache.
func ReadSenderNames() (senderNames []Scheme.Sender) {
	mutexCacheSenderNames.RLock()
	defer mutexCacheSenderNames.RUnlock()

	// Transform the values to the right format:
	for _, entry := range cacheSenderNames {
		senderNames = append(senderNames, Scheme.Sender(entry))
	}
	return
}

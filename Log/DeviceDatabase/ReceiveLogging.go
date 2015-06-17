package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

// This function is the interface between the logging system and the database logger.
func (dev Database) Log(entries []Meta.Entry) {

	//
	// Cannot log here to prevent endless loop (consumer is also producer)
	//

	// Write every incoming batch to the cache:
	write2Cache(entries)
}

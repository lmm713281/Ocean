package DeviceDatabase

import (
	"github.com/SommerEngineering/Ocean/Log/Meta"
)

type Database struct {
}

func (dev Database) Log(entries []Meta.Entry) {

	//
	// Cannot log here to prevent endless loop (consumer is also producer)
	//

	write2Cache(entries)
}

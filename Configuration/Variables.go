package Configuration

import (
	"github.com/SommerEngineering/Ocean/Configuration/Meta"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

var (
	filename                         = "configuration.json" // Where is the configuration located?
	configuration Meta.Configuration = Meta.Configuration{} // The loaded configuration
	isInit                           = false                // Is the configuration loaded?
	senderName    LM.Sender          = `System::Configuration`
)

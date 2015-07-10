package Admin

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
)

// Data for the admin's configuration management site
type AdminWebConfiguration struct {
	Configuration []ConfigurationDB.ConfigurationDBEntry
}

// Data for the admin's overview i.e. dashboard
type AdminWebOverview struct {
	Version string
}

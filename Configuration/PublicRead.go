package Configuration

import "github.com/SommerEngineering/Ocean/Configuration/Meta"

/*
Read the whole configuration and enable Ocean to get the configuration database.

Hint: Normally, you do not use this package at all, because the application configuration should persist
inside the configuration database.
*/
func Read() (result Meta.Configuration) {
	result = configuration
	return
}

/*
This package reads the configuration file from the disk (file: configuration.json).
This configuratin is just used to specific the configuration database to go further.

An example configuration.json file:

	{
		"ConfigDBHostname" : "localhost:27017",
		"ConfigDBDatabase" : "MyDatabase",
		"ConfigDBConfigurationCollection" : "Configuration",
		"ConfigDBConfigurationCollectionUsername" : "ConfigurationUsername",
		"ConfigDBConfigurationCollectionPassword" : "ConfigurationPassword"
	}

Hint #1: Ocean is using MongoDB as database ;-)
Hint #2: Normally, you do not use this package at all, because the application configuration should persist
inside the configuration database. Your task: Provide a configuration.json file at the installation directory ;-)
*/
package Configuration

/*
This package provides the configuration database access: Ocean uses MongoDB as database! The configuration is represented as
name value pairs. Both, the name and the value, are strings. If you need numbers for the configuration, you have to convert
these strings in to numbers afterwards. Use the Read() function to read a specific configuration value.

To provide own application configurations, use the function CheckSingleConfigurationPresentsAndAddIfMissing() to ensure, that
at least a default value is present in the database. Collect all these function calls inside a init() function somewhere at your
application.
*/
package ConfigurationDB

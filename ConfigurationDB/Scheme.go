package ConfigurationDB

// The type for a configuration entry.
type ConfigurationDBEntry struct {
	Name  string `bson:"Name"`
	Value string `bson:"Value"`
}

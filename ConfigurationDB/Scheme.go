package ConfigurationDB

type ConfigurationDBEntry struct {
	Name  string `bson:"Name"`
	Value string `bson:"Value"`
}

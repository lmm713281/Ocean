package NumGen

type NumberGenScheme struct {
	Name           string `bson:"Name"`
	NextFreeNumber int64  `bson:"NextFreeNumber"`
}

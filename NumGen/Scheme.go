package NumGen

// The scheme for the database.
type NumberGenScheme struct {
	Name           string `bson:"Name"`           // A name for this counter.
	NextFreeNumber int64  `bson:"NextFreeNumber"` // The next number.
}

package SystemMessages

// The type for any answer, which can be extended by using CommandData.
type DefaultAnswer struct {
	CommandSuccessful bool
	CommandAnswer     int64
}

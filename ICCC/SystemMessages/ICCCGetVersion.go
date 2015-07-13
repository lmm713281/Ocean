package SystemMessages

// Requests the version from a Ocean server
type ICCCGetVersion struct {
}

// Answer to the version request
type ICCCGetVersionAnswer struct {
	Kind    byte   // Ocean || Component
	Name    string // Ocean: Hostname; Components: Name
	Version string // The current version
}

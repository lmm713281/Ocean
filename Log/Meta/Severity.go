package Meta

type Severity byte

const (
	SeverityNone     = Severity(iota)
	SeverityLow      = Severity(iota)
	SeverityMiddle   = Severity(iota)
	SeverityHigh     = Severity(iota)
	SeverityCritical = Severity(iota)
	SeverityUnknown  = Severity(iota)
)

func FormatSeverity(pri Severity) (result string) {
	switch pri {
	case SeverityCritical:
		result = `S:CRITICAL`
	case SeverityHigh:
		result = `S:HIGH`
	case SeverityLow:
		result = `S:LOW`
	case SeverityMiddle:
		result = `S:MIDDLE`
	case SeverityNone:
		result = `S:NONE`
	case SeverityUnknown:
		result = `S:UNKNOWN`
	default:
		result = `S:N/A`
	}

	return
}

package Meta

type Severity byte

const (
	SeverityNone     = Severity(iota) // None severity
	SeverityLow      = Severity(iota) // Low severity
	SeverityMiddle   = Severity(iota) // Middle severity
	SeverityHigh     = Severity(iota) // High severity
	SeverityCritical = Severity(iota) // Critical severity
	SeverityUnknown  = Severity(iota) // Unknown severity
)

// Format the severity as string.
func (pri Severity) Format() (result string) {
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

// Parse the severity from a string.
func ParseSeverity(pri string) (value Severity) {
	switch pri {
	case `S:CRITICAL`:
		value = SeverityCritical
	case `S:HIGH`:
		value = SeverityHigh
	case `S:LOW`:
		value = SeverityLow
	case `S:MIDDLE`:
		value = SeverityMiddle
	case `S:NONE`:
		value = SeverityNone
	case `S:UNKNOWN`:
		value = SeverityUnknown
	default:
		value = SeverityUnknown
	}

	return
}

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

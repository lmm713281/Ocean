package Meta

type Impact byte

const (
	ImpactNone     = Impact(iota) // None impact
	ImpactLow      = Impact(iota) // Low impact
	ImpactMiddle   = Impact(iota) // Middle impact
	ImpactHigh     = Impact(iota) // High impact
	ImpactCritical = Impact(iota) // Critical impact
	ImpactUnknown  = Impact(iota) // Unknown impact
)

// Formats a impact as string.
func (pri Impact) Format() (result string) {
	switch pri {
	case ImpactCritical:
		result = `I:CRITICAL`
	case ImpactHigh:
		result = `I:HIGH`
	case ImpactLow:
		result = `I:LOW`
	case ImpactMiddle:
		result = `I:MIDDLE`
	case ImpactNone:
		result = `I:NONE`
	case ImpactUnknown:
		result = `I:UNKNOWN`
	default:
		result = `I:N/A`
	}

	return
}

// Parse a impact from a string.
func ParseImpact(pri string) (value Impact) {
	switch pri {
	case `I:CRITICAL`:
		value = ImpactCritical
	case `I:HIGH`:
		value = ImpactHigh
	case `I:LOW`:
		value = ImpactLow
	case `I:MIDDLE`:
		value = ImpactMiddle
	case `I:NONE`:
		value = ImpactNone
	case `I:UNKNOWN`:
		value = ImpactUnknown
	default:
		value = ImpactUnknown
	}

	return
}

package Meta

type Impact byte

const (
	ImpactNone     = Impact(iota)
	ImpactLow      = Impact(iota)
	ImpactMiddle   = Impact(iota)
	ImpactHigh     = Impact(iota)
	ImpactCritical = Impact(iota)
	ImpactUnknown  = Impact(iota)
)

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

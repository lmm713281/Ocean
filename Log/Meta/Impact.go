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

func FormatImpact(pri Impact) (result string) {
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

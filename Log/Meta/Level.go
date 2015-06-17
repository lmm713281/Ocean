package Meta

type Level byte

const (
	LevelWARN      = Level(iota) // Level: Warning
	LevelDEBUG     = Level(iota) // Level: Debug
	LevelERROR     = Level(iota) // Level: Error
	LevelINFO      = Level(iota) // Level: Information
	LevelTALKATIVE = Level(iota) // Level: Talkative (even more events as debug)
	LevelSECURITY  = Level(iota) // Level: Security
)

// Formats a level as string.
func (lvl Level) Format() (result string) {
	switch lvl {
	case LevelDEBUG:
		result = `L:DEBUG`
	case LevelERROR:
		result = `L:ERROR`
	case LevelINFO:
		result = `L:INFO`
	case LevelSECURITY:
		result = `L:SECURITY`
	case LevelTALKATIVE:
		result = `L:TALKATIVE`
	case LevelWARN:
		result = `L:WARN`
	default:
		result = `L:N/A`
	}

	return
}

// Parse a level from a string.
func ParseLevel(lvl string) (value Level) {
	switch lvl {
	case `L:DEBUG`:
		value = LevelDEBUG
	case `L:ERROR`:
		value = LevelERROR
	case `L:INFO`:
		value = LevelINFO
	case `L:SECURITY`:
		value = LevelSECURITY
	case `L:TALKATIVE`:
		value = LevelTALKATIVE
	case `L:WARN`:
		value = LevelWARN
	default:
		value = LevelERROR
	}

	return
}

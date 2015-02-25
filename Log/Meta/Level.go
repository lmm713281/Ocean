package Meta

type Level byte

const (
	LevelWARN      = Level(iota)
	LevelDEBUG     = Level(iota)
	LevelERROR     = Level(iota)
	LevelINFO      = Level(iota)
	LevelTALKATIVE = Level(iota)
	LevelSECURITY  = Level(iota)
)

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

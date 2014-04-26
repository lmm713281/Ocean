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

func FormatLevel(lvl Level) (result string) {
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

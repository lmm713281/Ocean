package Tools

type Language struct {
	Language string
	Factor   float32
}

type Languages []Language

func (lang Languages) Len() int {
	return len(lang)
}

func (lang Languages) Less(i, j int) bool {
	return lang[i].Factor > lang[j].Factor
}

func (lang Languages) Swap(i, j int) {
	lang[i], lang[j] = lang[j], lang[i]
}

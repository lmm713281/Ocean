package Tools

// The data type for the language.
type Language struct {
	Language string
	Factor   float32
}

// The data type for multiple languages.
type Languages []Language

// Ensure, that the languages are sortable.
func (lang Languages) Len() int {
	return len(lang)
}

// Ensure, that the languages are sortable.
func (lang Languages) Less(i, j int) bool {
	return lang[i].Factor > lang[j].Factor
}

// Ensure, that the languages are sortable.
func (lang Languages) Swap(i, j int) {
	lang[i], lang[j] = lang[j], lang[i]
}

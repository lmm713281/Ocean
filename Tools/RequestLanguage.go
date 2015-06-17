package Tools

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Gets the language of a request for enable I18N.
func GetRequestLanguage(request *http.Request) (resultLangs Languages) {

	// Read the goven data from the request:
	languageCodeForm := request.FormValue(`lang`)
	languageCodeHeader := request.Header.Get(`Accept-Language`)

	// Is only one language available?
	if languageCodeForm != `` && ((len(languageCodeForm) == 2) || (len(languageCodeForm) == 5 && strings.Contains(languageCodeForm, `-`))) {

		// Case: Only one language available.
		resultLangs = make(Languages, 1)
		resultLangs[0].Factor = 1.0
		resultLangs[0].Language = strings.ToLower(languageCodeForm)
		return
	}

	// Case: The Accept-Language fild is empty.
	if languageCodeHeader == `` {
		// Use the default language:
		resultLangs = make(Languages, 1)
		resultLangs[0].Factor = 1.0
		resultLangs[0].Language = defaultLanguage
		return
	}

	// Separate each language:
	values := strings.Split(languageCodeHeader, `,`)
	langs := make(Languages, len(values))

	// Loop over all languages:
	for n, langData := range values {

		// Is the factor given?
		if factorData := strings.Split(langData, `;q=`); len(factorData) == 2 {

			// Try to parse the factor:
			if factor, errFactor := strconv.ParseFloat(factorData[1], 32); errFactor != nil {
				// Case: Parsing was not possible. Use 1 instead:
				langs[n] = Language{}
				langs[n].Language = strings.ToLower(strings.Trim(factorData[0], ` `))
				langs[n].Factor = 1.0
			} else {
				// Case: Factor was parsed.
				langs[n] = Language{}
				langs[n].Language = strings.ToLower(strings.Trim(factorData[0], ` `))
				langs[n].Factor = float32(factor)
			}
		} else {
			// Case: No factor given. Use 1 instead.
			langs[n] = Language{}
			langs[n].Language = strings.ToLower(strings.Trim(langData, ` `))
			langs[n].Factor = 1.0
		}
	}

	// Sort all languages by factor:
	sort.Sort(langs)
	resultLangs = langs
	return
}

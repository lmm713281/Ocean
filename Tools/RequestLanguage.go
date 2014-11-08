package Tools

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func GetRequestLanguage(request *http.Request) (resultLangs Languages) {
	languageCodeForm := request.FormValue(`lang`)
	languageCodeHeader := request.Header.Get(`Accept-Language`)

	if languageCodeForm != `` && ((len(languageCodeForm) == 2) || (len(languageCodeForm) == 5 && strings.Contains(languageCodeForm, `-`))) {
		resultLangs = make(Languages, 1)
		resultLangs[0].Factor = 1.0
		resultLangs[0].Language = strings.ToLower(languageCodeForm)
		return
	}

	if languageCodeHeader == `` {
		resultLangs = make(Languages, 1)
		resultLangs[0].Factor = 1.0
		resultLangs[0].Language = defaultLanguage
		return
	}

	values := strings.Split(languageCodeHeader, `,`)
	langs := make(Languages, len(values))
	for n, langData := range values {
		if factorData := strings.Split(langData, `;q=`); len(factorData) == 2 {
			if factor, errFactor := strconv.ParseFloat(factorData[1], 32); errFactor != nil {
				langs[n] = Language{}
				langs[n].Language = strings.ToLower(strings.Trim(factorData[0], ` `))
				langs[n].Factor = 1.0
			} else {
				langs[n] = Language{}
				langs[n].Language = strings.ToLower(strings.Trim(factorData[0], ` `))
				langs[n].Factor = float32(factor)
			}
		} else {
			langs[n] = Language{}
			langs[n].Language = strings.ToLower(strings.Trim(langData, ` `))
			langs[n].Factor = 1.0
		}
	}

	sort.Sort(langs)
	resultLangs = langs
	return
}

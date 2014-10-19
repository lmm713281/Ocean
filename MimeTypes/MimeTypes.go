package MimeTypes

import (
	"strings"
)

type MimeType struct {
	MimeType      string
	FileExtension []string
}

func DetectType(filename string) (mime MimeType, err error) {
	for _, typeElement := range allTypes {
		for _, extension := range typeElement.FileExtension {
			if strings.HasSuffix(filename, extension) {
				mime = typeElement
				return
			}
		}
	}

	mime = TypeUnknown
	return
}

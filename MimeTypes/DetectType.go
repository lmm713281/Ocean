package MimeTypes

import (
	"strings"
)

// A function to detect a MIME type.
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

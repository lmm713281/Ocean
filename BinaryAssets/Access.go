package BinaryAssets

import (
	"github.com/SommerEngineering/Ocean/BinaryAssets/SourceCodePro"
)

// Reads the content for a file.
func GetData(filename string) (data []byte) {
	if obj, err := SourceCodePro.Asset(filename); err != nil {
		return
	} else {
		data = obj
		return
	}
}

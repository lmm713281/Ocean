package Configuration

import (
	"encoding/json"
	"github.com/SommerEngineering/Ocean/Log"
	"github.com/SommerEngineering/Ocean/Log/Meta"
	"os"
	"path/filepath"
)

// Function to read the configuration file.
func readConfiguration() {
	if isInit {
		Log.LogFull(senderName, Meta.CategorySYSTEM, Meta.LevelWARN, Meta.SeverityNone, Meta.ImpactNone, Meta.MessageNameINIT, `The configuration package is already init!`)
		return
	} else {
		Log.LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, Meta.MessageNameCONFIGURATION, `Init of configuration starting.`)
	}

	// Access to the working directory?
	currentDir, dirError := os.Getwd()
	if dirError != nil {
		panic(`Was not able to read the working directory: ` + dirError.Error())
		return
	}

	// Access to the configuration file?
	currentPath := filepath.Join(currentDir, filename)
	if _, errFile := os.Stat(currentPath); errFile != nil {
		if os.IsNotExist(errFile) {
			panic(`It was not possible to find the necessary configuration file 'configuration.json' at the application directory.`)
		} else {
			panic(`There was an error while open the configuration: ` + errFile.Error())
		}
	}

	// Open the file:
	file, fileError := os.Open(currentPath)
	defer file.Close()

	if fileError != nil {
		panic(`The configuration file is not accessible: ` + fileError.Error())
		return
	}

	// Try to decode / parse the file:
	decoder := json.NewDecoder(file)
	decError := decoder.Decode(&configuration)

	if decError != nil {
		panic(`Decoding of the configuration file was not possible: ` + decError.Error())
	}

	Log.LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, Meta.MessageNameINIT, `Init of configuration is done.`)
	isInit = true
	return
}

package Configuration

import (
	"encoding/json"
	"fmt"
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
		fmt.Printf("[Error] Was not able to read the working directory. %s\n", dirError.Error())
		os.Exit(0)
	}

	// Access to the configuration file?
	currentPath := filepath.Join(currentDir, filename)
	if _, errFile := os.Stat(currentPath); errFile != nil {
		if os.IsNotExist(errFile) {
			fmt.Printf("[Error] Cannot read the project name file 'configuration.json': File not found! Please read https://github.com/SommerEngineering/Ocean\n")
		} else {
			fmt.Printf("[Error] Cannot read the project name file 'configuration.json': %s. Please read https://github.com/SommerEngineering/Ocean\n", errFile.Error())
		}
		os.Exit(0)
	}

	// Open the file:
	file, fileError := os.Open(currentPath)
	defer file.Close()

	if fileError != nil {
		fmt.Printf("[Error] The configuration file 'configuration.json' is not accessible: %s. Please read https://github.com/SommerEngineering/Ocean\n", fileError.Error())
		os.Exit(0)
	}

	// Try to decode / parse the file:
	decoder := json.NewDecoder(file)
	decError := decoder.Decode(&configuration)

	if decError != nil {
		fmt.Printf("[Error] Decoding of the configuration file 'configuration.json' was not possible: %s. Please read https://github.com/SommerEngineering/Ocean\n", decError.Error())
		os.Exit(0)
	}

	Log.LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, Meta.MessageNameINIT, `Init of configuration is done.`)
	isInit = true
	return
}

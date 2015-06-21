package Log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Read the project name out of the local configuration file "project.name".
func readProjectName() {
	// Try to get access to the working directory:
	if currentDir, dirError := os.Getwd(); dirError != nil {
		// Case: Error! Stop the server.
		fmt.Printf("[Error] Was not able to read the working directory. %s\n", dirError.Error())
		os.Exit(0)
	} else {
		// Try to get access to the file:
		filename := filepath.Join(currentDir, `project.name`)
		if _, errFile := os.Stat(filename); errFile != nil {
			// Cases: Error.
			if os.IsNotExist(errFile) {
				fmt.Printf("[Error] Cannot read the project name file 'project.name': File not found! Please read https://github.com/SommerEngineering/Ocean\n")
			} else {
				fmt.Printf("[Error] Cannot read the project name file 'project.name': %s. Please read https://github.com/SommerEngineering/Ocean\n", errFile.Error())
			}
			os.Exit(0)
		}

		// Try to read the file:
		if projectNameBytes, errRead := ioutil.ReadFile(filename); errRead != nil {
			// Case: Error.
			fmt.Printf("[Error] Cannot read the project name file 'project.name': %s. Please read https://github.com/SommerEngineering/Ocean\n", errRead.Error())
			os.Exit(0)
		} else {
			// Store the project name:
			projectName = string(projectNameBytes)
			projectName = strings.TrimSpace(projectName)
		}
	}
}

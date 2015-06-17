package Log

import (
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
		panic(`Cannot read the current working directory and therefore cannot read the project name!`)
	} else {
		// Try to get access to the file:
		filename := filepath.Join(currentDir, `project.name`)
		if _, errFile := os.Stat(filename); errFile != nil {
			// Cases: Error.
			if os.IsNotExist(errFile) {
				panic(`Cannot read the project name file 'project.name': File not found!`)
			} else {
				panic(`Cannot read the project name file 'project.name': ` + errFile.Error())
			}
		}

		// Try to read the file:
		if projectNameBytes, errRead := ioutil.ReadFile(filename); errRead != nil {
			// Case: Error.
			panic(`Cannot read the project name file 'project.name': ` + errRead.Error())
		} else {
			// Store the project name:
			projectName = string(projectNameBytes)
			projectName = strings.TrimSpace(projectName)
		}
	}
}

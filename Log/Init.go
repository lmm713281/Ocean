package Log

import "github.com/SommerEngineering/Ocean/Log/Meta"
import "strconv"
import "container/list"
import "sync"
import "path/filepath"
import "os"
import "io/ioutil"
import "strings"

func readProjectName() {
	if currentDir, dirError := os.Getwd(); dirError != nil {
		panic(`Can not read the current working directory and therefore can not read the project name!`)
	} else {
		filename := filepath.Join(currentDir, `project.name`)
		if _, errFile := os.Stat(filename); errFile != nil {
			if os.IsNotExist(errFile) {
				panic(`Can not read the project name file 'project.name': File not found!`)
			} else {
				panic(`Can not read the project name file 'project.name': ` + errFile.Error())
			}
		}

		if projectNameBytes, errRead := ioutil.ReadFile(filename); errRead != nil {
			panic(`Can not read the project name file 'project.name': ` + errRead.Error())
		} else {
			projectName = string(projectNameBytes)
			projectName = strings.TrimSpace(projectName)
		}
	}
}

func init() {
	readProjectName()
	mutexDeviceDelays = sync.Mutex{}
	mutexPreChannelBuffer = sync.Mutex{}
	mutexChannel = sync.RWMutex{}
	preChannelBuffer = list.New()
	deviceDelayBuffer = list.New()
	devices = list.New()

	initTimer()
	initCode()
}

func initCode() {
	entriesBuffer = make(chan Meta.Entry, logBufferSize)

	LogShort(senderName, Meta.CategorySYSTEM, Meta.LevelINFO, `Starting`, `The logger is now starting.`, `logBufferSize=`+strconv.Itoa(int(logBufferSize)), `logBufferTimeoutSeconds=`+strconv.Itoa(int(logBufferTimeoutSeconds)))
	go scheduler(entriesBuffer)
}

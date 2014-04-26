package Tools

import "os"
import "time"
import "math/rand"
import "github.com/SommerEngineering/Ocean/ConfigurationDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {
	// Get this hostname:
	if hostText, errHost := os.Hostname(); errHost != nil {
		panic(`Was not able to read the hostname: ` + errHost.Error())
	} else {
		hostname = hostText
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `Log the hostname of this machine.`, hostname)
	}

	// Get all IP addresses:
	initIPAddresses4ThisHost()

	// Read the InternalCommPassword:
	internalCommPassword = ConfigurationDB.Read(`InternalCommPassword`)

	// Set the seed for random:
	rand.Seed(time.Now().Unix())
}

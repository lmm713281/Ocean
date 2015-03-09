package Tools

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

	// Build the local IP address and port:
	allHostsIPAddresses := ReadAllIPAddresses4ThisHost()
	port := ConfigurationDB.Read(`PublicWebServerPort`)
	if strings.Contains(allHostsIPAddresses[0], `:`) {
		// Case: IPv6
		localIPAddressAndPort = fmt.Sprintf("[%s]:%s", allHostsIPAddresses[0], port)
	} else {
		// Case: IPv4
		localIPAddressAndPort = fmt.Sprintf("%s:%s", allHostsIPAddresses[0], port)
	}

	// Read the default language:
	defaultLanguage = ConfigurationDB.Read(`DefaultLanguageCode`)
}

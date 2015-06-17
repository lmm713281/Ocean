package NumGen

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"strconv"
	"strings"
)

// Init this package.
func init() {
	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `Init the number generator.`)

	// Get the exklusive access to the channel list:
	channelListLock.Lock()
	defer channelListLock.Unlock()

	correctPassword = ConfigurationDB.Read(`InternalCommPassword`)
	activeHost := ConfigurationDB.Read(`NumGenActiveHosts`)
	isActive = strings.Contains(activeHost, Tools.ThisHostname())
	getHandler = ConfigurationDB.Read(`NumGenGetHandler`)

	if isActive {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.MessageNameCONFIGURATION, `The number generator is active on this host.`, `This host is producer and consumer.`)

		channelBufferSizeText := ConfigurationDB.Read(`NumGenBufferSize`)
		if bufferSizeNumber, errBufferSizeNumber := strconv.Atoi(channelBufferSizeText); errBufferSizeNumber != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityCritical, LM.ImpactMiddle, LM.MessageNameCONFIGURATION, `Was not able to parse the configuration value of NumGenBufferSize.`, errBufferSizeNumber.Error(), `Use the default value now!`)
		} else {
			channelBufferSize = bufferSizeNumber
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, `The buffer size for the number generator was loaded.`, `Buffer size=`+channelBufferSizeText)
		}

		channelList = make(map[string]chan int64)

		initDB()
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.MessageNameCONFIGURATION, `The number generator is not active on this host.`, `This host is just a consumer.`)
	}
}

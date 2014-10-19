package NumGen

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Shutdown"
)

func requestChannel4Name(name string) (result chan int64) {

	if Shutdown.IsDown() {
		return
	}

	if !isActive {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactNone, LM.MessageNameCONFIGURATION, `Called the requestChannel4Name() on an inactive host.`, `Wrong configuration?`)
		return
	}

	channelListLock.RLock()
	channel, isPresent := channelList[name]
	channelListLock.RUnlock()

	if isPresent {
		result = channel
		return
	}

	// Create the entry:
	newChannel := make(chan int64, channelBufferSize)
	result = newChannel

	channelListLock.Lock()
	channelList[name] = newChannel
	channelListLock.Unlock()

	// Create the new producer:
	go producer(name)
	return
}

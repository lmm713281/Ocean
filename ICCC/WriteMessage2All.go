package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// Function to broadcast a message to all listeners.
func WriteMessage2All(channel, command string, kind byte, message interface{}, answerPrototype interface{}) (results []interface{}) {
	cacheListenerDatabaseLock.RLock()
	defer cacheListenerDatabaseLock.RUnlock()

	// Convert the message to HTTP data:
	data := Message2Data(channel, command, message)

	// Store all matching listener:
	matchingListener := make([]Scheme.Listener, 0)

	// Loop over all listeners which are currently available at the cache:
	for entry := cacheListenerDatabase.Front(); entry != nil; entry = entry.Next() {
		listener := entry.Value.(Scheme.Listener)

		// If the channel, command and kind matches, deliver the message:
		if kind == KindALL {
			if listener.Channel == channel && listener.Command == command {
				matchingListener = append(matchingListener, listener)
			}
		} else {
			if listener.Channel == channel && listener.Command == command && listener.Kind == kind {
				matchingListener = append(matchingListener, listener)
			}
		}
	}

	// Was not able to find any matching listener?
	if len(matchingListener) == 0 {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not able to deliver this message, because no listener was found!`, `channel=`+channel, `command=`+command)
		return
	}

	// Create an channel to receive all the answers from all listeners:
	answerChannel := make(chan map[string][]string, len(matchingListener))

	// Start for every listener an own thread:
	for _, listener := range matchingListener {
		go func() {
			answerChannel <- sendMessage(listener, data)
		}()
	}

	// Reserve memory for the result data:
	results = make([]interface{}, len(matchingListener))

	// Read all answers:
	for n := 0; n < len(matchingListener); n++ {
		//
		// We use no timeout here. This way, it is also possible to execute
		// long running commands by ICCC. The caller can abort, if necessary.
		//
		answersData := <-answerChannel

		// Convert the data to the message type. The call will use answerPrototype
		// just as a prototype and will create a new instance for every call:
		_, _, results[n] = Data2Message(answerPrototype, answersData)
	}

	return
}

package ICCC

import (
	"github.com/SommerEngineering/Ocean/ICCC/Scheme"
	"github.com/SommerEngineering/Ocean/ICCC/SystemMessages"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

// The receiver function for the ICCC message, that yields the hosts.
func ICCCGetHostsReceiver(data map[string][]string) (result map[string][]string) {

	// Recover from errors:
	defer func() {
		if err := recover(); err != nil {
			Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNamePARSE, "Was not able to execute the ICCC get hosts message.")
			result = make(map[string][]string, 0)
			return
		}
	}()

	// Converts the HTTP form data into an object:
	_, _, obj := Data2Message(SystemMessages.ICCCGetHosts{}, data)

	// Was it possible to convert the data?
	if obj != nil {

		// We have to read from the cache:
		cacheHostDatabaseLock.RLock()

		// How many hosts we currently known?
		countHosts := cacheHostDatabase.Len()

		// Prepare the answer object:
		answerMessage := SystemMessages.ICCCGetHostsAnswer{}
		answerMessage.Hostnames = make([]string, countHosts, countHosts)
		answerMessage.IPAddressesPorts = make([]string, countHosts, countHosts)

		// Loop over all hosts which are currently available at the cache:
		n := 0
		for entry := cacheHostDatabase.Front(); entry != nil; entry = entry.Next() {
			host := entry.Value.(Scheme.Host)
			answerMessage.Hostnames[n] = host.Hostname
			answerMessage.IPAddressesPorts[n] = host.IPAddressPort
			n++
		}

		// Unlock the cache:
		cacheHostDatabaseLock.RUnlock()

		// Send the answer:
		return Message2Data(``, ``, answerMessage)
	} else {
		Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameSTARTUP, `ICCC message: Was not able to create the message.`)
	}

	// In any other error case:
	result = make(map[string][]string, 0)
	return
}

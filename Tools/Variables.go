package Tools

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

var (
	senderName            LM.Sender = `System::Tools`
	hostname              string    = `unknown`
	localIPAddressAndPort string    = `127.0.0.1:60000`
	ipAddresses           []string  = nil
	internalCommPassword  string    = ``
	defaultLanguage       string    = `en`
)

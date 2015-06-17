package Tools

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

var (
	senderName            LM.Sender = `System::Tools`   // This is the name for logging event from this package
	hostname              string    = `unknown`         // The local server's hostname
	localIPAddressAndPort string    = `127.0.0.1:60000` // The local server's IP address and port
	ipAddresses           []string  = nil               // All public IP addresses of this server
	internalCommPassword  string    = ``                // The communication password
	defaultLanguage       string    = `en`              // The default language for I18N
)

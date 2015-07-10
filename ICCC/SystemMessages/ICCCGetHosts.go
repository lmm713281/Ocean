package SystemMessages

// A message to request the known ICCC hosts.
type ICCCGetHosts struct {
}

// The answer to the hosts request.
type ICCCGetHostsAnswer struct {
	Hostnames        []string
	IPAddressesPorts []string
}

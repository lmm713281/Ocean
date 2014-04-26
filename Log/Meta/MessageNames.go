package Meta

type MessageName string

const (
	MessageNameSTARTUP       = `Startup`
	MessageNameINIT          = `Init`
	MessageNameSHUTDOWN      = `Shutdown`
	MessageNameEXECUTE       = `Execute`
	MessageNameDATABASE      = `Database`
	MessageNameNETWORK       = `Network`
	MessageNameLOGIN         = `Login`
	MessageNameLOGOUT        = `Logout`
	MessageNameSESSION       = `Session`
	MessageNameTIMEOUT       = `Timeout`
	MessageNameFILESYSTEM    = `Filesystem`
	MessageNameCOMMUNICATION = `Communication`
	MessageNameWRITE         = `Write`
	MessageNameREAD          = `Read`
	MessageNameALGORITHM     = `Algorithm`
	MessageNameCONFIGURATION = `Configuration`
	MessageNameTIMER         = `Timer`
	MessageNameINPUT         = `Input`
	MessageNameOUTPUT        = `Output`
	MessageNameBROWSER       = `Browser`
	MessageNameSECURITY      = `Security`
	MessageNameNOTFOUND      = `NotFound`
	MessageNameANALYSIS      = `Analysis`
	MessageNameSTATE         = `State`
	MessageNameGENERATOR     = `Generator`
	MessageNamePRODUCER      = `Producer`
	MessageNameCONSUMER      = `Consumer`
	MessageNamePASSWORD      = `Password`
)

package Meta

type MessageName string

// Some pre-defined message names:
const (
	MessageNameSTARTUP       = `Startup`       // e.g. the server startup
	MessageNameINIT          = `Init`          // some kind of init event
	MessageNameSHUTDOWN      = `Shutdown`      // some kind of shutdown event
	MessageNameEXECUTE       = `Execute`       // some kind of execution context
	MessageNameDATABASE      = `Database`      // events which are related to database issues
	MessageNameNETWORK       = `Network`       // events which are related to network issues
	MessageNameLOGIN         = `Login`         // some kind of login event
	MessageNameLOGOUT        = `Logout`        // some kind of logout event
	MessageNameSESSION       = `Session`       // some kind of session event
	MessageNameTIMEOUT       = `Timeout`       // some kind of timeout event
	MessageNameFILESYSTEM    = `Filesystem`    // events which are related to the filesystem
	MessageNameCOMMUNICATION = `Communication` // events which are related to communication issues
	MessageNameWRITE         = `Write`         // some kind of write event
	MessageNameREAD          = `Read`          // some kind of read event
	MessageNameALGORITHM     = `Algorithm`     // some kind of algorithm event
	MessageNameCONFIGURATION = `Configuration` // some kind of configuration event
	MessageNameTIMER         = `Timer`         // some kind of timer event
	MessageNameINPUT         = `Input`         // some kind of input event
	MessageNameOUTPUT        = `Output`        // some kind of output event
	MessageNameBROWSER       = `Browser`       // some kind of browser event
	MessageNameSECURITY      = `Security`      // some kind of security event
	MessageNameNOTFOUND      = `NotFound`      // something was not found
	MessageNameANALYSIS      = `Analysis`      // some kind of analysis event
	MessageNameSTATE         = `State`         // some kind of state-related event
	MessageNameGENERATOR     = `Generator`     // some kind of generator event
	MessageNamePRODUCER      = `Producer`      // some kind of producer event
	MessageNameCONSUMER      = `Consumer`      // some kind of consumer event
	MessageNamePASSWORD      = `Password`      // some kind of password event
	MessageNamePARSE         = `Parse`         // some kind of parser-related event
	MessageNameUSER          = `User`          // some kind of user event
	MessageNameREQUEST       = `Request`       // some kind of request-related event
	MessageNameRESPONSE      = `Response`      // some kind of response-related event
)

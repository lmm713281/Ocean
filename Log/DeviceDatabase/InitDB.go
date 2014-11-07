package DeviceDatabase

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"strconv"
	"strings"
	"time"
)

func initDatabase() {

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Checking and init the logging database collection.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Checking and init the logging database collection done.`)

	databaseHost := ConfigurationDB.Read(`LogDBHost`)
	databaseDB := ConfigurationDB.Read(`LogDBDatabase`)
	databaseUsername := ConfigurationDB.Read(`LogDBUsername`)
	databasePassword := ConfigurationDB.Read(`LogDBPassword`)
	expire := strings.ToLower(ConfigurationDB.Read(`LogDBEventsExpire`)) == `true`
	expireAfterDays := 36500

	if value, errValue := strconv.Atoi(ConfigurationDB.Read(`LogDBEventsExpireAfterDays`)); errValue != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not possible to read the configuration for the expire time of logging events. Log events will not expire any more.`, errValue.Error())
		expire = false
	} else {
		if expire {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("All logging events are expire after %d days.", value))
			expireAfterDays = value
		}
	}

	// Connect to MongoDB:
	if newSession, errDial := mgo.Dial(databaseHost); errDial != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to connect to the MongoDB host `+databaseHost, errDial.Error())
		return
	} else {
		logDBSession = newSession
	}

	// Use the correct database:
	logDB = logDBSession.DB(databaseDB)

	// Login:
	if errLogin := logDB.Login(databaseUsername, databasePassword); errLogin != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to login the user `+databaseUsername, errLogin.Error())
		return
	}

	// Get the collection:
	logDBCollection = logDB.C(`Logbook`)

	//
	// Take care about all the indexes:
	//
	indexTimeUTC := mgo.Index{}
	indexTimeUTC.Key = []string{`TimeUTC`}
	logDBCollection.EnsureIndex(indexTimeUTC)

	if expire {
		indexTTL := mgo.Index{}
		indexTTL.Key = []string{`TimeUTC`}
		indexTTL.ExpireAfter = time.Duration(expireAfterDays) * time.Hour * 24
		logDBCollection.EnsureIndex(indexTTL)
	}

	indexProject := mgo.Index{}
	indexProject.Key = []string{`Project`}
	logDBCollection.EnsureIndex(indexProject)

	indexSender := mgo.Index{}
	indexSender.Key = []string{`Sender`}
	logDBCollection.EnsureIndex(indexSender)

	indexCategory := mgo.Index{}
	indexCategory.Key = []string{`Category`}
	logDBCollection.EnsureIndex(indexCategory)

	indexLevel := mgo.Index{}
	indexLevel.Key = []string{`Level`}
	logDBCollection.EnsureIndex(indexLevel)

	indexSeverity := mgo.Index{}
	indexSeverity.Key = []string{`Severity`}
	logDBCollection.EnsureIndex(indexSeverity)

	indexImpact := mgo.Index{}
	indexImpact.Key = []string{`Impact`}
	logDBCollection.EnsureIndex(indexImpact)

	indexMessageName := mgo.Index{}
	indexMessageName.Key = []string{`MessageName`}
	logDBCollection.EnsureIndex(indexMessageName)

	indexMessageDescription := mgo.Index{}
	indexMessageDescription.Key = []string{`MessageDescription`}
	logDBCollection.EnsureIndex(indexMessageDescription)

	indexProjectTimeUTC := mgo.Index{}
	indexProjectTimeUTC.Key = []string{`Project`, `TimeUTC`}
	logDBCollection.EnsureIndex(indexProjectTimeUTC)

	indexProjectSender := mgo.Index{}
	indexProjectSender.Key = []string{`Project`, `Sender`}
	logDBCollection.EnsureIndex(indexProjectSender)

	indexProjectCategory := mgo.Index{}
	indexProjectCategory.Key = []string{`Project`, `Category`}
	logDBCollection.EnsureIndex(indexProjectCategory)

	indexProjectLevel := mgo.Index{}
	indexProjectLevel.Key = []string{`Project`, `Level`}
	logDBCollection.EnsureIndex(indexProjectLevel)

	indexProjectSeverity := mgo.Index{}
	indexProjectSeverity.Key = []string{`Project`, `Severity`}
	logDBCollection.EnsureIndex(indexProjectSeverity)

	indexProjectImpact := mgo.Index{}
	indexProjectImpact.Key = []string{`Project`, `Impact`}
	logDBCollection.EnsureIndex(indexProjectImpact)

	indexProjectMessageName := mgo.Index{}
	indexProjectMessageName.Key = []string{`Project`, `MessageName`}
	logDBCollection.EnsureIndex(indexProjectMessageName)

	indexProjectMessageDescription := mgo.Index{}
	indexProjectMessageDescription.Key = []string{`Project`, `MessageDescription`}
	logDBCollection.EnsureIndex(indexProjectMessageDescription)

	indexProjectTimeUTCSender := mgo.Index{}
	indexProjectTimeUTCSender.Key = []string{`Project`, `TimeUTC`, `Sender`}
	logDBCollection.EnsureIndex(indexProjectTimeUTCSender)

	indexProjectTimeUTCCategory := mgo.Index{}
	indexProjectTimeUTCCategory.Key = []string{`Project`, `TimeUTC`, `Category`}
	logDBCollection.EnsureIndex(indexProjectTimeUTCCategory)

	indexProjectTimeUTCLevel := mgo.Index{}
	indexProjectTimeUTCLevel.Key = []string{`Project`, `TimeUTC`, `Level`}
	logDBCollection.EnsureIndex(indexProjectTimeUTCLevel)

	indexProjectTimeUTCSeverity := mgo.Index{}
	indexProjectTimeUTCSeverity.Key = []string{`Project`, `TimeUTC`, `Severity`}
	logDBCollection.EnsureIndex(indexProjectTimeUTCSeverity)

	indexProjectTimeUTCImpact := mgo.Index{}
	indexProjectTimeUTCImpact.Key = []string{`Project`, `TimeUTC`, `Impact`}
	logDBCollection.EnsureIndex(indexProjectTimeUTCImpact)

	indexProjectTimeUTCMessageName := mgo.Index{}
	indexProjectTimeUTCMessageName.Key = []string{`Project`, `TimeUTC`, `MessageName`}
	logDBCollection.EnsureIndex(indexProjectTimeUTCMessageName)

	indexProjectTimeUTCMessageDescription := mgo.Index{}
	indexProjectTimeUTCMessageDescription.Key = []string{`Project`, `TimeUTC`, `MessageDescription`}
	logDBCollection.EnsureIndex(indexProjectTimeUTCMessageDescription)

	indexTimeUTCSender := mgo.Index{}
	indexTimeUTCSender.Key = []string{`TimeUTC`, `Sender`}
	logDBCollection.EnsureIndex(indexTimeUTCSender)

	indexTimeUTCCategory := mgo.Index{}
	indexTimeUTCCategory.Key = []string{`TimeUTC`, `Category`}
	logDBCollection.EnsureIndex(indexTimeUTCCategory)

	indexTimeUTCLevel := mgo.Index{}
	indexTimeUTCLevel.Key = []string{`TimeUTC`, `Level`}
	logDBCollection.EnsureIndex(indexTimeUTCLevel)

	indexTimeUTCSeverity := mgo.Index{}
	indexTimeUTCSeverity.Key = []string{`TimeUTC`, `Severity`}
	logDBCollection.EnsureIndex(indexTimeUTCSeverity)

	indexTimeUTCImpact := mgo.Index{}
	indexTimeUTCImpact.Key = []string{`TimeUTC`, `Impact`}
	logDBCollection.EnsureIndex(indexTimeUTCImpact)

	indexTimeUTCMessageName := mgo.Index{}
	indexTimeUTCMessageName.Key = []string{`TimeUTC`, `MessageName`}
	logDBCollection.EnsureIndex(indexTimeUTCMessageName)

	indexTimeUTCMessageDescription := mgo.Index{}
	indexProjectTimeUTCMessageDescription.Key = []string{`TimeUTC`, `MessageDescription`}
	logDBCollection.EnsureIndex(indexTimeUTCMessageDescription)

}

package DeviceDatabase

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"strconv"
	"strings"
	"time"
)

// Init the database for the logging.
func initDatabase() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[Error] Was not able to connect to the logging database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", err)
			os.Exit(0)
		}
	}()

	Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Checking and init the logging database collection.`)
	defer Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameINIT, `Checking and init the logging database collection done.`)

	// Read the configuration values for the logging database:
	databaseHost := ConfigurationDB.Read(`LogDBHost`)
	databaseDB := ConfigurationDB.Read(`LogDBDatabase`)
	databaseUsername := ConfigurationDB.Read(`LogDBUsername`)
	databasePassword := ConfigurationDB.Read(`LogDBPassword`)

	// Should the logging events at the database expire?
	expire := strings.ToLower(ConfigurationDB.Read(`LogDBEventsExpire`)) == `true`

	// The default values for the TTL (time to live):
	expireAfterDays := 21900              // 60 years ~ maximum of MongoDB
	expireValue4DisabledFunction := 21900 // 60 years ~ maximum of MongoDB

	// Try to read the configured value for the TTL:
	if value, errValue := strconv.Atoi(ConfigurationDB.Read(`LogDBEventsExpireAfterDays`)); errValue != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactUnknown, LM.MessageNameCONFIGURATION, `It was not possible to read the configuration for the expire time of logging events. Log events will not expire any more.`, errValue.Error())
		expire = false
	} else {
		if expire {
			Log.LogShort(senderName, LM.CategorySYSTEM, LM.LevelINFO, LM.MessageNameCONFIGURATION, fmt.Sprintf("All logging events are expire after %d days.", value))
			if value > expireValue4DisabledFunction {
				Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityLow, LM.ImpactLow, LM.MessageNameDATABASE, fmt.Sprintf("Cannot set the logging database's TTL to %d, because MongoDB does not allow more than %d (63 years). Use now the maximum instead.", value, expireValue4DisabledFunction))
			} else {
				expireAfterDays = value
			}
		}
	}

	// Connect to MongoDB:
	if newSession, errDial := mgo.Dial(databaseHost); errDial != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelERROR, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to connect to the MongoDB host `+databaseHost, errDial.Error())
		fmt.Printf("[Error] Was not able to connect to the logging database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", errDial.Error())
		os.Exit(0)
	} else {
		logDBSession = newSession
	}

	// Use the correct database:
	logDB = logDBSession.DB(databaseDB)

	// Login:
	if errLogin := logDB.Login(databaseUsername, databasePassword); errLogin != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to login the user `+databaseUsername, errLogin.Error())
		fmt.Printf("[Error] Was not able to connect to the logging database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", errLogin.Error())
		os.Exit(0)
	}

	// Get the collection:
	logDBCollection = logDB.C(`Logbook`)

	//
	// Take care about all the indexes:
	//
	expireAfterSeconds := expireAfterDays * 24 * 60 * 60
	indexTimeUTC := mgo.Index{}
	indexTimeUTC.Key = []string{`TimeUTC`}
	indexTimeUTC.ExpireAfter = time.Duration(expireValue4DisabledFunction * 24 * 60 * 60)
	logDBCollection.EnsureIndex(indexTimeUTC)

	// Update the expire policy:
	updateResult := TTLUpdateResult{}
	updateCommand := bson.D{
		{`collMod`, `Logbook`},
		{`index`,
			bson.D{
				{`keyPattern`, bson.D{{`TimeUTC`, 1}}},
				{`expireAfterSeconds`, expireAfterSeconds},
			},
		},
	}

	if errUpdate := logDB.Run(updateCommand, &updateResult); errUpdate != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to update the expire policy for the logging database.`, errUpdate.Error())
	} else {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelWARN, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, fmt.Sprintf(`Update the expire policy for the logging database done.`))
	}

	//
	// Ensure that all necessary indexes are existing:
	//
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

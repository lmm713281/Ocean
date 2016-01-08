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
			os.Exit(2)
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
		os.Exit(2)
	} else {
		logDBSession = newSession
	}

	// Use the correct database:
	logDB = logDBSession.DB(databaseDB)

	// Login:
	if errLogin := logDB.Login(databaseUsername, databasePassword); errLogin != nil {
		Log.LogFull(senderName, LM.CategorySYSTEM, LM.LevelSECURITY, LM.SeverityUnknown, LM.ImpactUnknown, LM.MessageNameDATABASE, `It was not possible to login the user `+databaseUsername, errLogin.Error())
		fmt.Printf("[Error] Was not able to connect to the logging database: %s. Please read https://github.com/SommerEngineering/Ocean.\n", errLogin.Error())
		os.Exit(2)
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
	logDBCollection.EnsureIndexKey(`Sender`)
	logDBCollection.EnsureIndexKey(`Category`)
	logDBCollection.EnsureIndexKey(`Level`)
	logDBCollection.EnsureIndexKey(`Severity`)
	logDBCollection.EnsureIndexKey(`Impact`)
	logDBCollection.EnsureIndexKey(`MessageName`)
	logDBCollection.EnsureIndexKey(`MessageDescription`)
	logDBCollection.EnsureIndexKey(`Project`, `Sender`)
	logDBCollection.EnsureIndexKey(`Project`, `Category`)
	logDBCollection.EnsureIndexKey(`Project`, `Level`)
	logDBCollection.EnsureIndexKey(`Project`, `Severity`)
	logDBCollection.EnsureIndexKey(`Project`, `Impact`)
	logDBCollection.EnsureIndexKey(`Project`, `MessageName`)
	logDBCollection.EnsureIndexKey(`Project`, `MessageDescription`)
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `Sender`)

	// Related to the logging viewer:
	logDBCollection.EnsureIndexKey(`Project`)                                                                                 // Logging viewer, case: No filter
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`)                                                                     // Logging viewer, case: Filter for time
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `Sender`, `MessageName`, `Level`, `Category`, `Impact`, `Severity`) // Logging viewer, case: All filters are active
	logDBCollection.EnsureIndexKey(`Project`, `Sender`, `MessageName`, `Level`, `Category`, `Impact`, `Severity`)             // Logging viewer, case: All filters are active, but no time filter
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `Level`, `Category`)                                                // Logging viewer, case: Filter for e.g. app errors from yesterday
	logDBCollection.EnsureIndexKey(`Project`, `Level`, `Category`)                                                            // Logging viewer, case: Filter for e.g. all app errors

	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `Category`)
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `Level`)
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `Severity`)
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `Impact`)
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `MessageName`)
	logDBCollection.EnsureIndexKey(`Project`, `-TimeUTC`, `MessageDescription`)
	logDBCollection.EnsureIndexKey(`-TimeUTC`, `Sender`)
	logDBCollection.EnsureIndexKey(`-TimeUTC`, `Category`)
	logDBCollection.EnsureIndexKey(`-TimeUTC`, `Level`)
	logDBCollection.EnsureIndexKey(`-TimeUTC`, `Severity`)
	logDBCollection.EnsureIndexKey(`-TimeUTC`, `Impact`)
	logDBCollection.EnsureIndexKey(`-TimeUTC`, `MessageName`)
	logDBCollection.EnsureIndexKey(`-TimeUTC`, `MessageDescription`)
}

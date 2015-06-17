package DeviceDatabase

import (
	"time"
)

// The type for the database logging.
type LogDBEntry struct {
	TimeUTC            time.Time `bson:"TimeUTC"`
	Project            string    `bson:"Project"`
	Sender             string    `bson:"Sender"`
	Category           string    `bson:"Category"`
	Level              string    `bson:"Level"`
	Severity           string    `bson:"Severity"`
	Impact             string    `bson:"Impact"`
	MessageName        string    `bson:"MessageName"`
	MessageDescription string    `bson:"MessageDescription"`
	Parameters         []string  `bson:"Parameters"`
}

// A type for the TTL (time to live) for the index.
type TTLUpdateResult struct {
	ExpireAfterSeconds_old int32 `bson:"expireAfterSeconds_old"`
	ExpireAfterSeconds_new int32 `bson:"expireAfterSeconds_new"`
	Ok                     int32 `bson:"ok"`
}

// The logging device.
type Database struct {
}

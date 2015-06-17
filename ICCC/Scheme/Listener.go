package Scheme

// Type for the listener entries at the database.
type Listener struct {
	Channel       string `bson:"Channel"`
	Command       string `bson:"Command"`
	IsActive      bool   `bson:"IsActive"`
	IPAddressPort string `bson:"IPAddressPort"`
}

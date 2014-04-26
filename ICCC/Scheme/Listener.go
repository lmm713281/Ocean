package Scheme

type Listener struct {
	Channel       string `bson:"Channel"`
	Command       string `bson:"Command"`
	IsActive      bool   `bson:"IsActive"`
	IPAddressPort string `bson:"IPAddressPort"`
}

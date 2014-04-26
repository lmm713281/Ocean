package Scheme

type Host struct {
	Hostname      string `bson:"Hostname"`
	IPAddressPort string `bson:"IPAddressPort"`
}

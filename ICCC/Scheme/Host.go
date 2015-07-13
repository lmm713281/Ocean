package Scheme

// Scheme for the host database entry.
type Host struct {
	Hostname      string `bson:"Hostname"`
	IPAddressPort string `bson:"IPAddressPort"`
	Kind          byte   `bson:"Kind"`
}

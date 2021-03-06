package CustomerDB

import (
	"gopkg.in/mgo.v2"
)

// Get the database instance of the mgo MongoDB driver.
func DB() (session *mgo.Session, database *mgo.Database) {
	session = mainSession.Copy()
	database = session.DB(databaseDB)
	database.Login(databaseUsername, databasePassword)
	return
}

// Get directly the GridFS instance of the mgo MongoDB driver.
func GridFS() (session *mgo.Session, filesystem *mgo.GridFS) {
	session = mainSession.Copy()
	database := session.DB(databaseDB)
	database.Login(databaseUsername, databasePassword)
	filesystem = database.GridFS(`fs`)
	return
}

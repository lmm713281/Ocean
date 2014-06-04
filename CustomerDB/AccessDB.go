package CustomerDB

import "labix.org/v2/mgo"

/*
Get the database instance of the MGo Mongo driver.
*/
func DB() (session *mgo.Session, database *mgo.Database) {
	session = mainSession.Copy()
	database = session.DB(databaseDB)
	database.Login(databaseUsername, databasePassword)

	return
}

/*
Get directly the GridFS instance of the Mgo Mongo driver.
*/
func GridFS() (session *mgo.Session, filesystem *mgo.GridFS) {
	session = mainSession.Copy()
	database := session.DB(databaseDB)
	database.Login(databaseUsername, databasePassword)
	filesystem = database.GridFS(`fs`)

	return
}

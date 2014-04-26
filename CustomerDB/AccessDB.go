package CustomerDB

import "labix.org/v2/mgo"

/*
Get the database instance of the MGo Mongo driver.
*/
func DB() (result *mgo.Database) {
	result = db
	return
}

/*
Get directly the GridFS instance of the Mgo Mongo driver.
*/
func GridFS() (result *mgo.GridFS) {
	result = gridFS
	return
}

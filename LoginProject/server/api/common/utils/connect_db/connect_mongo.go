package connect_db

import (
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/dalhelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	mgo "gopkg.in/mgo.v2"
)

//GENERIC METHOD TO CONNECT TO DB
func ConnecttoMongoDB() (*mgo.Session, error) {
	logginghelper.LogDebug("IN: ConnecttoMongoDB")

	dbStatus, err := dalhelper.GetMongoConnection()

	if err != nil {
		logginghelper.LogError(" UNABLE TO CONNECT TO DB Error : ", err)
		return dbStatus, err
	}
	logginghelper.LogDebug("OUT: ConnecttoMongoDB")
	return dbStatus, nil
}

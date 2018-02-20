package connect_db

import (
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	mgo "gopkg.in/mgo.v2"
)

func ConnecttoMongoDB(RemoteIP string) (*mgo.Session, error) {
	logginghelper.LogDebug("IN: ConnecttoMongoDB")

	dbStatus, err := mgo.Dial(RemoteIP)

	if err != nil {
		logginghelper.LogError(" UNABLE TO CONNECT TO DB Error : ", err)
		return dbStatus, err
	}
	return dbStatus, nil
}

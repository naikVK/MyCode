package clientConfiguration

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils/connect_db"
	"errors"

	"gopkg.in/mgo.v2/bson"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

//GET CLIENT CONFIG FROM DB
func GetClientConfigeDAO(clientId string) (model.Client, error) {
	logginghelper.LogDebug("IN: GetGetClientConfigeDAO")
	var result model.Client
	dbStatus, err := connect_db.ConnecttoMongoDB()

	if err != nil {
		logginghelper.LogError("GetGetClientConfigeDAO UNABLE TO CONNECT TO DB Error : ", err)
		return result, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	dbname := confighelper.GetConfig("mongo.collection_clientdetails")
	cd := dbStatus.DB(name).C(dbname)

	Ferr := cd.Find(bson.M{
		"CLIENT_ID": clientId}).One(&result)
	if Ferr != nil {
		logginghelper.LogError("GetGetClientConfigeDAO CLIENT NOT PRESENT IN DB Error : ", Ferr)
		return model.Client{}, errors.New("CLIENT NOT FOUND ")
	}
	logginghelper.LogDebug("OUT: GetGetClientConfigeDAO")
	return result, err
}

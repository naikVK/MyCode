package login

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils/connect_db"
	"LoginProject/server/api/common/utils/security"
	"errors"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"gopkg.in/mgo.v2/bson"
)

//VERIFY USERNAME FROM DB
func GetUserByUsernameDAO(Username string) (model.Login, error) {
	logginghelper.LogDebug("IN: GetUserByUsernameDAO")
	result := model.Login{}
	dbStatus, err := connect_db.ConnecttoMongoDB()

	if err != nil {
		logginghelper.LogError("GetUserByUserIDDAO UNABLE TO CONNECT TO DB Error : ", err)
		return result, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	cd := dbStatus.DB(name).C(collection_name)

	Ferr := cd.Find(bson.M{
		"USERNAME": Username}).One(&result)
	if Ferr != nil {
		logginghelper.LogError("GetUserByUserIDDAO USER NOT PRESENT IN DB Error : ", Ferr)
		return model.Login{}, errors.New("USER NOT FOUND ")
	}
	logginghelper.LogDebug("OUT: GetUserByUsernameDAO")
	return result, err
}

//VERIFY USERNAME & PASSWORD FROM DB
func GetUserByLoginIDPasswordDAO(Username string, Password string) (model.Login, error) {
	logginghelper.LogDebug("IN: GetUserByLoginIDPasswordDAO")
	result := model.Login{}
	dbStatus, err := connect_db.ConnecttoMongoDB()

	if err != nil {
		logginghelper.LogError("GetUserByUserIDDAO UNABLE TO CONNECT TO DB Error : ", err)
		return result, err
	}
	logginghelper.LogDebug("password" + Password)
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	cd := dbStatus.DB(name).C(collection_name)
	hashedPassword := security.PasswordHashService(Password)
	logginghelper.LogDebug(hashedPassword)
	Ferr := cd.Find(bson.M{
		"USERNAME": Username}).One(&result)
	if Ferr != nil {
		logginghelper.LogError("GetUserByUserIDDAO USER NOT PRESENT IN DB Error : ", Ferr)
		return model.Login{}, errors.New("USER NOT FOUND ")
	}
	_, cmperr := security.CompareHashPasswordService(Password, []byte(result.Password))
	logginghelper.LogDebug("OUT: GetUserByLoginIDPasswordDAO")
	if cmperr != nil {
		logginghelper.LogError("hash error", cmperr)
		return model.Login{}, cmperr
	}
	return result, nil
}

//USED TO GET USER SCANNED QRCODE OR NOT
func GetUserQRcodeDAO(Username string) (model.Login, error) {
	logginghelper.LogDebug("IN: GetUserProfile")
	result := model.Login{}
	dbStatus, err := connect_db.ConnecttoMongoDB()

	if err != nil {
		logginghelper.LogError("GetUserByUserIDDAO UNABLE TO CONNECT TO DB Error : ", err)
		return result, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	cd := dbStatus.DB(name).C(collection_name)

	Ferr := cd.Find(bson.M{
		"USERNAME": Username}).One(&result)
	if Ferr != nil {
		logginghelper.LogError("GetUserByUserIDDAO USER NOT PRESENT IN DB Error : ", Ferr)
		return model.Login{}, errors.New("USER NOT FOUND ")
	}
	return result, err
}

//INSERT LOGIN DETAILS AFTER SUCCESSFULL REGISTRATION
func InsertLoginDetailsDAO(userObject model.Login) error {
	if (model.Login{}) == userObject {
		logginghelper.LogError("InsertLoginDetails: EMPTY LOGIN OBJ")
		return errors.New("USEROBJ IS EMPTY ")
	}
	session, err := connect_db.ConnecttoMongoDB()
	if nil != err {
		return err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	loginCollection := session.DB(name).C(collection_name)
	hashed := security.PasswordHashService(userObject.Password)
	userObject.Password = hashed
	dberr := loginCollection.Insert(userObject)
	if nil != dberr {
		logginghelper.LogError("Error in inserting in DB:", dberr)
		return dberr
	}
	return nil
}

//SET IF QR CODE SCANNED
func SetQRcodeDAO(Username string) (model.Login, error) {
	logginghelper.LogDebug("IN: GetUserProfile")
	result := model.Login{}
	dbStatus, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError("GetUserByUserIDDAO UNABLE TO CONNECT TO DB Error : ", err)
		return result, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	cd := dbStatus.DB(name).C(collection_name)

	Ferr := cd.Find(bson.M{
		"USERNAME": Username}).One(&result)
	if Ferr != nil {
		logginghelper.LogError("GetUserByUserIDDAO USER NOT PRESENT IN DB Error : ", Ferr)
		return model.Login{}, errors.New("USER NOT FOUND ")
	}
	result.Google_Auth.QRcodeScan = true
	Serr := cd.Update(bson.M{
		"USERNAME": Username}, result)
	if Serr != nil {
		logginghelper.LogError("GetUserByUserIDDAO USER NOT PRESENT IN DB Error : ", Ferr)
		return model.Login{}, errors.New("USER NOT FOUND ")
	}

	return result, err
}

//UPDATE PASSWORD IF  PASSWORD CHANGED
func UpdatePasswordDAO(userObject model.Login) (bool, error) {
	session, err := connect_db.ConnecttoMongoDB()
	if nil != err {
		return false, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	loginCollection := session.DB(name).C(collection_name)
	loginobj, err := GetUserByUsernameDAO(userObject.Username)
	if err != nil {
		logginghelper.LogError("USER_NOT_FOUND")
		return false, err
	}
	selector := bson.M{"USERNAME": loginobj.Username}
	hashed := security.PasswordHashService(userObject.Password)
	userObject.Password = hashed
	updator := bson.M{"$set": bson.M{"PASSWORD": userObject.Password}}
	err = loginCollection.Update(selector, updator)
	if err != nil {
		logginghelper.LogError("UPDATING_FAILED", err)
		return false, err
	}
	return true, nil
}

func ActivityloggedDAO(ActivityInfo model.ActivityLog) (bool, error) {
	session, err := connect_db.ConnecttoMongoDB()
	if nil != err {
		return false, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_activitydetails")
	activityCollection := session.DB(name).C(collection_name)

	dberr := activityCollection.Insert(ActivityInfo)
	if err != nil {
		logginghelper.LogError("INSERT_FAILED", err)
		return false, dberr
	}
	return true, nil
}

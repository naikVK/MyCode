package otp

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils/connect_db"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"gopkg.in/mgo.v2/bson"
)

// VERIFY OTP FROM DB
func VerifyOTPDAO(otpValues model.OTP) (bool, error) {
	logginghelper.LogDebug("VerifyOTPDAO() Start: ")
	dbStatus, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError("Database Connection Error : ", err)
		return false, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_otpverify")
	cd := dbStatus.DB(name).C(collection_name)
	result := model.OTP{}
	err = cd.Find(bson.M{"USERNAME": otpValues.Username, "OTP": otpValues.OTP}).One(&result)

	if err != nil {
		logginghelper.LogError("UNABLE_TO_DETECT_OTP :: ", err)
		return false, err
	} else {
		logginghelper.LogInfo("OTP verified")
		logginghelper.LogInfo("otp::", otpValues.OTP)
		info, err := cd.RemoveAll(bson.M{"USERNAME": otpValues.Username})
		logginghelper.LogInfo(info)
		if err != nil {
			logginghelper.LogError("Error while deleting data", err)
			return true, err
		}
		logginghelper.LogInfo("Deleted Successfully......")
		logginghelper.LogDebug("VerifyOTPDAO() End: ")
		return true, err
	}
	return false, nil
}

//INSERT OTP AFTER GENERATING
func InsertOTP(otpValues model.OTP) (bool, error) {
	logginghelper.LogDebug("InsertOTP start()")
	dbStatus, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError("Error while creating DBConnection", err)
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_otpverify")
	cd := dbStatus.DB(name).C(collection_name)

	dberr := cd.Insert(otpValues)
	if nil != dberr {
		logginghelper.LogError("Error inserting otp in db", dberr)
		return false, dberr
	}
	logginghelper.LogDebug("InsertOTP end()")
	return true, nil
}

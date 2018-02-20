package forgotPassword

import (
	"LoginProject/server/api/common/model"

	"LoginProject/server/api/common/utils/connect_db"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"gopkg.in/mgo.v2/bson"
)

func UpdateOTPDAO(otpValues model.OTP) (bool, error) {
	logginghelper.LogDebug("UpdateOTPDAO Start()")
	dbStatus, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError("Error while creating DBConnection")
		return false, err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_otpverify")
	cd := dbStatus.DB(name).C(collection_name)
	selector := bson.M{"USERNAME": otpValues.Username}
	updator := bson.M{"$set": bson.M{"OTP": otpValues.OTP}}
	info, dberr := cd.UpdateAll(selector, updator)
	logginghelper.LogInfo(info)
	if nil != dberr {
		logginghelper.LogError("Error updating otp in db", dberr)
		return false, dberr
	}
	logginghelper.LogDebug("UpdateOTPDAO End()")
	return true, nil
}

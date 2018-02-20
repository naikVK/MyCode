package fetchProfile

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils/connect_db"
	"LoginProject/server/api/common/utils/security"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"

	"gopkg.in/mgo.v2/bson"
)

// Fucntion to fetch the profile data based on username
func GetByUserName(username string) (model.ProfileDetail, bool) {
	logginghelper.LogInfo("Inside registrationDAO:: GetByUserName")
	profileDetails := model.ProfileDetail{}
	session, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError("Error while connecting to database", err)
		return model.ProfileDetail{}, false
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_registration")
	registrationCollection := session.DB(name).C(collection_name)
	findErr := registrationCollection.Find(bson.M{"USERNAME": username}).One(&profileDetails)
	if nil != findErr {
		logginghelper.LogError("User not found", findErr)
		return model.ProfileDetail{}, false
	}
	decryptFullName, decryptDOB, err := security.Decrypt(profileDetails.PersonalDetails.FullName, profileDetails.PersonalDetails.Dob)
	if err != nil {
		logginghelper.LogError("Error while decrypting data", err)
	}
	profileDetails.PersonalDetails.FullName = decryptFullName
	profileDetails.PersonalDetails.Dob = decryptDOB
	logginghelper.LogInfo(profileDetails)
	return profileDetails, true
}

package registration

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils/connect_db"
	"LoginProject/server/api/common/utils/security"
	"math/rand"
	"strings"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"gopkg.in/mgo.v2/bson"
)

func Insert(profileDetail model.ProfileDetail) error {
	logginghelper.LogInfo("Inside registrationDAO:: Insert")
	logginghelper.LogInfo(profileDetail)
	encryptFullName, encryptDOB, errr := security.Encrypt(profileDetail.PersonalDetails.FullName, profileDetail.PersonalDetails.Dob)
	if errr != nil {
		logginghelper.LogError(errr)
		return errr
	}
	profileDetail.PersonalDetails.FullName = encryptFullName
	profileDetail.PersonalDetails.Dob = encryptDOB

	session, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError(err)
		return err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_registration")
	registrationCollection := session.DB(name).C(collection_name)
	insertErr := registrationCollection.Insert(profileDetail)
	if nil != insertErr {
		logginghelper.LogError(insertErr)
		return err
	}
	return nil
}

func GetByUserName(username string) (model.ProfileDetail, bool) {
	logginghelper.LogInfo("Inside registrationDAO:: GetByUserName")
	profileDetails := model.ProfileDetail{}
	isUserPresent := true
	session, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		return model.ProfileDetail{}, isUserPresent
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_registration")
	registrationCollection := session.DB(name).C(collection_name)
	findErr := registrationCollection.Find(bson.M{"USERNAME": username}).One(&profileDetails)
	if nil != findErr {
		logginghelper.LogError("User not found")
		isUserPresent = false
		return model.ProfileDetail{}, isUserPresent
	}
	return profileDetails, isUserPresent
}

func GetUsernameSuggessions(profileDetails model.ProfileDetail, count int) []string {
	usernameList := make([]string, 0)
	usernamesMap := make(map[string]string, 0)
	usernamecnt := 0
	for usernamecnt < count {
		uniqueUsername := generateUsername(profileDetails)
		_, isUserPresent := GetByUserName(uniqueUsername)
		if !isUserPresent {
			if _, ok := usernamesMap[uniqueUsername]; !ok {
				usernamesMap[uniqueUsername] = uniqueUsername
				usernamecnt++
			}
		}
	}

	for _, username := range usernamesMap {
		usernameList = append(usernameList, username)
	}
	return usernameList
}

func generateUsername(profileDetail model.ProfileDetail) string {

	dates := strings.Split(profileDetail.PersonalDetails.Dob, "-")
	fullname := strings.Trim(profileDetail.PersonalDetails.FullName, " ")
	names := strings.Split(fullname, " ")
	initials := make([]string, 0)

	for _, val := range names {
		initials = append(initials, strings.ToUpper(strings.Split(val, "")[0]))
	}
	rand.Seed(time.Now().UTC().UnixNano())
	var uniqueUsername string
	uniqueUsername = names[rand.Intn(len(names))]
	uniqueUsername += initials[rand.Intn(len(initials))]
	uniqueUsername += dates[rand.Intn(len(dates))]

	return uniqueUsername
}

package profileModule

import (
	"net/http"
	"DummySolar/server/api/common/utils/jwtUtils"
	"github.com/labstack/echo"
	"DummySolar/server/api/common/model"
	"DummySolar/server/api/common/constants"

	"gopkg.in/mgo.v2/bson"
	"fmt"
		"gopkg.in/mgo.v2"
		"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
		"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

func Init(o *echo.Group, r *echo.Group, c *echo.Group) {
	r.GET("/getData", GetData)
}

func GetData(c echo.Context) error {
	fmt.Println("inside profile module: GetData")
	confighelper.InitViper()
	secretKey := confighelper.GetConfig("serverConfig.jwtSecretKey")
	tokenClaims, tokenError := jwtUtils.GetDecodedLoginFromToken(c, secretKey)
	if tokenError != nil {
		logginghelper.LogError("error occured while calling GetDecodedLoginFromToken ", tokenError)
		return echo.ErrUnauthorized
	}
	username,_ := tokenClaims["username"].(string)
	fmt.Println(username)
	profileDetails, isPresent  := GetByUserName(username)
	if !isPresent {
		return c.JSON(http.StatusExpectationFailed, "failed")	
	}
	return c.JSON(http.StatusOK, profileDetails)
}


func GetByUserName(username string) (model.ProfileDetail, bool) {
	logginghelper.LogInfo("Inside registrationDAO:: GetByUserName")
	profileDetails := model.ProfileDetail{}
	isUserPresent := true
	session, err := mgo.Dial(constants.DB_CONNECTION_ADD)
	if err != nil {
		return model.ProfileDetail{}, isUserPresent
	}
	fmt.Println(username)
	registrationCollection := session.DB(constants.DB_NAME).C(constants.COLLECTION_REGISTRATION)
	findErr := registrationCollection.Find(bson.M{"USERNAME": username}).One(&profileDetails)
	if nil != findErr {
		logginghelper.LogError("User not found")
		isUserPresent = false
		return model.ProfileDetail{}, isUserPresent
	}
	return profileDetails, isUserPresent
}

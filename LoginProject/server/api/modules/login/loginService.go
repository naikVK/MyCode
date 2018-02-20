package login

import (
	"LoginProject/server/api/common/constants"
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils"
	"LoginProject/server/api/common/utils/fetchProfile"
	"LoginProject/server/api/common/utils/jwtUtils"
	"LoginProject/server/api/modules/otp"
	"io/ioutil"
	"net/http"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pquerna/ffjson/ffjson"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

//VALIDATE USERNAME FROM DB
func ValidateUsernameService(user model.Login) (model.Login, error) {
	logginghelper.LogDebug("IN: ValidateUsernameService")
	logindbObj, err := GetUserByUsernameDAO(user.Username)

	if err != nil {
		logginghelper.LogError("ValidateUsernameService GetUserByUsernameDAO : ", err)
		return logindbObj, err

	}
	logginghelper.LogDebug("OUT: ValidateUsernameService")
	return logindbObj, nil
}

//VALIDATE USERNAME & PASSWORD FROM DB
func ValidateCredentialsService(user model.Login) (model.Login, error) {
	logginghelper.LogDebug("IN: ValidateCredentialsService")
	logindbObj, err := GetUserByLoginIDPasswordDAO(user.Username, user.Password)

	if err != nil {
		logginghelper.LogError("ValidateCredentialsService GetUserByUserIDPasswordDAO : ", err)
		return model.Login{}, err

	}
	logginghelper.LogDebug("OUT: ValidateCredentialsService")
	return logindbObj, nil
}

//VERIFY CAPTCHA
func IsRobotService(CaptchaResponse string) bool {
	logginghelper.LogInfo("Inside Service: login::isRobot")
	res, _ := http.PostForm(constants.CAPTCHA_API, url.Values{
		"secret":   {constants.SECRET_KEY_CAPTCHA},
		"response": {CaptchaResponse},
	})
	logginghelper.LogInfo("cap res:", CaptchaResponse)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logginghelper.LogError("Error while reading recaptcha response body ", err)
		return true
	}
	captchaRes := model.CaptchResponse{}
	err = ffjson.Unmarshal(data, &captchaRes)
	if err != nil {
		logginghelper.LogError("Error while unmarshaling recaptcha response", err)
		return true
	}
	//CAPTCHA RESPONSE SUCCESS MEANS USER IS HUMAN
	if captchaRes.IsSuccess {
		return false
	} else {
		return true
	}
}

//SEND OTP PHONE SERVICE
func SendOtpPhoneService(user model.ProfileDetail) (model.ProfileDetail, bool) {
	logginghelper.LogDebug("IN: GetUserProfileService")

	profiledbObj, status := otp.SendOTPService(user, 1, constants.MSG_OTP_LOGIN)
	if status {
		return profiledbObj, true
	}
	return model.ProfileDetail{}, false
}

//SEND OTP EMAIL SERVICE

func SendOtpEmailService(user model.ProfileDetail) (model.ProfileDetail, bool) {
	logginghelper.LogDebug("IN: GetUserProfileService")
	profiledbObj, status := otp.SendOTPService(user, 2, constants.MSG_OTP_LOGIN)
	if status {
		return profiledbObj, true
	}
	return model.ProfileDetail{}, false
}

//GET USER FROFILE FROM DB
func CheckUserService(user model.ProfileDetail) model.ProfileDetail {
	logginghelper.LogDebug("IN: GetUserProfileService")
	profiledbObj, status := fetchProfile.GetByUserName(user.UserName)
	if status {
		return profiledbObj
	}
	return model.ProfileDetail{}
}

//GET TOKEN FOR RESTRICT URL
func GetTokenRestrictedurlService(user model.Login, clientID string) (string, error) {
	url_key := confighelper.GetConfig(constants.RESTRICTKEY_PATH)
	Random_key := utils.GetRandamStringOfLen(constants.KEYLEN)
	tokenString, err := jwtUtils.GenerateJwtToken(user, Random_key, clientID, url_key)

	if err != nil {
		logginghelper.LogError(" GetTokenRestrictedurlService GenerateToken : ", err)
		return tokenString, err
	}
	return tokenString, nil
}

func UpdatePasswordService(login model.Login) bool {
	logginghelper.LogDebug("UpdatePasswordService IN:")
	status, err := UpdatePasswordDAO(login)
	if err != nil {
		logginghelper.LogError("UpdatePasswordDAO Error", err)
		return status
	}
	logginghelper.LogDebug("UpdatePAsswordService OUT")
	return true
}

//DECODE TOKEN FOR RESTRICT URL
func GetDecodedTokenRestrictedurlService(c echo.Context) (jwt.MapClaims, error) {
	url_key := confighelper.GetConfig(constants.RESTRICTKEY_PATH)
	tokenClaims, err := jwtUtils.GetDecodedLoginFromToken(c, url_key)
	if err != nil {
		logginghelper.LogError(" GetTokenRestrictedurlService GenerateToken : ", err)
		return tokenClaims, err
	}
	return tokenClaims, nil
}

//LOG THE ACTIVITY
func ActivityloggedService(AcitivityInfo model.ActivityLog) (bool, error) {
	ActivityLoggedstatus, err := ActivityloggedDAO(AcitivityInfo)
	if err != nil {
		logginghelper.LogError("ActivityloggedService:", err)
		return ActivityLoggedstatus, err
	}
	return ActivityLoggedstatus, nil
}

package login

import (
	"LoginProject/server/api/common/constants"
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils"
	"LoginProject/server/api/common/utils/google_auth"
	"LoginProject/server/api/common/utils/jwtUtils"
	"LoginProject/server/api/common/utils/sms"
	"LoginProject/server/api/modules/clientConfiguration"
	"LoginProject/server/api/modules/otp"
	"LoginProject/server/redisSessionManager"
	"net/http"
	"strings"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"

	"github.com/labstack/echo"
)

//Initialize method
func Init(o *echo.Group, r *echo.Group, c *echo.Group) {
	//For Open API
	o.POST("/login/validateuser", ValidatePasswordRoute)
	r.GET("/logout", Logout)
	o.POST("/login/validateusername", ValidateUsernameRoute)
	o.POST("/login/googleauthapp", GoogleAuthApp)
	o.POST("/login/googleauthgetkey", GoogleAuthGetKey)
	o.POST("/login/googleauthappcheck", GoogleAuthAppCheck)
	o.POST("/login/sendOTPonPhone", sendOTPonPhone)
	o.POST("/login/sendOTPonEmail", sendOTPonEmail)
	o.POST("/login/verifyOTP", VerifyOTP)
	o.POST("/isValidToken", isValidToken)
}

func isValidToken(c echo.Context) error {
	logginghelper.LogError("Inside LoginRoute:: isValidToken")
	authModel := model.Auth{}
	c.Bind(&authModel)
	clientConfig, err := clientConfiguration.GetClientConfigeDAO(authModel.ClientId)
	if nil != err {
		return c.JSON(http.StatusBadRequest, "config not found for client: "+authModel.ClientId)
	}
	secretKey := clientConfig.Jwt
	tokenMap, tokenError := jwtUtils.ValidateToken(authModel.Token, secretKey)
	sessionKey, _ := tokenMap["sessionId"].(string)
	if tokenError != nil {
		logginghelper.LogError("error occured while calling GetDecodedLoginFromToken ", tokenError)
		return echo.ErrUnauthorized
	}
	// check token in session store
	gcActiveToken, cacheError := redisSessionManager.Get(sessionKey)
	if cacheError != nil {
		logginghelper.LogError("CheckForSession: error occured while fetching value from gcache ", cacheError)
	}

	// if failed to access session
	if cacheError != nil || gcActiveToken == "" {
		return c.JSON(http.StatusUnauthorized, "SESSION_EXPIRED")
	}
	// if tokenFromRequest dont match with session token

	if authModel.Token != ("Bearer " + gcActiveToken) {
		logginghelper.LogInfo("activation token from request is not matching with activation token from gcache")
		return c.JSON(http.StatusBadRequest, "SESSION_EXPIRED")
	} else {
		// token from request is equal to token from redis sessionStore token
		// sliding session
		redisSessionManager.SlideSession(sessionKey)
	}
	return c.JSON(http.StatusOK, "token valid")
}

//Validates Username From DB
func ValidateUsernameRoute(c echo.Context) error {
	logginghelper.LogDebug("IN: ValidateUsernameRoute")
	login := model.Login{}
	err := c.Bind(&login)
	logginghelper.LogInfo(login.Username)
	if err != nil {
		logginghelper.LogError("ValidateUsernameRoute Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_PARAMETER_BIND_ERROR")
	}
	if strings.Trim(login.Username, "") == "" {
		logginghelper.LogError("GetLearnerProfileRoute USERNAME  is empty")
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_REQUIRED_FIELD_VALIDATION_FAILED")
	}
	resultObj, error := ValidateUsernameService(login)
	if error != nil {
		logginghelper.LogError("ValidateUsernameRoute FAIL TO FETCH DATA : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_NO USER FOUND ")
	}
	logginghelper.LogDebug("OUT: ValidateUsernameRoute")
	return c.JSON(http.StatusOK, resultObj.Username)
}

//Validate Username & Password
func ValidatePasswordRoute(c echo.Context) error {
	logginghelper.LogDebug("IN: ValidatePasswordRoute")
	login := model.Login{}
	err := c.Bind(&login)
	if err != nil {
		logginghelper.LogError("ValidatePasswordRoute Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "PARAMETER_BIND_ERROR")
	}
	if strings.Trim(login.Username, " ") == "" || strings.Trim(login.Password, " ") == "" {
		logginghelper.LogError("GetLearnerProfileRoute USERNAME OR PASSWORD is empty")
		return c.JSON(http.StatusExpectationFailed, "REQUIRED_FIELD_VALIDATION_FAILED")
	}
	requestClientId := c.Request().Header.Get(constants.CLIENTID)

	clientConfigForRequest, err := clientConfiguration.GetClientConfigeDAO(requestClientId)
	if nil != err {
		logginghelper.LogError("Fail to get client config:", err)
		return c.JSON(http.StatusBadRequest, "Failed to fetch config")
	}
	if clientConfigForRequest.Purpose.Settings.Captcha {
		captchaverified := IsRobotService(login.Captcha)
		if captchaverified {
			logginghelper.LogError("Captcha verification failed")
			return c.JSON(http.StatusInternalServerError, "Captcha verification failed")
		}
	}
	resultObj1, error := ValidateCredentialsService(login)

	if error != nil {
		logginghelper.LogError("ValidatePasswordRoute FAIL TO FETCH DATA : ", err)
		return c.JSON(http.StatusExpectationFailed, "USER NOT FOUND ")
	}
	//IF executes for two step auth
	logginghelper.LogInfo("TWO STEP VERIFICATION: ", clientConfigForRequest.Purpose.Settings.TwoStepAuth.Set)
	if clientConfigForRequest.Purpose.Settings.TwoStepAuth.Set {
		restrictJWTToken, err := GetTokenRestrictedurlService(login, requestClientId)
		if err != nil {
			logginghelper.LogError(" GetTokenRestrictedurlService GenerateToken : ", err)
			return c.JSON(http.StatusExpectationFailed, "GET TOKEN FAILED")
		}
		c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+restrictJWTToken)
		logginghelper.LogInfo(restrictJWTToken)
		return c.JSON(http.StatusOK, resultObj1.Google_Auth.QRcodeScan)
	}

	if redisSessionManager.IsSessionLimitReached(login.Username) {
		logginghelper.LogError("MAX_SESSION_LIMIT_REACHED")
		return c.JSON(http.StatusAlreadyReported, "MAX_SESSION_LIMIT_REACHED")
	}

	secretKey := clientConfigForRequest.Jwt

	jwtTOKEN, err := GetToken(login, requestClientId, secretKey)
	logginghelper.LogDebug(jwtTOKEN)
	if err != nil {
		logginghelper.LogError("ValidatePasswordRoute FAIL TO GENERATE TOKEN : ", err)
		return c.JSON(http.StatusExpectationFailed, "FAIL TO GENERATE TOKEN ")
	}
	c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+jwtTOKEN)
	CreateActivityInfo(login.Username)
	logginghelper.LogDebug("OUT: ValidatePasswordRoute")

	return c.JSON(http.StatusOK, "SUCCESS")
}

//Logout Kill session
func Logout(c echo.Context) error {
	logginghelper.LogDebug("IN: LogoutRoute")
	requestClientId := c.Request().Header.Get("ClientId")
	clientConfigForRequest, err := clientConfiguration.GetClientConfigeDAO(requestClientId)

	if nil != err {
		logginghelper.LogError("Failed to fetch config : ", err)
		return c.JSON(http.StatusBadRequest, "Failed to fetch config")
	}
	secretKey := clientConfigForRequest.Jwt
	// deleting session
	tokenClaims, err := jwtUtils.GetDecodedLoginFromToken(c, secretKey)
	if err != nil {
		logginghelper.LogError("LogoutRoute GetDecodedLoginFromToken : ", err)
		return c.JSON(http.StatusExpectationFailed, "LogoutRoute GetDecodedLoginFromToken Error")
	}
	// Removing session
	sessionKey := tokenClaims["sessionId"].(string)
	err = redisSessionManager.Del(sessionKey)
	if nil != err {
		logginghelper.LogError("LogoutRoute redis session deletion failed : ", err)
		return c.JSON(http.StatusInternalServerError, "Failed")
	}

	return c.JSON(http.StatusOK, "success")
}

// Creates QR code
func GoogleAuthApp(c echo.Context) error {
	logginghelper.LogDebug("IN: GoogleAuthApp")
	login := model.Login{}
	err := c.Bind(&login)
	if err != nil {
		logginghelper.LogError("GoogleAuthApp Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "PARAMETER_BIND_ERROR")
	}

	QR_Code_fileName, err := google_auth.CreateQRCode(login.Username)
	if err != nil {
		logginghelper.LogError("GoogleAuthApp Create qr code : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERROR")
	}
	logginghelper.LogDebug("OUT: GoogleAuthApp")
	return c.JSON(http.StatusOK, QR_Code_fileName)

}

//SEND GOOGLE_AUTH SECRET KEY VIA SMS
func GoogleAuthGetKey(c echo.Context) error {
	logginghelper.LogDebug("IN: GoogleAuthGetKey")
	login1 := model.Login{}

	err := c.Bind(&login1)
	logginghelper.LogInfo("username>>>", login1.Username)
	if err != nil {
		logginghelper.LogError("GoogleAuthApp Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERA_ERRORCODE_PARAMETER_BIND_ERROR")
	}

	secretKey, err := google_auth.GetSecretKey(login1.Username)
	if err != nil {
		logginghelper.LogError("GoogleAuthGetKey : ", err)
		return c.JSON(http.StatusExpectationFailed, "GoogleAuthGetKey KEY NOT FOUND")
	}
	profile := model.ProfileDetail{}
	profile.UserName = login1.Username
	profileOBJ := CheckUserService(profile)

	phoneNumber := profileOBJ.ContactDetails.Mobile.Number
	msg := constants.MSG_GOAUTH_SECRET_KEY + secretKey
	Sendmsg, err := sms.SendSingleSMS(msg, phoneNumber, "")
	if err != nil {
		logginghelper.LogError("Error while sending message", err)
		return err
	}
	logginghelper.LogInfo(Sendmsg)
	logginghelper.LogInfo(phoneNumber)
	logginghelper.LogDebug("OUT: GoogleAuthGetKey")
	return c.JSON(http.StatusOK, phoneNumber)

}

//VERIFY TOTP FOR GOOGLE AUTHENTICATOR
func GoogleAuthAppCheck(c echo.Context) error {
	logginghelper.LogDebug("IN: GoogleAuthAppCheck")
	OTPCheck := model.Google_AuthCheck{}
	logginghelper.LogInfo(OTPCheck.UserName)
	err := c.Bind(&OTPCheck)
	if err != nil {
		logginghelper.LogError("GoogleAuthApp Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERA_ERRORCODE_PARAMETER_BIND_ERROR")
	}
	tokenClaims, err := GetDecodedTokenRestrictedurlService(c)
	if err != nil {
		logginghelper.LogError("GoogleAuthAppCheck GetDecodedLoginFromToken : ", err)
		return c.JSON(http.StatusExpectationFailed, "GoogleAuthAppCheck GetDecodedLoginFromToken Error")
	}
	username := tokenClaims["username"].(string)
	logginghelper.LogInfo(tokenClaims)
	OTPCheck.UserName = username
	logginghelper.LogInfo(OTPCheck.UserName)
	QR_CodeCheck := google_auth.AuthenticateOTP(OTPCheck)
	if !QR_CodeCheck {
		logginghelper.LogDebug("OUT: GoogleAuthAppCheck")
		return c.JSON(http.StatusExpectationFailed, "OTP verification failed")
	}
	logginghelper.LogInfo(QR_CodeCheck)
	if redisSessionManager.IsSessionLimitReached(OTPCheck.UserName) {
		logginghelper.LogError("Sessio limit reached: ")
		return c.JSON(http.StatusAlreadyReported, "MAX_SESSION_LIMIT_REACHED")
	}
	requestClientId := c.Request().Header.Get("ClientId")
	clientConfigForRequest, err := clientConfiguration.GetClientConfigeDAO(requestClientId)
	if nil != err {
		return c.JSON(http.StatusAlreadyReported, "Failed to fing Config")
	}
	login := model.Login{}
	login.Username = OTPCheck.UserName
	secretKey := clientConfigForRequest.Jwt
	_, error := SetQRcodeDAO(login.Username)
	if error != nil {
		logginghelper.LogError("ValidateCredentialsRoute FAIL TO GENERATE TOKEN : ", err)
		return c.JSON(http.StatusExpectationFailed, "FAIL TO GENERATE TOKEN ")
	}

	jwtTOKEN, err := GetToken(login, requestClientId, secretKey)

	if err != nil {
		logginghelper.LogError("ValidateCredentialsRoute FAIL TO GENERATE TOKEN : ", err)
		return c.JSON(http.StatusExpectationFailed, "FAIL TO GENERATE TOKEN ")
	}

	c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+jwtTOKEN)
	CreateActivityInfo(OTPCheck.UserName)

	logginghelper.LogDebug("OUT: GoogleAuthAppCheck")
	return c.JSON(http.StatusOK, "success")
}

// GENERATE TOKEN WHEN USER IS AUTHENTICATED
func GetToken(loginobj model.Login, ClientID string, secretKey string) (string, error) {
	sessionKeyPostfix := utils.GetRandamStringOfLen(constants.KEYLEN)
	tokenString, err := jwtUtils.GenerateJwtToken(loginobj, sessionKeyPostfix, ClientID, secretKey)

	if err != nil {
		logginghelper.LogError("ValidateCredentialsRoute GenerateToken : ", err)
		return tokenString, err
	}

	sessionKey := loginobj.Username + ":" + sessionKeyPostfix
	err = redisSessionManager.Set(sessionKey, tokenString)
	if nil != err {
		logginghelper.LogError("Failed to set key", err)
		return "", err
	}
	return tokenString, nil
}

//SEND OTP FOR LOGIN ON PHONE VIA SMS
func sendOTPonPhone(c echo.Context) error {
	logginghelper.LogDebug("IN: sendOTPonPhone")
	profile := model.ProfileDetail{}
	err := c.Bind(&profile)
	if err != nil {
		logginghelper.LogError("sendOTPonPhone Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_PARAMETER_BIND_ERROR")
	}
	if strings.Trim(profile.UserName, "") == "" {
		logginghelper.LogError("sendOTPonPhone USERNAME  is empty")
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_REQUIRED_FIELD_VALIDATION_FAILED")
	}
	result, status := SendOtpPhoneService(profile)
	if !status {
		logginghelper.LogInfo("SendOtpService : ", err)
		return c.JSON(http.StatusExpectationFailed, "OTP SENDING FAILED")

	}
	phoneNumber := result.ContactDetails.Mobile.Number
	logginghelper.LogDebug("OUT: sendOTPonPhone")
	return c.JSON(http.StatusOK, phoneNumber)

}

//SEND OTP FOR LOGIN  ON EMAIL
func sendOTPonEmail(c echo.Context) error {
	logginghelper.LogDebug("IN: sendOTPonEmail")
	profile := model.ProfileDetail{}
	err := c.Bind(&profile)
	if err != nil {
		logginghelper.LogError("sendOTPonEmail Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_PARAMETER_BIND_ERROR")
	}
	if strings.Trim(profile.UserName, "") == "" {
		logginghelper.LogError("sendOTPonEmail USERNAME  is empty")
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_REQUIRED_FIELD_VALIDATION_FAILED")
	}
	result, status := SendOtpEmailService(profile)
	if !status {
		logginghelper.LogInfo("SendOtpService : ", err)
		return c.JSON(http.StatusExpectationFailed, "OTP SENDING FAILED")

	}
	emailID := result.ContactDetails.Email.Address
	logginghelper.LogDebug("OUT: sendOTPonEmail")
	return c.JSON(http.StatusOK, emailID)

}

//VERIFY OTP ON EMAIL OR PHONE
func VerifyOTP(c echo.Context) error {
	logginghelper.LogDebug("IN: VerifyOTP")
	OTPval := model.OTP{}
	err := c.Bind(&OTPval)
	if err != nil {
		logginghelper.LogError("VerifyOTP Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_PARAMETER_BIND_ERROR")
	}
	if strings.Trim(OTPval.Username, "") == "" {
		logginghelper.LogError("VerifyOTP USERNAME  is empty")
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_REQUIRED_FIELD_VALIDATION_FAILED")
	}

	OTPstatus, error := otp.VerifyOTPDAO(OTPval)
	if error != nil {
		logginghelper.LogError("VerifyOTP FAIL TO FETCH DATA : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERRORCODE_NO USER FOUND ")
	}
	if OTPstatus {
		requestClientId := c.Request().Header.Get(constants.CLIENTID)
		clientConfigForRequest, err := clientConfiguration.GetClientConfigeDAO(requestClientId)
		if nil != err {
			return c.JSON(http.StatusBadRequest, "Failed to fetch config")
		}
		login := model.Login{}
		login.Username = OTPval.Username
		resultObj1, err := GetUserQRcodeDAO(login.Username)
		if err != nil {
			logginghelper.LogError("ValidateCredentialsRoute FAIL TO FETCH DATA : ", err)
			return c.JSON(http.StatusExpectationFailed, "USER NOT FOUND ")
		}
		if clientConfigForRequest.Purpose.Settings.TwoStepAuth.Set {
			restrictJWTToken, err := GetTokenRestrictedurlService(login, requestClientId)
			if err != nil {
				logginghelper.LogError(" GetTokenRestrictedurlService GenerateToken : ", err)
				return c.JSON(http.StatusExpectationFailed, "GET TOKEN FAILED")
			}
			c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+restrictJWTToken)
			logginghelper.LogInfo(restrictJWTToken)
			return c.JSON(http.StatusOK, resultObj1.Google_Auth.QRcodeScan)
		}
		if redisSessionManager.IsSessionLimitReached(OTPval.Username) {
			return c.JSON(http.StatusAlreadyReported, "MAX_SESSION_LIMIT_REACHED")
		}

		secretKey := clientConfigForRequest.Jwt

		jwtTOKEN, err := GetToken(login, requestClientId, secretKey)

		if err != nil {
			logginghelper.LogError("ValidateCredentialsRoute FAIL TO GENERATE TOKEN : ", err)
			return c.JSON(http.StatusExpectationFailed, "FAIL TO GENERATE TOKEN ")
		}

		c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+jwtTOKEN)
		CreateActivityInfo(OTPval.Username)
		logginghelper.LogDebug("OUT: VerifyOTP")
		return c.JSON(http.StatusOK, "SUCCESS")

	}
	return c.JSON(http.StatusExpectationFailed, "OTP VERIFICATION FAILED")
}

func CreateActivityInfo(username string) {
	ActivityInfo := model.ActivityLog{}
	ActivityInfo.Username = username
	ActivityInfo.ActivitType = "LOGGED IN "
	ActivityInfo.ActivityResult = "SUCCESS"
	ActivityInfo.ActivityBy = "USER"
	ActivityInfo.ActivityOn = time.Now().Format(time.RFC850)
	activitylogged, _ := ActivityloggedService(ActivityInfo)
	logginghelper.LogInfo("ACTIVITY LOGGED:", activitylogged)
}

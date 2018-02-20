package forgotPassword

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/modules/login"
	"LoginProject/server/api/modules/otp"
	"errors"
	"fmt"
	"net/http"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
)

//Init method
func Init(o *echo.Group, r *echo.Group, c *echo.Group) {
	//For Open API
	o.POST("/forgotpassword/sendotp", SendOTPRoute)
	o.POST("/forgotpassword/verifyotp", VerifyOTPRoute)
	o.POST("/forgotpassword/changepassword", ChangePasswordRoute)
	o.POST("/forgotpassword/resendotp", ResendOTPRoute)
}

func SendOTPRoute(c echo.Context) error {
	logginghelper.LogDebug("SendOTPRoute() Start")
	profile := model.ProfileDetail{}

	err := c.Bind(&profile)

	if err != nil {
		logginghelper.LogError("SendOTPRoute: ", err)
		return c.JSON(http.StatusExpectationFailed, errors.New("PARAMETERS_BIND_ERROR"))
	}
	msg := fmt.Sprintf("Hi %s, your OTP for Changing Password is : ", profile.PersonalDetails.FullName)
	result, isUserPresent := otp.SendOTPService(profile, 3, msg)
	if !isUserPresent {
		logginghelper.LogError("SendOTPService: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	data := result.UserName
	logginghelper.LogDebug("SendOTPRoute() End")
	return c.JSON(http.StatusOK, data)
}

//ChangePasswordRoute Changes changed password
func ChangePasswordRoute(c echo.Context) error {
	logginghelper.LogDebug("IN: ChangePasswordRoute")
	loginObj := model.Login{}
	err := c.Bind(&loginObj)
	if err != nil {
		logginghelper.LogError("ChangePasswordRoute Bind Error", err)
		return c.JSON(http.StatusExpectationFailed, "PARAMETER_BIND_ERROR")
	}

	if loginObj.Password == "" || loginObj.ConfirmPassword == "" {
		logginghelper.LogError("Required params are empty")
		return c.JSON(http.StatusExpectationFailed, "REQUIRED_FIELD_VALIDATION_FAILED")
	} else if loginObj.Password != loginObj.ConfirmPassword {
		logginghelper.LogError("New password and confirm password does not match")
		return c.JSON(http.StatusExpectationFailed, "NEW_PASSWORD_AND_CONFIRM_PASSWORD_MISMATCH")
	}
	tokenClaims, err := login.GetDecodedTokenRestrictedurlService(c)
	if err != nil {
		logginghelper.LogError("ChangePassword GetDecodedLoginFromToken : ", err)
		return c.JSON(http.StatusExpectationFailed, "ChangePassword GetDecodedLoginFromToken Error")
	}
	username := tokenClaims["username"].(string)
	logginghelper.LogInfo("Username = ", username)
	loginObj.Username = username
	result, err := ChangePasswordService(loginObj)
	if err != nil {
		logginghelper.LogError("ChangePasswordService()", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	CreateActivityInfo(loginObj.Username)
	logginghelper.LogDebug("OUT: ChangePasswordRoute")

	return c.JSON(http.StatusOK, result)
}

func VerifyOTPRoute(c echo.Context) error {
	logginghelper.LogDebug("VerifyOTPRoute() Start")
	otpvalues := model.OTP{}
	user := model.Login{}

	err := c.Bind(&otpvalues)
	if err != nil {
		logginghelper.LogError("VerifyOTPRoute: ", err)
		return c.JSON(http.StatusExpectationFailed, errors.New("PARAMETERS_BIND_ERROR"))
	}
	user.Username = otpvalues.Username
	requestClientId := c.Request().Header.Get("ClientId")
	restrictJWTToken, err := login.GetTokenRestrictedurlService(user, requestClientId)
	if err != nil {
		logginghelper.LogError(" GetTokenRestrictedurlService GenerateToken : ", err)
		return c.JSON(http.StatusExpectationFailed, "GET TOKEN FAILED")
	}
	c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+restrictJWTToken)
	logginghelper.LogInfo("RestrictJWTtoken = ", restrictJWTToken)
	result, err := VerifyOTPService(otpvalues)
	if err != nil {
		logginghelper.LogError("VerifyOTPService: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logginghelper.LogDebug("VerifyOTPRoute() End")
	return c.JSON(http.StatusOK, result)
}

func ResendOTPRoute(c echo.Context) error {
	logginghelper.LogDebug("ResendOTPRoute() Start")
	profile := model.ProfileDetail{}

	err := c.Bind(&profile)

	if err != nil {
		logginghelper.LogError("ResendOTPRoute: ", err)
		return c.JSON(http.StatusExpectationFailed, errors.New("PARAMETERS_BIND_ERROR"))
	}
	result, status := ResendOTPService(profile)
	if !status {
		logginghelper.LogError("ResendOTPService: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logginghelper.LogDebug("ResendOTPRoute() End")
	return c.JSON(http.StatusOK, result)
}
func CreateActivityInfo(username string) {
	ActivityInfo := model.ActivityLog{}
	ActivityInfo.Username = username
	ActivityInfo.ActivitType = "PASSWORD CHANGED "
	ActivityInfo.ActivityResult = "SUCCESS"
	ActivityInfo.ActivityBy = "USER"
	ActivityInfo.ActivityOn = time.Now().Format(time.RFC850)
	activitylogged, _ := ActivityloggedService(ActivityInfo)
	logginghelper.LogInfo("ACTIVITY LOGGED:", activitylogged)
}

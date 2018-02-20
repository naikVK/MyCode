package registration

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/modules/clientConfiguration"
	"net/http"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
)

//Init method
func Init(o *echo.Group, r *echo.Group, c *echo.Group) {
	//For Open API
	o.POST("/register", RegisterRoute)
	o.POST("/isUsernameAvailable", IsUsernameAvailable)
}

func IsUsernameAvailable(c echo.Context) error {
	logginghelper.LogInfo("Inside registrationRoute:: IsUsernameAvailable")

	profileDetail := model.ProfileDetail{}
	err := c.Bind(&profileDetail)
	if err != nil {
		logginghelper.LogError("BIND_ERROR", err)
		return c.JSON(http.StatusExpectationFailed, "PARAMETER_DOES_NOT_GET_BIND_ERROR")
	}

	availableUsernames, isAvailable, err := IsUsernameAvailableService(profileDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not Available")
	}
	if isAvailable {
		return c.JSON(http.StatusOK, "Available")
	} else {
		return c.JSON(http.StatusAlreadyReported, model.UsernameSuggestion{availableUsernames})
	}
}

func RegisterRoute(c echo.Context) error {
	logginghelper.LogInfo("Inside registrationRoute:: RegisterRoute")
	profileDetail := model.ProfileDetail{}
	c.Bind(&profileDetail)
	logginghelper.LogInfo(profileDetail)

	err := RegisterUser(profileDetail)

	if nil != err {
		return c.JSON(http.StatusBadRequest, err)
	}
	//REVIEW: Use ClientId from constants.
	requestClientId := c.Request().Header.Get("ClientId")
	clientConfigForRequest, err := clientConfiguration.GetClientConfigeDAO(requestClientId)
	if nil != err {
		return c.JSON(http.StatusAlreadyReported, "Failed to find Config")
	}

	if clientConfigForRequest.Purpose.Settings.RegistrationNotification.SMS {
		sendmsg, _ := SendSuccessMsgPhoneService(profileDetail)
		logginghelper.LogInfo("Registration sms send", sendmsg)
	}
	if clientConfigForRequest.Purpose.Settings.RegistrationNotification.Email {
		sendmsg, _ := SendSuccessMsgEmailService(profileDetail)
		logginghelper.LogInfo("Registration email send", sendmsg)
	}
	CreateActivityInfo(profileDetail.UserName)
	return c.JSON(http.StatusOK, "Success")
}

func CreateActivityInfo(username string) {
	ActivityInfo := model.ActivityLog{}
	ActivityInfo.Username = username
	ActivityInfo.ActivitType = "USER REGISTERED "
	ActivityInfo.ActivityResult = "SUCCESS"
	ActivityInfo.ActivityBy = "USER"
	ActivityInfo.ActivityOn = time.Now().Format(time.RFC850)
	activitylogged, _ := ActivityloggedService(ActivityInfo)
	logginghelper.LogInfo("ACTIVITY LOGGED:", activitylogged)
}

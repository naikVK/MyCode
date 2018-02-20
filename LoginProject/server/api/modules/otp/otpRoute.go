package otp

import (
	"LoginProject/server/api/common/model"
	"net/http"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"

	"github.com/labstack/echo"
)

//INIT METHOD
func Init(o *echo.Group, r *echo.Group, c *echo.Group) {
	//For Open API
	o.POST("/verifyotp", VerifyOTPRoute)
}

//VERIFY OTP
func VerifyOTPRoute(c echo.Context) error {
	logginghelper.LogInfo("Inside otpRoute: VerifyOTPRoute")
	otpModel := model.OTP{}
	err := c.Bind(&otpModel)
	if err != nil {
		logginghelper.LogError("BIND ERROR :", err)
		return c.JSON(http.StatusExpectationFailed, "bind error")
	}
	isVerified, err := VerifyOTPDAO(otpModel)
	if err != nil {
		logginghelper.LogError("OTP VERIFICATION FAILED :", err)
		return c.JSON(http.StatusExpectationFailed, "Failed")
	}
	return c.JSON(http.StatusOK, isVerified)
}

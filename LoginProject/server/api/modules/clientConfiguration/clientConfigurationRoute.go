package clientConfiguration

import (
	"LoginProject/server/api/common/model"
	"net/http"

	"strings"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
)

//Init method
func Init(o *echo.Group, r *echo.Group, c *echo.Group) {
	//For Open API
	o.POST("/getclientconfig", GetClientConfigRoute)
}

//GET CLIENT CONFIGURATION FROM DB
func GetClientConfigRoute(c echo.Context) error {
	logginghelper.LogDebug("IN: GetClientConfigRoute")
	client := model.Client{}
	err := c.Bind(&client)
	if err != nil {
		logginghelper.LogError("GetClientConfigRoute Bind : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERA_ERRORCODE_PARAMETER_BIND_ERROR")
	}
	if strings.Trim(client.ClientId, "") == "" {
		logginghelper.LogError("GetLearnerProfileRoute USERNAME OR PASSWORD is empty")
		return c.JSON(http.StatusExpectationFailed, "ERA_ERRORCODE_REQUIRED_FIELD_VALIDATION_FAILED")
	}
	resultObj, err := GetClientConfigService(client)
	logginghelper.LogInfo(resultObj)
	if err != nil {
		logginghelper.LogError("GetClientConfigRoute FAIL TO FETCH DATA : ", err)
		return c.JSON(http.StatusExpectationFailed, "ERA_ERRORCODE_NO CLIENT NOT  FOUND ")
	}
	return c.JSON(http.StatusOK, resultObj)
}

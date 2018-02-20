package api

import (
	"LoginProject/server/api/modules/forgotPassword"
	"LoginProject/server/api/modules/login"
	"LoginProject/server/api/modules/otp"
	"LoginProject/server/middleware"
	"LoginProject/server/redisSessionManager"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"

	"LoginProject/server/api/modules/clientConfiguration"
	"LoginProject/server/api/modules/registration"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	o := e.Group("/o")
	r := e.Group("/r")
	c := r.Group("/c")

	// applying middleware
	middleware.Init(e, o, r, c)
	// initialize session store
	redisSessionManager.Init()
	// initialize api
	registration.Init(o, r, c)
	login.Init(o, r, c)
	confighelper.InitViper()
	clientConfiguration.Init(o, r, c)
	forgotPassword.Init(o, r, c)
	otp.Init(o, r, c)
}

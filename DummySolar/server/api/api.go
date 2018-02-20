package api

import (
	"DummySolar/server/middleware"
	"DummySolar/server/redisSessionManager"

	"DummySolar/server/api/modules/profileModule"
	// "corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
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
	profileModule.Init(o, r, c)
}

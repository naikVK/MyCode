package main

import (
	"LoginProject/server/api"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
)

func main() {

	// Create a new instance of Echo
	e := echo.New()
	// Starting Logger
	logginghelper.Init("logs/loginproject.log", false, 0, 0, 0, true)
	// initializing all routes
	api.Init(e)
	// Start as a web server
	serverport := confighelper.GetConfig("app_port")
	e.Start(":" + serverport)
}

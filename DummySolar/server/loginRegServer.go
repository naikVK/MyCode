package main

import (
	"DummySolar/server/api"

	"github.com/labstack/echo"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()
	confighelper.InitViper()
	// initializing all routes
	api.Init(e)
	// Start as a web server
	e.Start(":3032")

}

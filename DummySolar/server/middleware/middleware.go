package middleware

import (
	"net/http"
	// confighelper "DummySolar/server/api/common/utils/configHelper"
	"DummySolar/server/api/common/utils/jwtUtils"
	"DummySolar/server/redisSessionManager"
	// "strings"
	// "time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"

	logginghelper "corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Init middleware
func Init(e *echo.Echo, o *echo.Group, r *echo.Group, c *echo.Group) {

	//FIXME: Can we change below URLS?
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://10.4.1.207:8082", "http://10.4.1.186:8080", "http://localhost"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.HEAD, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human} ${header:all} ${header:*} ${header} ${form} \n",
	}))
	//NOTE: STATTIC IMAGES

	r.Use(JwtMiddleware())
}
// This middleware will be called for every restricted URL
func JwtMiddleware() echo.MiddlewareFunc {
	logginghelper.LogInfo("JwtMiddleware called.............. ")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// tokenFromRequest := c.Request().Header.Get("authorization")
			requestClientId := c.Request().Header.Get("ClientId")
		
			logginghelper.LogDebug("client id" + requestClientId)
			// tokenCookie, errc := c.Cookie(requestClientId)
			
			// if errc != nil {
			// 	logginghelper.LogError("error occured while fetching token from request cookie")
			// 	return echo.ErrUnauthorized
			// }
			// tokenFromCookie := tokenCookie.Value
			// logginghelper.LogDebug("token", tokenFromCookie)
			// logginghelper.LogInfo("token received from request : ", tokenFromCookie)
		
						// requestClientId := c.Request().Header.Get("ClientId")
			// clientConfig, err := clientConfiguration.GetGetClientConfigeDAO(requestClientId)

			serverClientId := confighelper.GetConfig("serverConfig.serverForClient")
			secretKey := confighelper.GetConfig("serverConfig.jwtSecretKey")

			if requestClientId != serverClientId {
				logginghelper.LogError("client is not authorized")
				return echo.ErrUnauthorized
			}
			clientId := c.Request().Header.Get("ClientId")
			tokenCookie, errc := c.Cookie(clientId)
			if nil != errc {
				logginghelper.LogError("Error: ", errc)
				return echo.ErrUnauthorized
			}
			token:= tokenCookie.Value

			tokenMap, tokenError := jwtUtils.GetDecodedLoginFromToken(c, secretKey)
			if tokenError != nil {
				logginghelper.LogError("error occured while calling GetDecodedLoginFromToken ", tokenError)
				return echo.ErrUnauthorized
			}

			sessionKey, _ := tokenMap["sessionId"].(string)
			// check token in session store
			gcActiveToken, cacheError := redisSessionManager.Get(sessionKey)
			logginghelper.LogInfo("tokens")
			logginghelper.LogInfo(gcActiveToken)
			logginghelper.LogInfo(token)
			if tokenError != nil {
				logginghelper.LogError("error occured while calling GetDecodedLoginFromToken ", tokenError)
				return echo.ErrUnauthorized
			}

			// if failed to access session
			if cacheError != nil || gcActiveToken == nil || gcActiveToken == "" {
				return c.JSON(http.StatusUnauthorized, "SESSION_EXPIRED")
			}
			cacheToken, _ := gcActiveToken.(string)
			cacheToken = "Bearer " + cacheToken
			// if tokenFromCookie dont match with session token
			if token != cacheToken {
				logginghelper.LogInfo("activation token from request is not matching with activation token from redis store")
				return c.JSON(http.StatusUnauthorized, "SESSION_EXPIRED")
			} else {
				// token from request is equal to token from redis sessionStore token
				// sliding session
				redisSessionManager.SlideSession(sessionKey)
				return next(c)
			}
		}
	}
}

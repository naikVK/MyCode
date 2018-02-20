package middleware

import (
	"LoginProject/server/api/common/constants"

	// confighelper "LoginProject/server/api/common/utils/configHelper"
	"LoginProject/server/api/common/utils/jwtUtils"
	"LoginProject/server/api/modules/clientConfiguration"
	"LoginProject/server/redisSessionManager"
	"net/http"
	"os"
	"strings"
	"time"

	logginghelper "corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Init middleware
func Init(e *echo.Echo, o *echo.Group, r *echo.Group, c *echo.Group) {
	path := getPath()
	path = path + constants.QRCODE_PATH
	//FIXME: Can we change below URLS?
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*", "http://10.4.1.186:8082", "http://10.4.1.186:8083", "http://10.4.1.186:8084", "localhost"},
		AllowMethods:  []string{echo.GET, echo.PUT, echo.HEAD, echo.PATCH, echo.POST, echo.DELETE},
		ExposeHeaders: []string{echo.HeaderAuthorization},
		// AllowCredentials: true,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human} ${header:all} ${header:*} ${header} ${form} \n",
	}))
	//NOTE: STATTIC IMAGES
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		// path := getPath()
		Root: "D:/resources",
		Browse: false,
	}))
	r.Use(JwtMiddleware())
}

// This middleware will be called for every restricted URL
func JwtMiddleware() echo.MiddlewareFunc {
	logginghelper.LogInfo("JwtMiddleware called.............. ")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// tokenFromRequest := c.Request().Header.Get("authorization")
			requestClientId := c.Request().Header.Get("ClientId")
			if requestClientId == "" || len(requestClientId) == 0 {
				return c.JSON(http.StatusBadRequest, "Header clientId not present: "+requestClientId)
			}
			clientConfig, err := clientConfiguration.GetClientConfigeDAO(requestClientId)

			if nil != err {
				return c.JSON(http.StatusBadRequest, "config not found for client: "+requestClientId)
			}
			tokenFromRequest := c.Request().Header.Get("Authorization")

			if tokenFromRequest == "" {
				logginghelper.LogError("error occured while fetching token from request header")
				return echo.ErrUnauthorized
			}

			tokenArray := strings.Split(tokenFromRequest, "Bearer")

			if len(tokenArray) <= 1 {
				logginghelper.LogError("error occured while splitting token")
				return echo.ErrUnauthorized
			}

			tokenFromRequest = strings.Trim(tokenArray[1], " ")
			time.Sleep(100)

			secretKey := clientConfig.Jwt
			tokenMap, tokenError := jwtUtils.GetDecodedLoginFromToken(c, secretKey)
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
			if tokenFromRequest != gcActiveToken {
				logginghelper.LogInfo("activation token from request is not matching with activation token from gcache")
				return c.JSON(http.StatusAlreadyReported, "SESSION_EXPIRED")
			} else {
				// token from request is equal to token from redis sessionStore token
				// sliding session
				redisSessionManager.SlideSession(sessionKey)
				return next(c)
			}
		}
	}
}
func getPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		logginghelper.LogError(err)
	}
	return pwd
}

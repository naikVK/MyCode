package jwtUtils

import (
	"DummySolar/server/api/common/constants"
	"DummySolar/server/api/common/model"
	"fmt"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/authhelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	logger "corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateJwtToken(login model.Login, sessionKeyPostfix, clientId, secretKey string) (string, error) {
	// token := jwt.New(jwt.SigningMethodHS256)
	logger.LogInfo("Inside: generateJwtToken")
	expAt := time.Now().Add(constants.JWT_TOKEN_EXPAT).Unix()
	sessionId := login.Username + ":" + sessionKeyPostfix
	fmt.Println(sessionId)

	// expireToken := time.Now().Add(time.Hour * 48).Unix()
	claims := model.JwtCustomClaims{
		login.Username,
		false,
		sessionId,
		clientId,
		jwt.StandardClaims{
			ExpiresAt: expAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if nil != err {
		logginghelper.LogError("GenerateToken SignedString() Error: ", err)
		return "", nil
	}
	return tokenString, nil
}

func GetDecodedLoginFromToken(c echo.Context, secretKey string) (jwt.MapClaims, error) {
	logginghelper.LogInfo("Inside jwtUtils: GetDecodedLoginFromToken")
	logginghelper.LogInfo("secreet key" + secretKey)
	clientId := c.Request().Header.Get("ClientId")
	tokenCookie, errc := c.Cookie(clientId)
	if nil != errc {
		logginghelper.LogError("GetDecodedLoginFromToken DecodeToken() Error: ", errc)
		return jwt.MapClaims{}, errc
	}
	token:= tokenCookie.Value
	decodedTokenClaims, err := authhelper.DecodeToken(token, secretKey)
	if err != nil {
		logginghelper.LogError("GetDecodedLoginFromToken DecodeToken() Error: ", err)
		return jwt.MapClaims{}, err
	}
	return decodedTokenClaims, nil
}

func ValidateToken(token, secretKey string) (jwt.MapClaims, error) {
	logginghelper.LogDebug("Inside jwtUtils:: ValidateToken")
	decodedTokenClaims, err := authhelper.DecodeToken(token, secretKey)
	if err != nil {
		logginghelper.LogError("ValidateToken DecodeToken() Error: ", err)
		return jwt.MapClaims{}, err
	}
	return decodedTokenClaims, nil
}
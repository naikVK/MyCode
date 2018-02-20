package google_auth

import (
	"LoginProject/server/api/common/constants"
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils"
	"LoginProject/server/api/common/utils/connect_db"
	"crypto/rand"
	"encoding/base32"
	"net/url"
	"os"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/filehelper"

	dgoogauth "github.com/dgryski/dgoogauth"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	qr "github.com/rsc/qr"
	"gopkg.in/mgo.v2/bson"
)

//create qr code for scanning
func CreateQRCode(username string) (string, error) {
	logginghelper.LogDebug("IN: CreateQRCode")
	login := model.Login{}
	login.Username = username
	secret := make([]byte, 6)
	_, err := rand.Read(secret)
	if err != nil {
		logginghelper.LogDebug(err)
	}

	secret = append(secret, []byte{0xDE, 0xAD, 0xBE, 0xEF}...)
	secretBase32 := base32.StdEncoding.EncodeToString(secret)
	account := confighelper.GetConfig("GOOGLE_AUTH.account")
	issuer := confighelper.GetConfig("GOOGLE_AUTH.issuer")

	URL, err := url.Parse(constants.GOOGLE_AUTH_URL)
	if err != nil {
		logginghelper.LogDebug(err)
	}

	URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)

	params := url.Values{}
	params.Add("secret", secretBase32)
	params.Add("issuer", issuer)

	URL.RawQuery = params.Encode()
	logginghelper.LogInfo("URL is %s\n", URL.String())

	code, err := qr.Encode(URL.String(), qr.Q)
	if err != nil {
		logginghelper.LogDebug(err)
	}
	b := code.PNG()
	filename := utils.GetRandamStringOfLen(4) + constants.QRCODE_EXT
	projectPath := getPath()
	QR_CODE_IMAGE_PATH := projectPath + constants.QRCODE_PATH
	qrFilename := QR_CODE_IMAGE_PATH + filename
	err = filehelper.WriteFile(qrFilename, b)
	if err != nil {
		logginghelper.LogError(err)
		return "", err
	}
	status := UpdateSecretKey(login, secretBase32)
	logginghelper.LogInfo("QR code is in %s. Please scan it into Google Authenticator app.\n", qrFilename)
	logginghelper.LogDebug("OUT: CreateQRCode")
	if status {
		return filename, nil
	}
	return "", err
}

//Authenticate totp of google authenticator
func AuthenticateOTP(Auth_check model.Google_AuthCheck) bool {
	logginghelper.LogDebug("IN: AuthenticateOTP")
	secretBase32, err := GetSecretKey(Auth_check.UserName)
	if err != nil {
		logginghelper.LogError("AuthenticateOTP did not get secret key", err)
		return false
	}
	if secretBase32 != "" {
		otpc := &dgoogauth.OTPConfig{
			Secret:      secretBase32,
			WindowSize:  3,
			HotpCounter: 0,
			// UTC:         true,
		}
		var token string
		token = Auth_check.OTP
		if token == "q" {
			logginghelper.LogInfo("BREAK")
		}

		val, err := otpc.Authenticate(token)
		if err != nil {
			logginghelper.LogError("google authentication failed", err)
		}
		logginghelper.LogDebug("OUT: AuthenticateOTP")
		if !val {
			logginghelper.LogInfo("Sorry, Not Authenticated")
			return false
		}

		logginghelper.LogInfo("Authenticated!")
		return true
	}
	logginghelper.LogDebug("OUT: AuthenticateOTP")
	return false
}

//upadate secret key after qr code generation
func UpdateSecretKey(userObject model.Login, key string) bool {
	logginghelper.LogDebug("IN: UpdateSecretKey")
	var result model.Login
	dbStatus, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError("UpdateSecretKey UNABLE TO CONNECT TO DB Error : ", err)
		return false
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	cd := dbStatus.DB(name).C(collection_name)

	Ferr := cd.Find(bson.M{
		"USERNAME": userObject.Username}).One(&result)
	if Ferr != nil {
		logginghelper.LogError("UpdateSecretKey USER NOT PRESENT IN DB Error : ", Ferr)
		return false
	}
	result.Google_Auth.SecretKey = key
	logginghelper.LogInfo("secret key:", result.Google_Auth.SecretKey)
	Serr := cd.Update(bson.M{
		"USERNAME": userObject.Username}, result)
	if Serr != nil {
		logginghelper.LogError("UpdateSecretKey USER NOT PRESENT IN DB Error : ", Ferr)
		return false
	}
	logginghelper.LogDebug("OUT: UpdateSecretKey")

	return true
}

//get secret key from DB
func GetSecretKey(username string) (string, error) {
	logginghelper.LogDebug("IN: GetSecretKey")
	var result model.Login
	dbStatus, err := connect_db.ConnecttoMongoDB()
	if err != nil {
		logginghelper.LogError("GetSecretKey UNABLE TO CONNECT TO DB Error : ", err)
		return "", err
	}
	name := confighelper.GetConfig("mongo.dbname")
	collection_name := confighelper.GetConfig("mongo.collection_logindetails")
	cd := dbStatus.DB(name).C(collection_name)

	Ferr := cd.Find(bson.M{
		"USERNAME": username}).One(&result)
	if Ferr != nil {
		logginghelper.LogError("GetSecretKey USER NOT PRESENT IN DB Error : ", Ferr)
		return "", Ferr
	}
	logginghelper.LogDebug("OUT: GetSecretKey")
	return result.Google_Auth.SecretKey, nil
}

//get path f project to set qrcode img path
func getPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		logginghelper.LogError("COULD NOT GET PROJECT PATH", err)
	}
	return pwd
}

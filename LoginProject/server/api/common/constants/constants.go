package constants

import (
	"time"
)

const (
	No_OF_USERNAME_SUGGESSION = 2

	REDIS_CONNECTION_ADDR       = "localhost:6379"
	DEFAULT_SESSION_EXPIRE_TIME = time.Minute * 30
	MAX_SESSION_ALLOWED         = 10

	JWT_TOKEN_EXPAT = time.Minute * 30
	KEYLEN          = 8
	GOOGLE_AUTH_URL = "otpauth://totp"
	QRCODE_EXT      = ".png"

	CAPTCHA_API              = "https://www.google.com/recaptcha/api/siteverify"
	SECRET_KEY_CAPTCHA       = "6Lf-hzoUAAAAALiYXYaW-8fLuWkpw4FSyyzdW9eX"
	CLIENTID                 = "ClientId"
	RESTRICTKEY_PATH         = "RESTRICT_KEY.URLkey"
	ENCRYPT_DECRYPT_KEY_PATH = "ENCRYPT_DECRYPT.key"
	QRCODE_PATH              = "/server/cdn/qr_code/"
	MSG_GOAUTH_SECRET_KEY    = "Your Secret Key for Google Authenticator is :  "
	MSG_OTP_LOGIN            = "Hi YOUR OTP FOR lOGIN IS : "
	FROM_ADDRESS             = "priyankam@mkcl.org"
)

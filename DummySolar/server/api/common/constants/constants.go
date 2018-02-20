package constants

import "time"

const (
	DB_CONNECTION_ADD        = "10.2.10.160:27017"
	DB_NAME                  = "Login_RegDB"
	COLLECTION_REGISTRATION  = "PROFILE_DETAILS"
	COLLECTION_LOGIN_DETAILS = "LOGIN_DETAILS"
	COLLECTION_OTPVERIFY     = "OTPVERIFY"

	No_OF_USERNAME_SUGGESSION = 5

	REDIS_CONNECTION_ADDR       = "localhost:6379"
	DEFAULT_SESSION_EXPIRE_TIME = time.Minute * 30
	MAX_SESSION_ALLOWED         = 2

	JWT_SECRET_KEY  = "secretkey"
	JWT_TOKEN_EXPAT = time.Minute * 30

	CAPTCHA_API        = "https://www.google.com/recaptcha/api/siteverify"
	SECRET_KEY_CAPTCHA = "6Lf9KjkUAAAAAJOEkkMStE-A8AiaCKl59Usk6ghr"
	//NOTE: Change according to ur filesystem
	QR_CODE_IMAGE_PATH = "/home/administrator/go/src/LoginProject/server/cdn/qr_code/"

	TOKEN_VALIDATION_API = "http://10.4.1.186:3030/o/isValidToken"
)

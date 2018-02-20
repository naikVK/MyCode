package model

// Login keeps information of all Logins
type Login struct {
	Username string `json:"username,omitempty" bson:"USERNAME" validate:"required"`

	Password        string `json:"password,omitempty" bson:"PASSWORD" validate:"required"`
	ConfirmPassword string `json:"confirmpassword,omitempty" validate:"required"`

	Captcha         string      `json:"captchaResponse,omitempty"`
	CreatedOn       string      `json:"createdon,omitempty" bson:"CREATED_ON"`
	IsActive        string      `json:"isactive,omitempty" bson:"ISACTIVE"`
	LoginExpiresOn  string      `json:"loginexpireson,omitempty" bson:"LOGIN_EXPIRES_ON"`
	LoginModifiedOn string      `json:"loginmodifiedon,omitempty" bson:"LOGIN_MODIFIED_ON"`
	Google_Auth     Google_Auth `json:"google_auth,omitempty" bson:"GOOGLE_AUTH"`

	Hash string `json:"hash,omitempty"`
}

type Google_AuthCheck struct {
	UserName string `json:"username"`
	OTP      string `json:"otp"`
}
type CaptchResponse struct {
	IsSuccess bool   `json:"success"`
	HostName  string `json:"hostname"`
}
type Google_Auth struct {
	QRcodeScan bool   `json:"qrcodescan,omitempty" bson:"QR_CODE_SCAN"`
	SecretKey  string `json:"secretkey,omitempty" bson:"SECRET_KEY"`
}

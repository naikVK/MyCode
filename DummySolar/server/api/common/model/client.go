package model

type Client struct {
	Id         string  `json:"id" bson:"_id"`
	ClientName string  `json:"clientname" bson:"CLIENT_NAME"`
	ClientId   string  `json:"clientid" bson:"CLIENT_ID" validate:"required"`
	ClientCode string  `json:"clientcode" bson:"CLIENT_CODE"`
	ExpiresOn  string  `json:"expireson" bson:"EXPIRES_ON"`
	Group      Group   `json:"group" bson:"GROUP"`
	Jwt        string  `json:"jwt" bson:"JWT"`
	Purpose    Purpose `json:"purpose" bson:"PURPOSE"`
}
type Group struct {
	GroupCode        string   `json:"groupcode" bson:"GROUP_CODE"`
	GroupDescription string   `json:"groupdescription" bson:"GROUP_DESCRIPTION"`
	GroupMembers     []string `json:"groupmembers" bson:"GROUP_MEMBERS"`
	GroupName        string   `json:"groupname" bson:"GROUP_NAME"`
}
type Purpose struct {
	Settings    Settings    `json:"settings" bson:"SETTINGS"`
	ServiceType Servicetype `json:"servicetype" bson:"SERVICE_TYPE"`
	ReturnURL   string      `json:"returnurl" bson:"RETURN_URL"`
}
type Settings struct {
	MultipleLogin     bool        `json:"multiplelogin" bson:"MULTIPLE_LOGIN"`
	MultipleSession   bool        `json:"multiplesession" bson:"MULITPLE_SESSION"`
	TwoStepAuth       TwoStepAuth `json:"twostepauth" bson:"TWO_STEP_AUTHENTICATION"`
	Captcha           bool        `json:"captcha" bson:"CAPTCHA"`
	EmailNotification bool        `json:"emailnotification" bson:"EMAIL_NOTIFICATION"`
	OtpRequired       OtpRequired `json:"otprequired" bson:"OTP_REQUIRED"`
}
type TwoStepAuth struct {
	Set  bool  `json:"set" bson:"SET"`
	Type Type1 `json:"type1" bson:"TYPE1"`
}
type Type1 struct {
	Email                bool `json:"email" bson:"EMAIL"`
	SMS                  bool `json:"sms" bson:"SMS"`
	Google_Authenticator bool `json:"google_authenticator" bson:"GOOGLE_AUTHENTICATOR"`
}
type OtpRequired struct {
	Set  bool  `json:"set" bson:"SET"`
	Type Type2 `json:"type2" bson:"TYPE2"`
}
type Type2 struct {
	Email  bool `json:"email" bson:"EMAIL"`
	Mobile bool `json:"mobile" bson:"MOBILE"`
}
type Servicetype struct {
	API   bool `json:"api" bson:"API"`
	UI    bool `json:"ui" bson:"UI"`
	EMBED bool `json:"embed" bson:"EMBED"`
}

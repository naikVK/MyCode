package model

import (
	"gopkg.in/mgo.v2/bson"
)

type UsernameSuggestion struct {
	Usernames []string `json:"usernames"`
}
type AccountSetting struct {
	NotificationAllow    bool   `json:"notificationAllow" bson:"NOTIFICATION_ALLOW"`
	OptedOutPermanently  bool   `json:"optedOutPermanently" bson:"OPTED_OUT_PERMANENT"`
	OptedOutTemporary    bool   `json:"optedOutTemporary" bson:"OPTED_OUT_TEMPORARY"`
	PasswordChangedCount string `json:"passwordChangedCount" bson:"PASSWORD_CHANGED_COUNT"`
	PasswordChangedOn    string `json:"passwordChangedOn" bson:"PASSWORD_CHANGED_ON"`
}

type ContactDetails struct {
	Email  Email  `json:"email" bson:"EMAIL"`
	Mobile Mobile `json:"mobile" bson:"MOBILE"`
}
type Email struct {
	Address    string `json:"address" bson:"ADDRESS"`
	VerifiedOn string `json:"verifiedOn" bson:"VERIFED_ON,omitempty"`
}
type Mobile struct {
	Number     string `json:"number" bson:"NUMBER"`
	VerifiedOn string `json:"verifiedOn" bson:"VERFIED_ON,omitempty"`
}

type PersonalDetails struct {
	Dob      string `json:"dob" bson:"DATE_OF_BIRTH"`
	FullName string `json:"fullName" bson:"FULLNAME"`
	Gender   string `json:"gender" bson:"GENDER"`
}
type ProfileDetail struct {
	ID              bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	AccountSetting  AccountSetting  `json:"accountSetting" bson:"ACCOUNT_SETTINGS"`
	ContactDetails  ContactDetails  `json:"contactDetails" bson:"CONTACT_DETAILS"`
	CreatedOn       string          `json:"createdOn" bson:"CREATED_ON"`
	ModifiedBy      string          `json:"modifiedBy" bson:"MODIFIED_BY"`
	ModifiedOn      string          `json:"modifiedOn" bson:"MODIFIED_ON"`
	PersonalDetails PersonalDetails `json:"personalDetails" bson:"PERSONAL_DETAILS"`
	UserName        string          `json:"userName" bson:"USERNAME"`
	Password        string          `json:"password"  bson:"-"`
	IsEmail         bool            `json:"isEmail,omitempty" bson:"ISEMAIL"`
}

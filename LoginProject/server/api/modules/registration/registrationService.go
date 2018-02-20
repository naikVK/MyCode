package registration

import (
	"LoginProject/server/api/common/constants"
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils/email"
	"LoginProject/server/api/common/utils/sms"
	"LoginProject/server/api/common/utils/validations"
	"LoginProject/server/api/modules/login"
	"fmt"
	"strings"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

//RegisterUser registers user onto the system
func RegisterUser(profileDetail model.ProfileDetail) error {
	logginghelper.LogInfo("Inside registrationService:: RegisterUser")
	if !isValideRegistrationDetails(profileDetail) {
		logginghelper.LogInfo("Inside registrationService:: Invalid registration Details")
		return fmt.Errorf("Invalid registration details")
	}
	username := strings.Trim(profileDetail.UserName, " ")

	// checking username availability
	// if user found with username then return err
	_, isUserAlreadyPresent := GetByUserName(username)
	if isUserAlreadyPresent {
		logginghelper.LogError("Username already exist")
		return fmt.Errorf("Username already exist")
	}

	err := Insert(profileDetail)
	if err != nil {
		logginghelper.LogError("ERROR WHILE INSERTING PROFILE DETAILS IN DATABASE", err)
		return err
	}

	loginDetail := model.Login{}
	loginDetail.Username = profileDetail.UserName
	loginDetail.Password = profileDetail.Password
	loginDetail.CreatedOn = time.Now().Format(time.RFC3339)
	logDberr := login.InsertLoginDetailsDAO(loginDetail)
	if nil != logDberr {
		logginghelper.LogError("ERROR WHILLE INSERTING LOGIN DATA IN DATABASE", logDberr)
		return logDberr
	}
	return nil
}

func IsUsernameAvailableService(profileDetail model.ProfileDetail) ([]string, bool, error) {
	logginghelper.LogInfo("Inside registrationService:: IsUsernameAvailableService")

	username := strings.Trim(profileDetail.UserName, " ")

	if validations.IsEmpty(username) {
		return nil, false, fmt.Errorf("Invalid data")
	}

	_, isUserPresent := GetByUserName(username)

	if !isUserPresent {
		return nil, true, nil
	}

	availableUSernames := GetUsernameSuggessions(profileDetail, constants.No_OF_USERNAME_SUGGESSION)
	return availableUSernames, false, nil
}

func isValideRegistrationDetails(profileDetails model.ProfileDetail) bool {
	logginghelper.LogInfo("Inside registrationService:: isValideRegistrationDetails")
	email := profileDetails.ContactDetails.Email.Address
	mobile := profileDetails.ContactDetails.Mobile.Number
	fullName := profileDetails.PersonalDetails.FullName
	userName := profileDetails.UserName
	gender := profileDetails.PersonalDetails.Gender
	dob := profileDetails.PersonalDetails.Dob
	//password := profileDetails.Password

	if validations.IsEmpty(email, mobile, fullName, userName, gender, dob) {
		logginghelper.LogInfo("BadRequest: required fields not provided")
		return false
	}

	if !validations.IsMobileNumber(mobile) {
		logginghelper.LogInfo("BadRequest: Invalid Mobile number")
		return false
	}

	if !validations.IsValidEmail(email) {
		logginghelper.LogInfo("BadRequest: Invalid Email address")
		return false
	}
	return true
}
func SendSuccessMsgPhoneService(user model.ProfileDetail) (bool, error) {
	logginghelper.LogDebug("IN: GetUserProfileService")
	msg := fmt.Sprintf("Hi %s, you have successfully Registered for MKCL login component", user.PersonalDetails.FullName)
	phoneNumber := user.ContactDetails.Mobile.Number
	Sendmsg, err := sms.SendSingleSMS(msg, phoneNumber, "")
	if err != nil {
		logginghelper.LogError("Error while sending registration sms")
		return false, err
	}
	if Sendmsg != "" {
		return false, nil
	}
	return true, nil
}
func SendSuccessMsgEmailService(user model.ProfileDetail) (bool, error) {
	logginghelper.LogDebug("IN: GetUserProfileService")
	msg := fmt.Sprintf("Hi %s, you have successfully Registered for MKCL login component", user.PersonalDetails.FullName)
	emailid := user.ContactDetails.Email.Address
	Sendmsg, err := email.Sendmail(emailid, msg, "")
	if err != nil {
		logginghelper.LogError("ERROR WHILE SENDING MESSAGE", err)
		return false, err
	}
	return Sendmsg, nil
}

//LOG THE ACTIVITY
func ActivityloggedService(AcitivityInfo model.ActivityLog) (bool, error) {
	ActivityLoggedstatus, err := login.ActivityloggedService(AcitivityInfo)
	if err != nil {
		logginghelper.LogError("ActivityloggedService:", err)
		return ActivityLoggedstatus, err
	}
	return ActivityLoggedstatus, nil
}

package otp

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils"
	"LoginProject/server/api/common/utils/email"
	"LoginProject/server/api/common/utils/fetchProfile"
	"LoginProject/server/api/common/utils/sms"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

//SEND OTP SERVICE
func SendOTPService(profileobj model.ProfileDetail, flag int, body string) (model.ProfileDetail, bool) {
	logginghelper.LogDebug("SendOTPService() Start: ")
	isEmail := profileobj.IsEmail
	var otp string
	var err error
	profileData, isUserPresent := fetchProfile.GetByUserName(profileobj.UserName)
	if !isUserPresent {
		logginghelper.LogError("SendOTPDAO() Error")
		return profileData, isUserPresent
	}
	//Flag = 1 will send otpmessage only on mobile
	if flag == 1 {
		otp, err = sms.SendSingleSMS(body, profileData.ContactDetails.Mobile.Number, "otpmsg")
		if err != nil {
			logginghelper.LogError("ERROR WHILE SENDING SMS", err)
			return model.ProfileDetail{}, false
		}
	}
	//Flag = 2 indicates that otpmessage will only be sent on email
	if flag == 2 {
		emailID := profileData.ContactDetails.Email.Address
		logginghelper.LogInfo(emailID)
		otp = utils.GetRandomOTP()
		emailotp, err := email.Sendmail(emailID, body, otp)
		if err != nil {
			logginghelper.LogError("ERROR WHILE SENDING EMAIL", err)
			return model.ProfileDetail{}, false
		}
		logginghelper.LogInfo(otp)
		logginghelper.LogInfo("emailotp=", emailotp)
	}
	//Flag = 3 indicates that otpmessage will be sent on mobile as well as email
	if flag == 3 {
		otp, err := sms.SendSingleSMS(body, profileData.ContactDetails.Mobile.Number, "otpmsg")
		if err != nil {
			logginghelper.LogError("ERROR WHILE SENDING SMS", err)
			return model.ProfileDetail{}, false
		}
		if isEmail {
			email.Sendmail(profileData.ContactDetails.Email.Address, body, otp)
		}
	}

	otpValues := model.OTP{}
	otpValues.Username = profileobj.UserName
	otpValues.OTP = otp
	status, err := InsertOTP(otpValues)
	if err != nil {
		logginghelper.LogError("Error while inserting otp data in database", err)
		return model.ProfileDetail{}, false
	}
	logginghelper.LogInfo("Status=", status)
	logginghelper.LogDebug("SendOTPService() Ended")
	return profileData, true
}

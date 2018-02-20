package forgotPassword

import (
	"LoginProject/server/api/common/model"
	"LoginProject/server/api/common/utils/email"
	"LoginProject/server/api/common/utils/sms"
	"LoginProject/server/api/modules/login"
	"LoginProject/server/api/modules/otp"
	"fmt"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

func VerifyOTPService(otpvalues1 model.OTP) (bool, error) {
	logginghelper.LogDebug("VerifyOTPService() Start: ")
	status, err := otp.VerifyOTPDAO(otpvalues1)
	if err != nil {
		logginghelper.LogError("VerifyOTPDAO() Error")
		return status, err
	}
	logginghelper.LogDebug("VerifyOTPService() Ended")
	return status, err
}

func ChangePasswordService(loginObj model.Login) (bool, error) {

	logginghelper.LogDebug("ChangePasswordService() Start: ")
	if loginObj.Password == loginObj.ConfirmPassword {
		status := login.UpdatePasswordService(loginObj)
		logginghelper.LogInfo(status)
	}
	logginghelper.LogDebug("ChangePasswordService() Ended")
	return true, nil
}

func ResendOTPService(profileobj model.ProfileDetail) (model.ProfileDetail, bool) {
	logginghelper.LogDebug("ResendOTPService() Start: ")
	var isEmail bool
	isEmail = profileobj.IsEmail
	msg := fmt.Sprintf("Hi %s ,your OTP for Changing Password is : ", profileobj.PersonalDetails.FullName)
	profile, status := otp.SendOTPService(profileobj, 3, msg)
	if !status {
		logginghelper.LogError("SendOTPDAO() Error")
		return profile, status
	}
	otp, err := sms.SendSingleSMS("", profile.ContactDetails.Mobile.Number, "otpmsg")
	if err != nil {
		logginghelper.LogError("Error while sending sms", err)
		return model.ProfileDetail{}, false
	}

	if isEmail {
		status, err := email.Sendmail(profile.ContactDetails.Email.Address, "Verifcaiton code : ", otp)
		logginghelper.LogInfo("Status=", status)
		if err != nil {
			logginghelper.LogError("Error while sending mail", err)
			return model.ProfileDetail{}, false
		}
	}
	otpValues := model.OTP{}
	otpValues.Username = profileobj.UserName
	otpValues.OTP = otp
	status, err = UpdateOTPDAO(otpValues)
	logginghelper.LogInfo("Status-Update OTP=", status)
	if err != nil {
		logginghelper.LogError("UpdateOTPDAO() Error", err)
		return model.ProfileDetail{}, false
	}
	logginghelper.LogDebug("ResendOTPService() Ended")
	return profile, status
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

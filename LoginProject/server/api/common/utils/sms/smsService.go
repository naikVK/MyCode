package sms

import (
	"LoginProject/server/api/common/utils"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/ddliu/go-httpclient"
)

//SendSingleSMS utility is for sending message on mobile
func SendSingleSMS(message string, mobileno string, smsservicetype string) (string, error) {
	var otpmsg string
	logginghelper.LogDebug("SMS Service() Start")
	activeSMSGateWay := confighelper.GetConfig("activeSMSGateWay")
	logginghelper.LogInfo("ActiveSMSGateway =", activeSMSGateWay)
	smsGatewayConfig := confighelper.GetConfig("activeSMSGateWay")
	logginghelper.LogInfo("SMSGatewayConfig =", smsGatewayConfig)
	url := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.url")
	username := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.username")
	password := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.password")
	senderId := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.senderId")
	cdmaHeader := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.cdmaHeader")

	if smsservicetype == "otpmsg" {
		otpmsg = utils.GetRandomOTP()
		message = message + otpmsg
		logginghelper.LogInfo("Message:", message)
	}
	res, err := httpclient.Post(url, map[string]string{
		"username":       username,
		"password":       password,
		"mobileno":       mobileno,
		"senderId":       senderId,
		"cdmaHeader":     cdmaHeader,
		"message":        message,
		"smsservicetype": smsservicetype,
	})
	logginghelper.LogInfo("Response=", res.StatusCode)
	if err != nil {
		logginghelper.LogError("Error while posting data :", err)
		return otpmsg, err
	}
	logginghelper.LogDebug("SMS service() End")
	return otpmsg, nil

}

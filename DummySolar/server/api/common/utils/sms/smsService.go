package sms

import (
	"LoginProject/server/api/common/utils"
	"fmt"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"

	"github.com/ddliu/go-httpclient"
)

func SendSingleSMS(message string, mobileno string, smsservicetype string) string {
	logginghelper.LogDebug("SMS Service() Start")
	confighelper.InitViper()
	activeSMSGateWay := confighelper.GetConfig("activeSMSGateWay")
	fmt.Println("activeSMSGateway", activeSMSGateWay)
	smsGatewayConfig := confighelper.GetConfig("activeSMSGateWay")
	fmt.Println("smsGatewayConfig", smsGatewayConfig)
	url := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.url")
	fmt.Println("url...", url)
	username := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.username")
	fmt.Println("username...", username)
	password := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.password")
	fmt.Println("password...", password)
	senderId := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.senderId")
	fmt.Println("senderId...", senderId)
	cdmaHeader := confighelper.GetConfig("smsGateway.smsGatewayURLMKCL.cdmaHeader")
	fmt.Println("cdmaHeader...", cdmaHeader)

	if smsservicetype == "otpmsg" {
		message = utils.GetRandomOTP()
		fmt.Println("message>>>>>>>>>>>>>>>>", message)

	}

	res, err := httpclient.Post(url, map[string]string{
		"Username":       username,
		"password":       password,
		"mobileno":       mobileno,
		"senderId":       senderId,
		"cdmaHeader":     cdmaHeader,
		"message":        message,
		"smsservicetype": smsservicetype,
	})
	if err != nil {
		logginghelper.LogError("Error while posting data :", err)
	}

	println(res.StatusCode)
	logginghelper.LogDebug("SMS service() End")
	return message

}

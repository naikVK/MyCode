package email

import (
	"LoginProject/server/api/common/constants"
	"github.com/spf13/viper"

	"crypto/tls"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"

	"github.com/go-gomail/gomail"
)

// SendEmail function to send otp as message or normal message
func Sendmail(toAddress string, body string, otp string) (bool,error) {
	logginghelper.LogDebug("SendEmail() Start:")

	message := gomail.NewMessage()

	fromAddress := constants.FROM_ADDRESS
	message.SetHeader("From", fromAddress)
	message.SetHeader("To", toAddress)

	content := body + otp

	message.SetHeader("Subject", "Email-Verfication")
	message.SetBody("text/html", content)
	port := viper.GetInt("emailConfig.port")
	host := confighelper.GetConfig("emailConfig.host")
	ssl := viper.GetBool("emailConfig.SSL")
	d := gomail.Dialer{Host: host, Port: port, SSL: ssl}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(message); err != nil {
		logginghelper.LogError(err)
	}
	logginghelper.LogDebug("SendEmail() End:")
	//REVIEW: optmsg must be generated properly. In case OTP is not generated then notify caller or throw an error so that appropriate handling can be done.
	return true,nil
}

package email

import (
	"crypto/tls"
	"fmt"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"

	"github.com/go-gomail/gomail"
)

func Sendmail(toAddress string, body string, otp string) string {
	logginghelper.LogDebug("SendEmail() Start:")

	confighelper.InitViper()
	message := gomail.NewMessage()
	fromAddress := confighelper.GetConfig("email.From")
	fmt.Println("fromaddress>>>>>", fromAddress)
	message.SetHeader("From", fromAddress)
	message.SetHeader("To", toAddress)

	content := body + otp
	fmt.Println("content>>>>>>>.", content)
	message.SetHeader("Subject", "Email-Verfication")
	message.SetBody("text/html", content)
	d := gomail.Dialer{Host: "10.1.70.100", Port: 25, SSL: false}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(message); err != nil {
		fmt.Println(err)
	}
	logginghelper.LogDebug("SendEmail() End:")
	return otp
}

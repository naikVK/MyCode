package google_auth

import (

	//  "crypto/rand"
	"LoginProject/server/api/common/constants"
	"LoginProject/server/api/common/model"

	"encoding/base32"
	"fmt"
	"io/ioutil"
	"net/url"

	dgoogauth "github.com/dgryski/dgoogauth"
	qr "github.com/rsc/qr"
)

const (
	qrFilename = "../qr.png"
)

var (
	// secret       = []byte{'D', 'e', 'l', 'l', 'o', '!', 0xDE, 0xAD, 0xBE, 0xEF}
	secret2 = ""
	account = "MKCL"
	issuer  = "loginproject"
)

func CreateQRCode(username string) bool {
	// secret := make([]byte, 6)
	// _, err := rand.Read(secret)
	// if err != nil {
	// 	panic(err)
	// }
	// secret := []byte{username}
	// secretBase32 = base32.StdEncoding.EncodeToString(secret)
	//  key := utils.GetRandamStringOfLen(16)
	// data := string(rand.Intn(99999999999)) + string(0xDE)
	data := username + string(0xDE)
	fmt.Println(data)
	secret := []byte(data)
	secretBase32 := base32.StdEncoding.EncodeToString(secret)
	secret2 = secretBase32
	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		panic(err)
	}

	URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)

	params := url.Values{}
	params.Add("secret", secretBase32)
	params.Add("issuer", issuer)

	URL.RawQuery = params.Encode()
	fmt.Printf("URL is %s\n", URL.String())

	code, err := qr.Encode(URL.String(), qr.Q)
	if err != nil {
		panic(err)
	}
	b := code.PNG()
	qrFilename := constants.QR_CODE_IMAGE_PATH + username + ".png"
	err = ioutil.WriteFile(qrFilename, b, 0600)
	if err != nil {
		panic(err)
	}
	fmt.Println("secretkey:", secretBase32)

	fmt.Printf("QR code is in %s. Please scan it into Google Authenticator app.\n", qrFilename)

	return true
}

func AuthenticateOTP(Auth_check model.Google_AuthCheck) bool {
	secret := []byte(Auth_check.UserName)
	secretBase32 := base32.StdEncoding.EncodeToString(secret)
	secret2 = secretBase32
	otpc := &dgoogauth.OTPConfig{
		Secret:      secret2,
		WindowSize:  3,
		HotpCounter: 0,
		// UTC:         true,
	}
	fmt.Println(Auth_check.OTP)
	fmt.Println(secret2)

	// fmt.Printf("Please enter the token value (or q to quit): ")
	// fmt.Scanln(&token)
	var token string
	token = Auth_check.OTP
	if token == "q" {
		fmt.Printf("break")
	}

	val, err := otpc.Authenticate(token)
	if err != nil {
		fmt.Println(err)

	}

	if !val {
		fmt.Println("Sorry, Not Authenticated")
		return false
	}

	fmt.Println("Authenticated!")

	return true
}

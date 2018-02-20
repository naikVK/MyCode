package utils

import (
	cryptorand "crypto/rand"
	"io"
	"math/rand"
	"time"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

//GET RANDOM STRING
func GetRandamStringOfLen(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//ENCODE STRING
func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(cryptorand.Reader, b, max)
	if n != max {
		logginghelper.LogError(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

//GET RANDOM 6 DIGIT PIN FOR OTP VERIFICATION
func GetRandomOTP() (content string) {
	content = encodeToString(6)
	return content
}

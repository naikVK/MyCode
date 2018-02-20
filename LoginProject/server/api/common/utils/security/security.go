package security

import (
	"LoginProject/server/api/common/constants"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"io"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"golang.org/x/crypto/bcrypt"
)

//Encrypt data & return
func Encrypt(name string, dob string) (string, string, error) {
	text1 := []byte(name)
	text2 := []byte(dob)
	keyconfig := confighelper.GetConfig(constants.ENCRYPT_DECRYPT_KEY_PATH)
	key := []byte(keyconfig) // 32 bytes
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}
	b1 := base64.StdEncoding.EncodeToString(text1)
	ciphertext1 := make([]byte, aes.BlockSize+len(b1))
	iv1 := ciphertext1[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv1); err != nil {
		return "", "", err
	}
	cfb1 := cipher.NewCFBEncrypter(block, iv1)
	cfb1.XORKeyStream(ciphertext1[aes.BlockSize:], []byte(b1))
	b2 := base64.StdEncoding.EncodeToString(text2)
	ciphertext2 := make([]byte, aes.BlockSize+len(b2))
	iv2 := ciphertext2[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv2); err != nil {
		return "", "", err
	}
	cfb2 := cipher.NewCFBEncrypter(block, iv2)
	cfb2.XORKeyStream(ciphertext2[aes.BlockSize:], []byte(b2))
	return string(ciphertext1), string(ciphertext2), nil
}

//Decrypt data & return
func Decrypt(name string, dob string) (string, string, error) {
	text1 := []byte(name)
	text2 := []byte(dob)
	confighelper.InitViper()
	keyconfig := confighelper.GetConfig(constants.ENCRYPT_DECRYPT_KEY_PATH)
	key := []byte(keyconfig) // 32 bytes
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}
	if len(text1) < aes.BlockSize {
		return "", "", errors.New("ciphertext too short")
	}
	iv1 := text1[:aes.BlockSize]
	text1 = text1[aes.BlockSize:]
	cfb1 := cipher.NewCFBDecrypter(block, iv1)
	cfb1.XORKeyStream(text1, text1)
	data1, err := base64.StdEncoding.DecodeString(string(text1))
	if err != nil {
		return "", "", err
	}
	if len(text2) < aes.BlockSize {
		return "", "", errors.New("ciphertext too short")
	}
	iv2 := text2[:aes.BlockSize]
	text2 = text2[aes.BlockSize:]
	cfb2 := cipher.NewCFBDecrypter(block, iv2)
	cfb2.XORKeyStream(text2, text2)
	data2, err := base64.StdEncoding.DecodeString(string(text2))
	if err != nil {
		return "", "", err
	}
	return string(data1), string(data2), nil
}
//Generating hash from password

func PasswordHashService(password string) string {
	logginghelper.LogDebug("PasswordHashService Start :")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logginghelper.LogError("ERROR_IN_GENERATING_HASH")
	}
	hashed := string(hash)
	logginghelper.LogDebug("PasswordHashService End :")
	return hashed
}

//Comparing db hashed pssword and password entered by user

func CompareHashPasswordService(userpassword string, hash []byte) (bool, error) {
	logginghelper.LogDebug("CompareHashPasswordService Start() :")
	err := bcrypt.CompareHashAndPassword(hash, []byte(userpassword))
	if err != nil {
		logginghelper.LogError("ERROR_WHILE_COMPARING")
		return false, err
	}
	return true, nil
}

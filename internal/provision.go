package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const provisionURL = "https://services.vip.symantec.com/prov"
const validateURL = "https://vip.symantec.com/otpCheck"

func Provision() error {
	request, err := GenerateRequest()
	if err != nil {
		return err
	}
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	requestReader := strings.NewReader(request)
	httpResponse, err := netClient.Post(provisionURL, "application/xml", requestReader)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}
	response, err := ProcessResponse(body)
	if err != nil {
		return err
	}
	secret, err := getSecret(response)

	otp := Generate6DigitTOTP(secret, time.Now(), 6)

	validationResponse, err := validate(response.SecretContainer.Secret.Id, otp)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", validationResponse)
	return nil
}

func getSecret(response *GetSharedSecretResponse) ([]byte, error) {
	var text []byte
	key := []byte{0x01, 0xad, 0x9b, 0xc6, 0x82, 0xa3, 0xaa, 0x93, 0xa9, 0xa3, 0x23, 0x9a, 0x86, 0xd6, 0xcc, 0xd9}
	text, err := base64.StdEncoding.DecodeString(response.SecretContainer.Secret.Cipher)
	if err != nil {
		return text, err
	}
	if len(text) < aes.BlockSize {
		return text, fmt.Errorf("ciphertext too short")
	}
	if len(text)%aes.BlockSize != 0 {
		return text, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return text, err
	}
	iv, err := base64.StdEncoding.DecodeString(response.SecretContainer.IV)
	if err != nil {
		return text, err
	}
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, text)
	return text, nil
}

func validate(ID string, otp string) (string, error) {
	form := url.Values{
		"cred":     {ID},
		"continue": {"otp_check"},
		"cr1":      {otp[0:1]},
		"cr2":      {otp[1:2]},
		"cr3":      {otp[2:3]},
		"cr4":      {otp[3:4]},
		"cr5":      {otp[4:5]},
		"cr6":      {otp[5:6]},
	}
	resp, err := http.PostForm(validateURL, form)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

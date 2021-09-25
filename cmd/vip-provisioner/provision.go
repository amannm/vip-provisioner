package main

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"
)

type GetSharedSecret struct {
	XMLName xml.Name `xml:"GetSharedSecret"`

	Id       int64  `xml:"Id,attr"`
	Version  string `xml:"Version,attr"`
	XmlNs    string `xml:"xmlns,attr"`
	XmlNsXsi string `xml:"xmlns:xsi,attr"`

	TokenModel                 string       `xml:"TokenModel"`
	ActivationCode             string       `xml:"ActivationCode"`
	OtpAlgorithm               OtpAlgorithm `xml:"OtpAlgorithm"`
	SharedSecretDeliveryMethod string       `xml:"SharedSecretDeliveryMethod"`
	Manufacturer               string       `xml:"DeviceId>Manufacturer"`
	SerialNo                   string       `xml:"DeviceId>SerialNo"`
	Model                      string       `xml:"DeviceId>Model"`
	Extension                  Extension    `xml:"Extension"`
}

type OtpAlgorithm struct {
	OtpAlgorithmType string `xml:"type,attr"`
}

type Extension struct {
	ExtVersion string `xml:"extVersion,attr"`
	XsiType    string `xml:"xsi:type,attr"`
	XmlNsVip   string `xml:"xmlns:vip,attr"`

	AppHandle       string `xml:"AppHandle"`
	ClientIDType    string `xml:"ClientIDType"`
	ClientID        string `xml:"ClientID"`
	DistChannel     string `xml:"DistChannel"`
	Os              string `xml:"ClientInfo>os"`
	Platform        string `xml:"ClientInfo>platform"`
	ClientTimestamp int64  `xml:"ClientTimestamp"`
	Data            string `xml:"Data"`
}

const vipServiceXmlNs = "http://www.verisign.com/2006/08/vipservice"
const digitsAndUppercaseLetters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const hexCharacters = "0123456789ABCDEF"

func printRequest() error {

	currentTime := time.Now().Unix()
	modelOs := fmt.Sprintf("MacBookPro%d,%d", randomInt(1, 12), randomInt(1, 4))
	serialNo := randomString(digitsAndUppercaseLetters, 12)
	clientID := fmt.Sprintf("Mac-%s", randomString(hexCharacters, 16))

	var request = &GetSharedSecret{
		Id:             currentTime,
		Version:        "2.0",
		XmlNs:          vipServiceXmlNs,
		XmlNsXsi:       "http://www.w3.org/2001/XMLSchema-instance",
		TokenModel:     "SYMC",
		ActivationCode: "",
		OtpAlgorithm: OtpAlgorithm{
			OtpAlgorithmType: "HMAC-SHA1-TRUNC-6DIGITS",
		},
		SharedSecretDeliveryMethod: "HTTPS",
		Manufacturer:               "Apple Inc.",
		SerialNo:                   serialNo,
		Model:                      modelOs,
		Extension: Extension{
			ExtVersion:      "auth",
			XsiType:         "vip:ProvisionInfoType",
			XmlNsVip:        vipServiceXmlNs,
			AppHandle:       "iMac010200",
			ClientIDType:    "BOARDID",
			ClientID:        clientID,
			DistChannel:     "Symantec",
			Os:              modelOs,
			Platform:        "iMac",
			ClientTimestamp: currentTime,
			Data:            "",
		},
	}

	out, err := xml.MarshalIndent(request, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(out))

	return nil
}

func randomString(choiceString string, outputLength int) string {
	numChoices := len(choiceString)
	output := make([]byte, outputLength)
	for i := 0; i < outputLength; i++ {
		randomIndex := rand.Intn(numChoices)
		output[i] = choiceString[randomIndex]
	}
	return string(output)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

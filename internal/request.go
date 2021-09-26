package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"time"
)

type GetSharedSecret struct {
	XMLName                    xml.Name     `xml:"GetSharedSecret"`
	Id                         int64        `xml:"Id,attr"`
	Version                    string       `xml:"Version,attr"`
	XmlNs                      string       `xml:"xmlns,attr"`
	XmlNsXsi                   string       `xml:"xmlns:xsi,attr"`
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
	ExtVersion      string `xml:"extVersion,attr"`
	XsiType         string `xml:"xsi:type,attr"`
	XmlNsVip        string `xml:"xmlns:vip,attr"`
	AppHandle       string `xml:"AppHandle"`
	ClientIDType    string `xml:"ClientIDType"`
	ClientID        string `xml:"ClientID"`
	DistChannel     string `xml:"DistChannel"`
	Os              string `xml:"ClientInfo>os"`
	Platform        string `xml:"ClientInfo>platform"`
	ClientTimestamp int64  `xml:"ClientTimestamp"`
	Data            string `xml:"Data"`
}

func GenerateRequest() (string, error) {
	const vipServiceXmlNs = "http://www.verisign.com/2006/08/vipservice"
	const digitsAndUppercaseLetters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const clientIDType = "BOARDID"
	const distChannel = "Symantec"
	const hexCharacters = "0123456789ABCDEF"
	currentTime := time.Now().Unix()
	modelOs := fmt.Sprintf("MacBookPro%d,%d", randomInt(1, 12), randomInt(1, 4))
	serialNo := randomString(digitsAndUppercaseLetters, 12)
	clientID := fmt.Sprintf("Mac-%s", randomString(hexCharacters, 16))
	messageData := []byte(fmt.Sprintf("%d%d%s%s%s", currentTime, currentTime, clientIDType, clientID, distChannel))
	var hmacKey = []byte{0xdd, 0x0b, 0xa6, 0x92, 0xc3, 0x8a, 0xa3, 0xa9, 0x93, 0xa3, 0xaa, 0x26, 0x96, 0x8c, 0xd9, 0xc2, 0xaa, 0x2a, 0xa2, 0xcb, 0x23, 0xb7, 0xc2, 0xd2, 0xaa, 0xaf, 0x8f, 0x8f, 0xc9, 0xa0, 0xa9, 0xa1}
	encodedHmacData := computeEncodedHmac(messageData, hmacKey)
	request := &GetSharedSecret{
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
			ClientIDType:    clientIDType,
			ClientID:        clientID,
			DistChannel:     distChannel,
			Os:              modelOs,
			Platform:        "iMac",
			ClientTimestamp: currentTime,
			Data:            encodedHmacData,
		},
	}
	requestXmlData, err := xml.MarshalIndent(request, "", "    ")
	if err != nil {
		return "", err
	}
	encodedXmlRequest := xml.Header + string(requestXmlData)
	return encodedXmlRequest, nil
}

func computeEncodedHmac(messageData []byte, hmacKey []byte) string {
	mac := hmac.New(sha256.New, hmacKey)
	mac.Write(messageData)
	digest := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(digest)
}

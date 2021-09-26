package internal

import (
	"encoding/xml"
	"time"
)

type GetSharedSecretResponse struct {
	XMLName                    xml.Name        `xml:"GetSharedSecretResponse"`
	RequestId                  int64           `xml:"RequestId,attr"`
	Version                    string          `xml:"Version,attr"`
	StatusReasonCode           string          `xml:"Status>ReasonCode"`
	StatusMessage              string          `xml:"Status>StatusMessage"`
	SharedSecretDeliveryMethod string          `xml:"SharedSecretDeliveryMethod"`
	SecretContainer            SecretContainer `xml:"SecretContainer"`
}
type SecretContainer struct {
	Version           string `xml:"Version,attr"`
	PBESalt           string `xml:"EncryptionMethod>PBESalt"`
	PBEIterationCount int    `xml:"EncryptionMethod>PBEIterationCount"`
	IV                string `xml:"EncryptionMethod>IV"`
	Secret            Secret `xml:"Device>Secret"`
}
type Secret struct {
	Type         string    `xml:"type,attr"`
	Id           string    `xml:"Id,attr"`
	Issuer       string    `xml:"Issuer"`
	Usage        Usage     `xml:"Usage"`
	FriendlyName string    `xml:"FriendlyName"`
	Cipher       string    `xml:"Data>Cipher"`
	Digest       Digest    `xml:"Data>Digest"`
	Expiry       time.Time `xml:"Expiry"`
}
type Usage struct {
	Otp        bool `xml:"otp,attr"`
	AI         AI   `xml:"AI"`
	TimeStep   int  `xml:"TimeStep"`
	Time       int  `xml:"Time"`
	ClockDrift int  `xml:"ClockDrift"`
}
type AI struct {
	Type string `xml:"type,attr"`
}
type Digest struct {
	Algorithm string `xml:"algorithm,attr"`
	Data      string `xml:",chardata"`
}

func ProcessResponse(responseData []byte) (*GetSharedSecretResponse, error) {
	var response GetSharedSecretResponse
	if err := xml.Unmarshal(responseData, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

package internal

import (
	"encoding/json"
	"testing"
)

const input = `<GetSharedSecretResponse RequestId="1632619149" Version="2.0" xmlns="http://www.verisign.com/2006/08/vipservice">
  <Status>
    <ReasonCode>0000</ReasonCode>
    <StatusMessage>Success</StatusMessage>
  </Status>
  <SharedSecretDeliveryMethod>HTTPS</SharedSecretDeliveryMethod>
  <SecretContainer Version="1.0">
    <EncryptionMethod>
      <PBESalt>Ma2L7oHJdBGVnJzUDr+aGY/aYYc=</PBESalt>
      <PBEIterationCount>50</PBEIterationCount>
      <IV>vqkQNXKvVUFE5MkLf9LjgA==</IV>
    </EncryptionMethod>
    <Device>
      <Secret type="HOTP" Id="SYMC88734197">
        <Issuer>OU = ID Protection Center, O = Symantec</Issuer>
        <Usage otp="true">
          <AI type="HMAC-SHA1-TRUNC-6DIGITS"/>
          <TimeStep>30</TimeStep>
          <Time>0</Time>
          <ClockDrift>4</ClockDrift>
        </Usage>
        <FriendlyName>OU = ID Protection Center, O = Symantec</FriendlyName>
        <Data>
          <Cipher>v0STrB5yFAngrWjYBgzn1dtsW9yCsFnmTsOSi68AXvk=</Cipher>
          <Digest algorithm="HMAC-SHA1">oIm/wd6x7km9Mw49IhKkzPENJ6A=</Digest>
        </Data>
        <Expiry>2024-09-25T01:19:08.966Z</Expiry>
      </Secret>
    </Device>
  </SecretContainer>
  <UTCTimestamp>1632619148</UTCTimestamp>
</GetSharedSecretResponse>`

func TestResponseParsing(t *testing.T) {
	response, err := ProcessResponse([]byte(input))
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	responseSerialized, _ := json.MarshalIndent(response, "", "    ")
	t.Logf("%s", string(responseSerialized))
}

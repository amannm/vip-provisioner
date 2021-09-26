package internal

import (
	"log"
	"testing"
	"time"
)

func TestTotp(t *testing.T) {
	totp := Generate6DigitTOTP([]byte("12345678901234567890"), time.Unix(59, 0), 8)
	if totp != "94287082" {
		log.Panicf("error: %s", totp)
	}
	totp = Generate6DigitTOTP([]byte("12345678901234567890"), time.Unix(20000000000, 0), 8)
	if totp != "65353130" {
		log.Panicf("error: %s", totp)
	}
}

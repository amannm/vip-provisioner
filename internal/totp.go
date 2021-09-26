package internal

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"
)

func Generate6DigitTOTP(secret []byte, t time.Time, digits int) string {
	counter := make([]byte, 8)
	binary.BigEndian.PutUint64(counter, uint64(t.Unix()/30))
	mac := hmac.New(sha1.New, secret)
	mac.Write(counter)
	hash := mac.Sum(nil)
	offset := hash[19] & 0xf
	value := int(hash[offset]&0x7f) << 24
	value |= int(hash[offset+1]&0xff) << 16
	value |= int(hash[offset+2]&0xff) << 8
	value |= int(hash[offset+3]&0xff) << 0
	return fmt.Sprintf("%0"+strconv.Itoa(digits)+"d", value%(int(math.Pow10(digits))))
}

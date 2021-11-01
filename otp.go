package password

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	Digits6 [2]int = [2]int{6, 1000000}
	Digits8 [2]int = [2]int{8, 100000000}

	Digits = Digits6

	OTP otp
)

type otp struct{}

func (otp) HOTP(secret string, interval int64) string {
	key, err := base64.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(interval))
	hasher := hmac.New(sha512.New, key)
	hasher.Write(buf)
	h := hasher.Sum(nil)
	offset := h[len(h)-1] & 0xf
	r := bytes.NewReader(h[offset : offset+4])

	var data uint32
	err = binary.Read(r, binary.BigEndian, &data)
	if err != nil {
		fmt.Println(err)
	}
	h12 := (int(data) & 0x7fffffff) % Digits[1]
	passcode := strconv.Itoa(h12)

	length := len(passcode)
	if length == Digits[0] {
		return passcode
	}
	for i := (Digits[0] - length); i > 0; i-- {
		passcode = "0" + passcode
	}
	return passcode
}

func (o *otp) TOTP(secret string) string {
	interval := time.Now().Unix() / 30
	return o.HOTP(secret, interval)
}

func (o *otp) Verify(secret string, input uint) bool {
	otp := o.TOTP(secret)
	return otp == strconv.Itoa(int(input))
}

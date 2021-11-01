package password

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base32"
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

	timestamp = time.Now().Unix() / 30

	OTP otp
)

type otp struct{}

func (otp) GenSecret() (string, error) {
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.BigEndian, timestamp)
	if err != nil {
		return "", err
	}
	hasher := hmac.New(sha512.New, buf.Bytes())
	secret := base32.StdEncoding.EncodeToString(hasher.Sum(nil))
	return secret, nil
}

func (otp) HOTP(secret string) string {
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(timestamp))
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
	return o.HOTP(secret)
}

func (o *otp) Verify(secret string, input uint) bool {
	otp := o.TOTP(secret)
	return otp == strconv.Itoa(int(input))
}

func NewOTP(account, issuer string) (string, error) {
	const uri string = "otpauth://totp/%s:%s?secret=%s&issuer=%s"
	secret, err := OTP.GenSecret()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(uri, issuer, account, secret, issuer), nil
}

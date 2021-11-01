package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/linzeyan/password"
)

const (
	usage = `Generate random string

Usage: password [option...]

Options:
`
)

var (
	operator  = flag.String("o", "pass", "Specify function(pass, encrypt)")
	length    = flag.Uint("length", 16, "Specify the password length")
	minLower  = flag.Uint("lower", 4, "Number of lowercase letters to include in the password")
	minUpper  = flag.Uint("upper", 2, "Number of uppercase letters to include in the password")
	minSymbol = flag.Uint("symbol", 2, "Number of symbols to include in the password")
	minNumber = flag.Uint("digit", 4, "Number of digits to include in the password")
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
	}
	flag.Parse()
	switch *operator {
	case "pass":
		pass()
	case "encrypt":
		encrypt()
	case "check":
		compareHash()
	case "otp":
		otp()
	}
}

func pass() {
	pass := password.Password.GeneratePassword(*length, *minLower, *minUpper, *minSymbol, *minNumber)
	fmt.Println(pass)
}

func encrypt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Encrypt password...")
	password.Encrypt.Hashed(text)
}

func compareHash() {
	var input *bufio.Reader
	input = bufio.NewReader(os.Stdin)
	fmt.Print("Enter hash: ")
	hash, _ := input.ReadString('\n')
	input = bufio.NewReader(os.Stdin)
	fmt.Print("Enter salt: ")
	salt, _ := input.ReadString('\n')
	input = bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	passWord, _ := input.ReadString('\n')
	result := password.Hash.CheckHash([]byte(hash), []byte(salt+passWord))
	if result {
		fmt.Println("Match")
	}
}

func otp() {
	secret := "ihI24nok/BpGqCu3W3FA6HqYEZo="
	otp := password.OTP.TOTP(secret)
	fmt.Println(otp)
}

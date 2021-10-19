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
	length    = flag.Int("length", 16, "Specify the password length")
	minLower  = flag.Int("lower", 4, "Number of lowercase letters to include in the password")
	minUpper  = flag.Int("upper", 2, "Number of uppercase letters to include in the password")
	minSymbol = flag.Int("symbol", 2, "Number of symbols to include in the password")
	minNumber = flag.Int("digit", 4, "Number of digits to include in the password")
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
	}

}

func pass() {
	check := *length - *minLower - *minUpper - *minSymbol - *minNumber
	if check >= 0 {
		pass := password.GeneratePassword(*length, *minLower, *minUpper, *minSymbol, *minNumber)
		fmt.Println(pass)
	} else {
		fmt.Println("Wrong values")
		os.Exit(1)
	}
}

func encrypt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Encrypt password...")
	salt := password.RandomBytes([]byte(password.GenAll(15)))
	result := password.HashPassword(salt + text)
	fmt.Printf(`{"salt":"%s","password":"%s"}`, salt, result)
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
	result := password.CheckHash(hash, salt, passWord)
	if result {
		fmt.Println("Match")
	}
}

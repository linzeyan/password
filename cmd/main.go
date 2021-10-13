package main

import (
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

	check := *length - *minLower - *minUpper - *minSymbol - *minNumber
	if check >= 0 {
		pass := password.GeneratePassword(*length, *minLower, *minUpper, *minSymbol, *minNumber)
		fmt.Println(pass)
	} else {
		fmt.Println("Wrong values")
		os.Exit(1)
	}
}

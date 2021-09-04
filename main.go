package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols      = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
	numbers      = "0123456789"
	allSet       = lowerLetters + upperLetters + symbols + numbers
)

func genString(length int, charSet string) []rune {
	rand.Seed(time.Now().Local().UnixNano())
	var s strings.Builder
	for i := 0; i < length; i++ {
		s.WriteByte(charSet[rand.Intn(len(charSet))])
	}
	random := []rune(s.String())
	rand.Shuffle(len(random), func(i, j int) {
		random[i], random[j] = random[j], random[i]
	})
	return random
}

func GenLower(length int) []rune {
	return genString(length, lowerLetters)
}

func GenUpper(length int) []rune {
	return genString(length, upperLetters)
}

func GenSymbol(length int) []rune {
	return genString(length, symbols)
}

func GenNumber(length int) []rune {
	return genString(length, numbers)
}

func GenAll(length int) []rune {
	return genString(length, allSet)
}

func GeneratePassword(length, minLower, minUpper, minSymbol, minNumber int) []rune {
	password := make([]rune, length)
	lower := GenLower(minLower)
	upper := GenUpper(minUpper)
	symbol := GenSymbol(minSymbol)
	num := GenNumber(minNumber)
	remain := GenAll(length - minLower - minUpper - minSymbol - minNumber)
	password = append(password, lower...)
	password = append(password, upper...)
	password = append(password, symbol...)
	password = append(password, num...)
	password = append(password, remain...)
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})
	return password
}

const (
	usage = `Generate random string

Usage: password [option...]

Options:
`
)

var (
	length    = flag.Int("length", 16, "Specify the password length")
	minLower  = flag.Int("lower", 4, "Number of lowercase letters to include in the password")
	minUpper  = flag.Int("upper", 2, "Number of uppsercase letters to include in the password")
	minSymbol = flag.Int("symbol", 2, "Number of symbols to include in the password")
	minNumber = flag.Int("digits", 4, "Number of digits to include in the password")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	check := *length - *minLower - *minUpper - *minSymbol - *minNumber
	if check >= 0 {
		password := GeneratePassword(*length, *minLower, *minUpper, *minSymbol, *minNumber)
		fmt.Println(string(password))
	} else {
		fmt.Println("Wrong values")
		os.Exit(1)
	}
}

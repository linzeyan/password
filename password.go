package password

import (
	"math/rand"
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

func genString(length int, charSet string) string {
	rand.Seed(time.Now().Local().UnixNano())
	var s strings.Builder
	for i := 0; i < length; i++ {
		s.WriteByte(charSet[rand.Intn(len(charSet))])
	}
	return s.String()
}

func GenLower(length int) string {
	return genString(length, lowerLetters)
}

func GenUpper(length int) string {
	return genString(length, upperLetters)
}

func GenSymbol(length int) string {
	return genString(length, symbols)
}

func GenNumber(length int) string {
	return genString(length, numbers)
}

func GenAll(length int) string {
	return genString(length, allSet)
}

func GeneratePassword(length, minLower, minUpper, minSymbol, minNumber int) string {
	lower := GenLower(minLower)
	upper := GenUpper(minUpper)
	symbol := GenSymbol(minSymbol)
	num := GenNumber(minNumber)
	remain := GenAll(length - minLower - minUpper - minSymbol - minNumber)
	password := []byte(lower + upper + symbol + num + remain)
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})
	return string(password)
}

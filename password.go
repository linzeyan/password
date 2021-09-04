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

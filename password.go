package password

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Character string

const (
	lowerLetters Character = "abcdefghijklmnopqrstuvwxyz"
	upperLetters Character = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols      Character = "~!@#$%^&*()_+`-={}|[]:<>?,./"
	numbers      Character = "0123456789"
	allSet       Character = lowerLetters + upperLetters + symbols + numbers
)

type password struct{}

func (*password) genString(length uint, charSet Character) string {
	rand.Seed(time.Now().Local().UnixNano())
	var s strings.Builder
	for i := uint(0); i < length; i++ {
		s.WriteByte(charSet[rand.Intn(len(charSet))])
	}
	return s.String()
}

func GenLower(length uint) string {
	var p password
	return p.genString(length, lowerLetters)
}

func GenUpper(length uint) string {
	var p password
	return p.genString(length, upperLetters)
}

func GenSymbol(length uint) string {
	var p password
	return p.genString(length, symbols)
}

func GenNumber(length uint) string {
	var p password
	return p.genString(length, numbers)
}

func GenAll(length uint) string {
	var p password
	return p.genString(length, allSet)
}

func GeneratePassword(length, minLower, minUpper, minSymbol, minNumber uint) string {
	var remain string
	var leave uint
	sum := minLower + minUpper + minSymbol + minNumber
	if length >= sum {
		leave = length - sum
	} else {
		log.Println("Wrong number")
		os.Exit(1)
	}
	lower := GenLower(minLower)
	upper := GenUpper(minUpper)
	symbol := GenSymbol(minSymbol)
	num := GenNumber(minNumber)
	if leave != 0 {
		remain = GenAll(leave)
	}
	result := []byte(lower + upper + symbol + num + remain)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return string(result)
}

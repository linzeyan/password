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

var now = time.Now().Local().UnixNano()

var Password password

type password struct{}

func (password) genString(length uint, charSet string) string {
	rand.Seed(now)
	var s strings.Builder
	for i := uint(0); i < length; i++ {
		s.WriteByte(charSet[rand.Intn(len(charSet))])
	}
	return s.String()
}

func (p *password) GenLower(length uint) string {
	return p.genString(length, lowerLetters)
}

func (p *password) GenUpper(length uint) string {
	return p.genString(length, upperLetters)
}

func (p *password) GenSymbol(length uint) string {
	return p.genString(length, symbols)
}

func (p *password) GenNumber(length uint) string {
	return p.genString(length, numbers)
}

func (p *password) GenAll(length uint) string {
	return p.genString(length, allSet)
}

func (p *password) GeneratePassword(length, minLower, minUpper, minSymbol, minNumber uint) string {
	var remain string
	leave := length - minLower - minUpper - minSymbol - minNumber
	lower := p.GenLower(minLower)
	upper := p.GenUpper(minUpper)
	symbol := p.GenSymbol(minSymbol)
	num := p.GenNumber(minNumber)
	if leave != 0 {
		remain = p.GenAll(leave)
	}
	result := []byte(lower + upper + symbol + num + remain)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return string(result)
}

package password

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func RandomBytes() string {
	/* Generate random salt */
	salt := make([]byte, 2)
	_, err := rand.Read(salt[:])
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	/* Create sha-512 hasher */
	hasher := sha512.New()
	salt = append(salt, byte(now))
	hasher.Write(salt)
	/* Convert the hashed to a base64 encoded string */
	s := base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println(hasher)
	fmt.Println(string(hasher.Sum(nil)))
	return s
}

func HashPassword(password string) string {
	const cost int = 15
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(passHash)
}

func CheckHash(hash, salt, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password)); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

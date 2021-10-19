package password

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func RandomBytes(seed []byte) string {
	/* Generate random salt */
	var salt = make([]byte, 2)
	if seed == nil {
		rand.Seed(now)
		_, err := rand.Read(seed[:])
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}
	}
	/* Create sha-512 hasher */
	hasher := sha512.New()
	salt = append(seed, byte(now))
	hasher.Write(salt)
	/* Convert the hashed to a base64 encoded string */
	s := base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
	return s
}

func HashPassword(password []byte) []byte {
	const cost int = 15
	passHash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return passHash
}

func CheckHash(hash, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

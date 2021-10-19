package password

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost int = 15
)

func RandomBytes(seed []byte, t int64) string {
	/* Generate random salt */
	var salt = make([]byte, 2)
	if seed == nil {
		rand.Seed(t)
		_, err := rand.Read(seed[:])
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}
	}
	/* Create sha-512 hasher */
	hasher := sha512.New()
	salt = append(seed, byte(t))
	hasher.Write(salt)
	/* Convert the hashed to a base64 encoded string */
	s := base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
	return s
}

func HashPassword(password []byte, cost int) []byte {
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

type Encrypt struct {
	seed string
	salt string
	hash string
	time int64
	cost int
}

func (e *Encrypt) Hashed(p string) {
	e.cost = cost
	e.time = now
	e.seed = GenAll(uint(len(p) + e.cost))
	e.salt = RandomBytes([]byte(p+e.seed), e.time)
	e.hash = string(HashPassword([]byte(e.salt+p), e.cost))
	fmt.Printf(`{"salt":"%s","password":"%s"}`, e.salt, e.hash)
}

func Enc(p string) {
	e := new(Encrypt)
	e.Hashed(p)
}

type Decrypt struct {
	hash string
	seed string
	time int64
}

func (d *Decrypt) Compare(p string) bool {
	salt := RandomBytes([]byte(p+d.seed), d.time)
	pw := HashPassword([]byte(salt+p), cost)
	return CheckHash([]byte(d.hash), pw)
}

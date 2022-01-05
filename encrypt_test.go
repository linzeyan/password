package password

import (
	"testing"
	"time"
)

func BenchmarkEncryptRandomBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash.RandomBytes([]byte(GenAll(64)), time.Now().Local().UnixNano())
	}
}

func BenchmarkEncryptRandomBytesWithNil(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash.RandomBytes(nil, time.Now().Local().UnixNano())
	}
}

func BenchmarkEncryptHashPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash.HashPassword([]byte(Hash.RandomBytes([]byte(GenAll(64)), time.Now().Local().UnixNano())), Cost)
	}
}

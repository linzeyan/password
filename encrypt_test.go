package password

import "testing"

func BenchmarkEncryptRandomBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash.RandomBytes([]byte(Password.GenAll(64)), now)
	}
}

func BenchmarkEncryptRandomBytesWithNil(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash.RandomBytes(nil, now)
	}
}

func BenchmarkEncryptHashPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash.HashPassword([]byte(Hash.RandomBytes([]byte(Password.GenAll(64)), now)), cost)
	}
}

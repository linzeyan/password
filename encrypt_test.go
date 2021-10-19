package password

import "testing"

func BenchmarkEncryptRandomBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := GenAll(64)
		RandomBytes([]byte(s))
	}
}

func BenchmarkEncryptRandomBytesWithNil(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomBytes(nil)
	}
}

func BenchmarkEncryptHashPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := GenAll(64)
		p := RandomBytes([]byte(s))
		HashPassword(p)
	}
}

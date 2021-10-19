package password

import "testing"

func BenchmarkEncryptRandomBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomBytes([]byte(GenAll(64)))
	}
}

func BenchmarkEncryptRandomBytesWithNil(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomBytes(nil)
	}
}

func BenchmarkEncryptHashPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		HashPassword([]byte(RandomBytes([]byte(GenAll(64)))))
	}
}

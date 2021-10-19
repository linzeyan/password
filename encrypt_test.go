package password

import "testing"

func BenchmarkRandomBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := GenAll(64)
		RandomBytes([]byte(s))
	}
}

func BenchmarkRandomBytesWithNil(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomBytes(nil)
	}
}

func BenchmarkHashPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := GenAll(64)
		p := RandomBytes([]byte(s))
		HashPassword(p)
	}
}

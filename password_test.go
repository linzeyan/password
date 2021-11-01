package password

import "testing"

func BenchmarkGenNumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Password.GenNumber(64)
	}
}

func BenchmarkGenSymbol(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Password.GenSymbol(64)
	}
}

func BenchmarkGenLower(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Password.GenLower(64)
	}
}

func BenchmarkGenUpper(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Password.GenUpper(64)
	}
}

func BenchmarkGenAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Password.GenAll(64)
	}
}

func BenchmarkGeneratePassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Password.GeneratePassword(64, 12, 12, 12, 12)
	}
}

package password

import "testing"

func BenchmarkGenNumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenNumber(64)
	}
}

func BenchmarkGenSymbol(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenSymbol(64)
	}
}

func BenchmarkGenLower(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenLower(64)
	}
}

func BenchmarkGenUpper(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenUpper(64)
	}
}

func BenchmarkGenAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenAll(64)
	}
}

func BenchmarkGeneratePassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GeneratePassword(64, 12, 12, 12, 12)
	}
}

package int2string

import "testing"

const number = 100

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	b.Logf("b.N is %d\n", b.N)
	for i := 0; i < b.N; i++ {
		ConvertSprintf(number)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	b.ResetTimer()
	// b.Logf("b.N is %d\n", b.N)
	for i := 0; i < b.N; i++ {
		ConvertFormatInt(number)
	}
}

func BenchmarkItoa(b *testing.B) {
	b.ResetTimer()
	// b.Logf("b.N is %d\n", b.N)
	for i := 0; i < b.N; i++ {
		ConvertItoa(number)
	}
}
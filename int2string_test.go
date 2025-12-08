package int2string

import "testing"

const number = 100

func BenchmarkSprintf(b *testing.B) {
	for b.Loop() {
		ConvertSprintf(number)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for b.Loop() {
		ConvertFormatInt(number)
	}
}

func BenchmarkItoa(b *testing.B) {
	for b.Loop() {
		ConvertItoa(number)
	}
}

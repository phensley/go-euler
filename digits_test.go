package euler

import (
	"fmt"
	"testing"
)

func blackhole(s string) {

}

func TestNewUintToDigits(t *testing.T) {
	for i := uint(0); i < 1000000; i++ {
		// fmt.Println(NumDigitsBase10(uint64(i)))
		UintToDigits(i)
	}
}

func BenchmarkUintToDigits(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UintToDigits(uint(i))
	}
}

func BenchmarkUintToString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blackhole(fmt.Sprintf("%d", uint(i)))
	}
}

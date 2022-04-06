package main

import (
	"testing"
)

func benchmarkFac(n int64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fac(n)
	}
}

func BenchmarkFac1(b *testing.B)  { benchmarkFac(1, b) }
func BenchmarkFac2(b *testing.B)  { benchmarkFac(2, b) }
func BenchmarkFac3(b *testing.B)  { benchmarkFac(3, b) }
func BenchmarkFac10(b *testing.B) { benchmarkFac(10, b) }
func BenchmarkFac20(b *testing.B) { benchmarkFac(20, b) }
func BenchmarkFac40(b *testing.B) { benchmarkFac(40, b) }

func TestFac(t *testing.T) {
	if Fac(5).Int64() != 120 {
		t.Error(Fac(5))
		t.Error("Factorial failed!!!")
	}
}

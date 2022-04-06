package utils

import (
	"testing"
)

func benchmarkSplit(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split(Gen(n), 5)
	}
}

// func BenchmarkSplit1(b *testing.B)     { Split(Gen(10), 5) }
func BenchmarkSplit100(b *testing.B)   { benchmarkSplit(100, b) }
func BenchmarkSplit200(b *testing.B)   { benchmarkSplit(200, b) }
func BenchmarkSplit300(b *testing.B)   { benchmarkSplit(300, b) }
func BenchmarkSplit1000(b *testing.B)  { benchmarkSplit(1000, b) }
func BenchmarkSplit2000(b *testing.B)  { benchmarkSplit(2000, b) }
func BenchmarkSplit4000(b *testing.B)  { benchmarkSplit(4000, b) }
func BenchmarkSplit10000(b *testing.B) { benchmarkSplit(10000, b) }

func TestSplitmator(t *testing.T) {
	if len(Split(Gen(5), 1)) != 5 {
		t.Error(Split(Gen(5), 1))
		t.Error("Splitmator failed!!!")
	}
}

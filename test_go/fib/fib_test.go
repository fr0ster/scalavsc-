package fib_test

import (
	. "test_go/fib"
	"testing"
)

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func TestFib(t *testing.T) {
	if Fib(44) != 701408733 {
		t.Error(Fib(44))
		t.Error("Factorial failed!!!")
	}
}

func ProxyTestAll(t *testing.T) {
	TestFib(t)
}

package fib_test

import (
	. "test_go/fib"
	"testing"
)

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib[int, int](i)
	}
}

func BenchmarkFib0(b *testing.B)     { benchmarkFib(0, b) }
func BenchmarkFib1(b *testing.B)     { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)     { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)     { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B)    { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B)    { benchmarkFib(20, b) }
func BenchmarkFib30(b *testing.B)    { benchmarkFib(30, b) }
func BenchmarkFib40(b *testing.B)    { benchmarkFib(40, b) }
func BenchmarkFib100(b *testing.B)   { benchmarkFib(100, b) }
func BenchmarkFib10000(b *testing.B) { benchmarkFib(10000, b) }

func _FibN(n int, res int, t *testing.T) {
	if Fib[int, int](n) != res {
		t.Error(Fib[int, int](n))
		t.Error("Factorial failed!!!")
	}
}
func TestFib0(t *testing.T) {
	_FibN(0, 0, t)
}
func TestFib5(t *testing.T) {
	_FibN(5, 5, t)
}
func TestFib10(t *testing.T) {
	_FibN(10, 55, t)
}

func ProxyTestAll(t *testing.T) {
	TestFib0(t)
}

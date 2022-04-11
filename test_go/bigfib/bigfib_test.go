package fib_test

import (
	"math/big"
	. "test_go/bigfib"
	"testing"
)

func benchmarkFib(i int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		BigFib(i)
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

func _FibN(n int64, res *big.Int, t *testing.T) {
	if BigFib(n).Cmp(res) != 0 {
		t.Error(BigFib(n))
		t.Error("Factorial failed!!!")
	}
}
func TestFib0(t *testing.T) {
	_FibN(0, big.NewInt(0), t)
}
func TestFib5(t *testing.T) {
	_FibN(5, big.NewInt(5), t)
}
func TestFib10(t *testing.T) {
	_FibN(10, big.NewInt(55), t)
}

func ProxyTestAll(t *testing.T) {
	TestFib0(t)
}

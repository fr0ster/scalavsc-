package sum_tc

import (
	. "test_go/utils"
	"testing"
)

func benchmarkSum(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(Gen(n))
	}
}

func BenchmarkSum100(b *testing.B)     { benchmarkSum(100, b) }
func BenchmarkSum200(b *testing.B)     { benchmarkSum(200, b) }
func BenchmarkSum300(b *testing.B)     { benchmarkSum(300, b) }
func BenchmarkSum1000(b *testing.B)    { benchmarkSum(1000, b) }
func BenchmarkSum2000(b *testing.B)    { benchmarkSum(2000, b) }
func BenchmarkSum4000(b *testing.B)    { benchmarkSum(4000, b) }
func BenchmarkSum10000(b *testing.B)   { benchmarkSum(10000, b) }
func BenchmarkSum1000000(b *testing.B) { benchmarkSum(1000000, b) }

// func BenchmarkSum100000000(b *testing.B) { benchmarkSum(100000000, b) }

func TestSum(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	if Sum(xs) != 15 {
		t.Error(Sum(xs))
		t.Error("Sum failed!!!")
	}
}

func TestSumLong(t *testing.T) {
	xs := make([]int, STEPS)
	for i := 0; i < len(xs); i++ {
		xs[i] = 1
	}
	if Sum(xs) != STEPS {
		t.Error(Sum(xs))
		t.Error("Sum failed!!!")
	}
}

func TestCrashSum(t *testing.T) {
	xs := Gen(1000)
	count := 10000
	for i := 0; i < count; i++ {
		Sum(xs)
	}
	t.Log("Loop Sum(1000) at {} times - OK!!!", count)
}

func ProxyTestAll(t *testing.T) {
	TestSum(t)
	TestSumLong(t)
}

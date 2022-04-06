package summator_test

import (
	. "test_go/summator"
	. "test_go/utils"
	"testing"
)

func benchmarkSummator(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Summator(Gen(n))
	}
}

func BenchmarkSummator100(b *testing.B)   { benchmarkSummator(100, b) }
func BenchmarkSummator200(b *testing.B)   { benchmarkSummator(200, b) }
func BenchmarkSummator300(b *testing.B)   { benchmarkSummator(300, b) }
func BenchmarkSummator1000(b *testing.B)  { benchmarkSummator(1000, b) }
func BenchmarkSummator2000(b *testing.B)  { benchmarkSummator(2000, b) }
func BenchmarkSummator4000(b *testing.B)  { benchmarkSummator(4000, b) }
func BenchmarkSummator10000(b *testing.B) { benchmarkSummator(10000, b) }

func TestSummator(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	if Summator(xs) != 15 {
		t.Error(Summator(xs))
		t.Error("Summatormator failed!!!")
	}
}

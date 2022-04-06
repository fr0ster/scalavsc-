package fold

import (
	. "test_go/utils"
	"testing"
)

func benchmarkFoldLeft(n int, b *testing.B) {
	xs := Gen(n)
	for i := 0; i < b.N; i++ {
		FoldLeft(xs[1:], xs[0], func(a int, b int) int { return a + b })
	}
}

func BenchmarkFoldLeft100(b *testing.B)   { benchmarkFoldLeft(100, b) }
func BenchmarkFoldLeft200(b *testing.B)   { benchmarkFoldLeft(200, b) }
func BenchmarkFoldLeft300(b *testing.B)   { benchmarkFoldLeft(300, b) }
func BenchmarkFoldLeft1000(b *testing.B)  { benchmarkFoldLeft(1000, b) }
func BenchmarkFoldLeft2000(b *testing.B)  { benchmarkFoldLeft(2000, b) }
func BenchmarkFoldLeft4000(b *testing.B)  { benchmarkFoldLeft(4000, b) }
func BenchmarkFoldLeft10000(b *testing.B) { benchmarkFoldLeft(10000, b) }

func TestFoldLeft(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	if FoldLeft(xs[1:], xs[0], func(a int, b int) int { return a + b }) != 15 {
		t.Error(FoldLeft(xs[1:], xs[0], func(a int, b int) int { return a + b }))
		t.Error("FoldLeft failed!!!")
	}
}

func TestCrashFoldLeft(t *testing.T) {
	xs := Gen(1000)
	count := 10000
	for i := 0; i < count; i++ {
		FoldLeft(xs[1:], xs[0], func(a int, b int) int { return a + b })
	}
	t.Log("Loop FoldLeft at {} times - OK!!!", count)
}

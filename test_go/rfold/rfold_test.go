package rfold

import (
	. "test_go/utils"
	"testing"
)

func benchmarkRFoldLeft(n int, b *testing.B) {
	xs := Gen(n)
	for i := 0; i < b.N; i++ {
		RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a + b })
	}
}

func BenchmarkRFoldLeft100(b *testing.B)   { benchmarkRFoldLeft(100, b) }
func BenchmarkRFoldLeft200(b *testing.B)   { benchmarkRFoldLeft(200, b) }
func BenchmarkRFoldLeft300(b *testing.B)   { benchmarkRFoldLeft(300, b) }
func BenchmarkRFoldLeft1000(b *testing.B)  { benchmarkRFoldLeft(1000, b) }
func BenchmarkRFoldLeft2000(b *testing.B)  { benchmarkRFoldLeft(2000, b) }
func BenchmarkRFoldLeft4000(b *testing.B)  { benchmarkRFoldLeft(4000, b) }
func BenchmarkRFoldLeft10000(b *testing.B) { benchmarkRFoldLeft(10000, b) }

func TestRFoldLeft(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	if RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a + b }) != 15 {
		t.Error(RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a + b }))
		t.Error("RFoldLeft failed!!!")
	}
}

func TestRFoldLeftBigGOMAX(t *testing.T) {
	xs := make([]int, 10000)
	for i := 0; i < len(xs); i++ {
		xs[i] = 1
	}
	if RFoldLeft(10000, xs[1:], xs[0], func(a int, b int) int { return a + b }) != 10000 {
		t.Error(RFoldLeft(10000, xs[1:], xs[0], func(a int, b int) int { return a + b }))
		t.Error("Summator failed!!!")
	}
}

func TestCrashRFoldLeft(t *testing.T) {
	xs := Gen(1000)
	count := 10000
	for i := 0; i < count; i++ {
		RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a + b })
	}
	t.Log("Loop RFoldLeft at {} times - OK!!!", count)
}

func ProxyTestAll(t *testing.T) {
	TestRFoldLeft(t)
	TestRFoldLeftBigGOMAX(t)
}

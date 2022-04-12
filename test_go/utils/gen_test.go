package utils

import (
	"testing"
)

func benchmarkGen(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gen(n)
	}
}

func BenchmarkGen100(b *testing.B)   { benchmarkGen(100, b) }
func BenchmarkGen200(b *testing.B)   { benchmarkGen(200, b) }
func BenchmarkGen300(b *testing.B)   { benchmarkGen(300, b) }
func BenchmarkGen1000(b *testing.B)  { benchmarkGen(1000, b) }
func BenchmarkGen2000(b *testing.B)  { benchmarkGen(2000, b) }
func BenchmarkGen4000(b *testing.B)  { benchmarkGen(4000, b) }
func BenchmarkGen10000(b *testing.B) { benchmarkGen(10000, b) }

func TestGenmator(t *testing.T) {
	if len(Gen(5)) != 5 {
		t.Error(Gen(5))
		t.Error("Genmator failed!!!")
	}
}

func test(t *testing.T) {

}

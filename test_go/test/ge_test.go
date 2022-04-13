package test

import (
	. "test_go/biglib"
	"testing"
)

func BenchmarkNewT(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		NewFromFloat(1000).Add(NewFromFloat(3000)).Cmp(NewFromFloat(4000))
	}
}

func TestNewT(t *testing.T) {
	if NewFromFloat(1000).Add(NewFromFloat(3000)).Cmp(NewFromFloat(4000)) != 0 {
		t.Error("Uncorrect Generic Add BigNumber method!!!")
	}
}

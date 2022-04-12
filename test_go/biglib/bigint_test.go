package biglib

import (
	"math"
	. "test_go/utils"
	"testing"
)

func gTest1[T Float](a T, b T) T {
	la := float64(a)
	lb := float64(b)
	return T(math.Pow(la, lb))
}

func gTest2[T Int](a T, b T) T {
	la := float64(a)
	lb := float64(b)
	return T(math.Pow(la, lb))
}

func gTest3[T Number](a T, b T) T {
	la := float64(a)
	lb := float64(b)
	return T(math.Pow(la, lb))
}

func BigAdd[T IBigNumber](a T, b T) IBigNumber {
	return a.Add(b)
}

func testI(x IBigNumber) string {
	return x.String()
}

func TestNewT(t *testing.T) {
	if NewFromInt(5555).Cmp(newInt(5555)) != 0 {
		t.Error("Uncorrect Generic Float BigNumber!!!")
	}
	if NewFromFloat(5555.55).Cmp(newFloat(5555.55)) != 0 {
		t.Error("Uncorrect Generic Float BigNumber!!!")
	}
	if NewFromFloat(5555.55).Add(NewFromFloat(1111.11)).Cmp(newFloat(6666.66)) != 0 {
		t.Error("Uncorrect Generic Add BigNumber method!!!")
	}
	if NewBigNumber[myBigFloat]().fromFloat(100).Cmp(newInt(100)) != 0 {
		t.Error("Uncorrect Generic Init BigNumber constructor!!!")
	}
}

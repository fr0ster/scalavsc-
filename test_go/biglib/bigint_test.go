package biglib

import (
	"math"
	"math/big"
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
	if NewInt(1234).value.Cmp(big.NewInt(1234)) != 0 {
		t.Error("Uncorrect MyBigInt!!!")
	}
	if NewFloat(4321).value.Cmp(big.NewFloat(4321)) != 0 {
		t.Error("Uncorrect MyBigFloat!!!")
	}
	if NewInt(1234).value.Cmp(big.NewInt(1234)) != 0 {
		t.Error("Uncorrect MyBigInt!!!")
	}
	if NewFloat(4321).value.Cmp(big.NewFloat(4321)) != 0 {
		t.Error("Uncorrect MyBigFloat!!!")
	}
	if NewBigNumber(NewInt(3333)).values.compare(NewInt(3333)) != 0 {
		t.Error("Uncorrect Generic Int BigNumber!!!")
	}
	if NewBigNumber(NewFloat(5555.55)).values.compare(NewFloat(5555.55)) != 0 {
		t.Error("Uncorrect Generic Float BigNumber!!!")
	}
	if NewBigNumber(NewFloat(5555.55)).compare(NewFloat(5555.55)) != 0 {
		t.Error("Uncorrect Generic Float BigNumber compare method!!!")
	}
	// if BigAdd(NewFloat(1000), NewFloat(3000)).compare(NewFloat(4000)) != 0 {
	// 	t.Error("Uncorrect Generic Add BigNumber method!!!")
	// }
}
